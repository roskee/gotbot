package gotbot

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/roskee/gotbot/entity"
	"github.com/roskee/gotbot/router"
)

var apiString = "https://api.telegram.org/bot%s/%s"

// Bot is a gateway to manage a telegram bot
type Bot interface {

	// SendRawRequest sends a request to the telegram server and returns the result part of the response as a serialized json body
	SendRawRequest(httpMethod, function string, getBody func() (io.Reader, BodyOptions, error), setReq func(req *http.Request) error) ([]byte, error)

	// RegisterMethod registers a new bot command with its name, description and implementation to the telegram server
	RegisterMethod(name, description string, function func(update entity.Update)) error

	// SendMessage is the implementation of the builtin sendMessage function of the bot.
	// It sends the given message to the sender user
	SendMessage(msgType MessageType, message entity.MessageEnvelop, files ...entity.FileEnvelop) (entity.Message, error)

	// GetMyCommands is the implementation of the builtin getMyCommands function of the bot.
	// It returns the list of all currently registered commands
	GetMyCommands() ([]entity.Command, error)

	// GetMe is the implementation of the builtin getMe function of the bot
	GetMe() (entity.User, error)

	// Listen creates a http server to listen for updates as a webhook handler.
	// It returns on failure only
	Listen(port int, webhook entity.Webhook, config entity.UpdateConfig) error

	// Poll initiates a manual poll to get updates from the telegram server.
	// instructions on what to do on the updates should be set on config.
	// note that registered methods are automatically called.
	//
	// It returns on failure only
	Poll(duration time.Duration, configs entity.UpdateConfig) error

	// AnswerCallbackQuery send answers to callback queries sent from entity.InlineKeyboardMarkup.
	AnswerCallbackQuery(options entity.AnswerCallbackQueryEntity) error
}

// bot is in-package implementation of the Bot interface
type bot struct {
	apiKey  string
	methods []router.Handler
}

// NewBot returns a new bot with the token apiKey
func NewBot(apiKey string) Bot {
	return &bot{
		apiKey: apiKey,
	}
}

// SendRawRequest sends a request to the telegram server and returns the result part of the response as a serialized json body
func (b *bot) SendRawRequest(httpMethod, function string, getBody func() (io.Reader, BodyOptions, error), setReq func(req *http.Request) error) ([]byte, error) {
	var body io.Reader
	var options BodyOptions

	if getBody != nil {
		var err error
		body, options, err = getBody()
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(httpMethod, fmt.Sprintf(apiString, b.apiKey, function), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", options.ContentType)

	if setReq != nil {
		err = setReq(req)
		if err != nil {
			return nil, err
		}
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var resData entity.Response
	err = json.Unmarshal(resBody, &resData)
	if err != nil {
		return nil, err
	}
	if !resData.OK {
		return nil, fmt.Errorf("error response: code = %d, description = %s, parameters: %+v", resData.ErrorCode, resData.Description, resData.Parameters)
	}

	resultBody, err := json.Marshal(resData.Result)
	return resultBody, err
}

// RegisterMethod registers a new bot command with its name, description and implementation to the telegram server
func (b *bot) RegisterMethod(name, description string, function func(update entity.Update)) error {
	commands, err := b.GetMyCommands()
	if err != nil {
		return err
	}
	commands = append(commands, entity.Command{
		Command:     name,
		Description: description,
	})
	err = b.setMyCommands(commands)
	if err != nil {
		return err
	}
	b.methods = append(b.methods, router.Handler{
		Name:     name,
		Function: function,
	})
	return nil
}

// executeMethod executes the method specified by name. if the method with the name was not found it simply returns
func (b *bot) executeMethod(name string, update entity.Update) {
	for _, method := range b.methods {
		if method.Name == name {
			method.Function(update)
			break
		}
	}
}

// Listen creates a http server to listen for updates as a webhook handler.
// It returns on failure only
func (b *bot) Listen(port int, webhook entity.Webhook, config entity.UpdateConfig) error {
	_, err := b.SendRawRequest("POST", "setWebhook", func() (io.Reader, BodyOptions, error) {
		return GetJSONBody(webhook)
	}, SetApplicationJSON)
	if err != nil {
		return err
	}

	return http.ListenAndServe(fmt.Sprintf(":%d", port), http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		// check token
		if webhook.SecretToken != req.Header.Get("X-Telegram-Bot-Api-Secret-Token") {
			log.Printf("invalid secret token: %v", req.Header.Get("X-Telegram-Bot-Api-Secret-Token"))
			res.WriteHeader(http.StatusUnauthorized)
			return
		}

		body, err := io.ReadAll(req.Body)
		if err != nil {
			res.WriteHeader(400)
			return
		}

		var update entity.Update
		err = json.Unmarshal(body, &update)
		if err != nil {
			res.WriteHeader(500)
			return
		}

		b.executeUpdate(update, config)
		res.WriteHeader(200)
	}))
}

// Poll initiates a manual poll to get updates from the telegram server.
// instructions on what to do on the updates should be set on config.
// note that registered methods are automatically called.
//
// It returns on failure only
func (b *bot) Poll(duration time.Duration, config entity.UpdateConfig) error {
	log.Println("deleting webhook if exists")
	_, err := b.SendRawRequest(http.MethodPost, "deleteWebhook", nil, nil)
	if err != nil {
		return err
	}

	log.Printf("Polling with duration %v", duration)
	var lastUpdate entity.Update
	for {
		time.Sleep(duration)
		updatesJSON, err := b.SendRawRequest(http.MethodGet, fmt.Sprintf("getUpdates?offset=%d", lastUpdate.UpdateID+1), nil, nil)
		if err != nil {
			log.Printf("Error while polling: %+v", err)
			continue
		}
		var updates []entity.Update
		err = json.Unmarshal(updatesJSON, &updates)
		if err != nil {
			log.Printf("Error while polling: %+v", err)
			continue
		}
		for _, update := range updates {
			b.executeUpdate(update, config)
			lastUpdate = update
		}
	}
}

func (b *bot) executeUpdate(update entity.Update, config entity.UpdateConfig) {
	if update.Message != nil {
		if command := update.Message.GetCommand(); command != "" {
			b.executeMethod(command, update)
			log.Printf("%s command executed", command)
		}
		if config.OnMessage != nil {
			config.OnMessage(*update.Message)
		}
	} else if update.EditedMessage != nil {
		if config.OnEditedMessage != nil {
			config.OnEditedMessage(*update.EditedMessage)
		}
	} else if update.ChannelPost != nil {
		if config.OnChannelPost != nil {
			config.OnChannelPost(*update.ChannelPost)
		}
	} else if update.EditedChannelPost != nil {
		if config.OnEditedChannelPost != nil {
			config.OnEditedChannelPost(*update.EditedChannelPost)
		}
	} else if update.InlineQuery != nil {
		if config.OnInlineQuery != nil {
			config.OnInlineQuery(*update.InlineQuery)
		}
	} else if update.ChosenInlineResult != nil {
		if config.OnChosenInlineResult != nil {
			config.OnChosenInlineResult(*update.ChosenInlineResult)
		}
	} else if update.CallbackQuery != nil {
		if config.OnCallbackQuery != nil {
			config.OnCallbackQuery(*update.CallbackQuery)
		}
	} else { // TODO: missing update types
		log.Printf("unknown update: %+v", update)
	}
}
