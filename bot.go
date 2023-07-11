package gotbot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/roskee/gotbot/entity"
	"github.com/roskee/gotbot/envelop"
	"github.com/roskee/gotbot/router"
)

var apiString = "https://api.telegram.org/bot%s/%s"
var apiFileString = "https://api.telegram.org/file/bot%s/%s"

// Bot is a gateway to manage a telegram bot
type Bot interface {

	// SendRawRequest sends a request to the telegram server and returns the result part of the response as a serialized json body
	SendRawRequest(httpMethod, function string, getBody func() (io.Reader, BodyOptions, error), setReq func(req *http.Request) error) ([]byte, error)

	// RegisterMethod registers a new bot command with its name, description and implementation to the telegram server
	RegisterMethod(name, description string, function func(update entity.Update)) error

	// SendMessage is the implementation of the builtin sendMessage function of the bot.
	// It sends the given message to the sender user
	SendMessage(msg entity.MessageEnvelop) (entity.Message, error)

	// SendMessageAny can be used to send any kind of message manually.
	// All other Send* messages use this method internally.
	SendMessageAny(msgType MessageType, message entity.MessageEnvelop, response any, attachedFiles ...entity.FileEnvelop) error

	// GetMyCommands is the implementation of the builtin getMyCommands function of the bot.
	// It returns the list of all currently registered commands
	GetMyCommands() ([]entity.Command, error)

	// SetMyCommands is the implementation of the builtin setMyCommands function of the bot.
	SetMyCommands(commands []entity.Command) error

	// DeleteMyCommands is the implementation of the builtin deleteMyCommands function of the bot.
	DeleteMyCommands(commandScope envelop.DeleteMyCommandsEnvelop) (bool, error)

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

	// ForwardMessage is used to forward messages of any kind.
	// Service messages can't be forwarded.
	ForwardMessage(msgEnvelop envelop.ForwardMessageEnvelop) (entity.Message, error)

	// CopyMessage is used to copy messages of any kind.
	// Service messages and invoice messages can't be copied.
	// A quiz poll can be copied only if the value of
	// the field 'correct_option_id' is known to the bot.
	CopyMessage(msgEnvelop envelop.CopyMessageEnvelop) (int64, error)

	// SendPhoto is used to send photos.
	SendPhoto(msg entity.MessageEnvelop) (entity.Message, error)

	// SendAudio is used to send audio files.
	// For telegrm to show the audio in the music player,
	// it must be in the format .mp3 or .m4a.
	//
	// Note: bots can only send audio files up to 50 MB in size.
	SendAudio(msg entity.MessageEnvelop) (entity.Message, error)

	// SendVideo is used to send video files.
	// Only MPEG4 videos are supported.
	// (other formats can be sent as a document)
	//
	// Note: bots can only send video files up to 50 MB in size.
	SendVideo(msg entity.MessageEnvelop) (entity.Message, error)

	// SendLocation is used to send location
	SendLocation(msg entity.MessageEnvelop) (entity.Message, error)
	// SendDocument is used to send general files.
	//
	// Note: bots can only send files of any type up to 50 MB in size.

	SendDocument(msg entity.MessageEnvelop) (entity.Message, error)
	// SendVoice is used to send audio files.
	//
	// Note: bots can only send voice messages up to 50 MB in size.

	SendVoice(msg entity.MessageEnvelop) (entity.Message, error)

	// SendMediaGroup is used to send a group of photos, videos, documents or audios as an album.
	//
	// Note: Documents and audio files can be only grouped in an album with messages of the same type.
	SendMediaGroup(msg entity.MessageEnvelop) ([]entity.Message, error)

	// SendVideoNote is used to send rounded square mp4 videos of up to 1 minute long.
	SendVideoNote(msg entity.MessageEnvelop) (entity.Message, error)

	// SendContact is used to send phone contacts.
	SendContact(msg entity.MessageEnvelop) (entity.Message, error)
	// GetUserProfilePhotos is used to get a list of profile pictures for a user.
	GetUserProfilePhotos(options envelop.GetUserProfilePhotos) (entity.UserProfilePhotos, error)
	// GetFile is used to get basic info about a file and prepare it for downloading.
	// For the moment, bots can download files of up to 20MB in size.
	GetFile(options envelop.GetFile) (entity.File, error)
	// DownloadFile downloads a file from the telegram server.
	DownloadFile(file entity.File) ([]byte, error)
	// SendPoll is used to send a native poll.
	SendPoll(msg entity.MessageEnvelop) (entity.Message, error)
	// SendChatAction is used to send a chat action.
	SendChatAction(msg entity.MessageEnvelop) (bool, error)
	// SendAnimation is used to send animation files.
	// For the moment, bots can send animation files of up to 50 MB in size.
	SendAnimation(msg entity.MessageEnvelop) (entity.Message, error)
	// SendDice is used to send an animated emoji that will display a random value.
	SendDice(msg entity.MessageEnvelop) (entity.Message, error)
	// SendVenue is used to send information about a venue.
	SendVenue(msg entity.MessageEnvelop) (entity.Message, error)
}

// BotOptions hold the options for the bot
type BotOptions struct {
	// Logger is the logger to use for logging
	Logger Logger
	// Client is the http client to use for sending requests
	Client http.Client
}

func setDefaultOptions(o BotOptions) BotOptions {
	if o.Logger == nil {
		o.Logger = &JSONLogger{
			TimeFormat: time.RFC3339,
		}
	}

	return o
}

// bot is in-package implementation of the Bot interface
type bot struct {
	apiKey  string
	methods []router.Handler
	options BotOptions
}

// NewBot returns a new bot with the token apiKey
func NewBot(apiKey string, options BotOptions) Bot {
	return &bot{
		apiKey:  apiKey,
		options: setDefaultOptions(options),
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

	res, err := b.options.Client.Do(req)
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

// SendMessageAny can be used to send any kind of message manually.
// All the default send functions use this internally.
func (b *bot) SendMessageAny(messageType MessageType, message entity.MessageEnvelop, response any, attachedFiles ...entity.FileEnvelop) error {
	var res []byte
	var err error

	res, err = b.SendRawRequest(http.MethodPost, string(messageType), func() (io.Reader, BodyOptions, error) {
		return GetMultipartBody(message, attachedFiles...)
	}, nil)
	if err != nil {
		return err
	}

	return json.Unmarshal(res, response)
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
	err = b.SetMyCommands(commands)
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
			b.options.Logger.Warn("invalid secret token", Fields{
				"token": req.Header.Get("X-Telegram-Bot-Api-Secret-Token"),
			})
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
	b.options.Logger.Info("deleting webhook if exists", Fields{})

	_, err := b.SendRawRequest(http.MethodPost, "deleteWebhook", nil, nil)
	if err != nil {
		b.options.Logger.Error("error while deleting webhook", Fields{
			"error": err.Error(),
		})

		return err
	}

	b.options.Logger.Info("Polling started", Fields{
		"duration": fmt.Sprintf("%fs", duration.Seconds()),
	})

	var lastUpdate entity.Update
	for {
		time.Sleep(duration)
		updatesJSON, err := b.SendRawRequest(http.MethodGet, fmt.Sprintf("getUpdates?offset=%d", lastUpdate.UpdateID+1), nil, nil)
		if err != nil {
			b.options.Logger.Error("Error while polling", Fields{
				"error": err.Error(),
			})

			continue
		}
		var updates []entity.Update
		err = json.Unmarshal(updatesJSON, &updates)
		if err != nil {
			b.options.Logger.Error("Error while polling", Fields{
				"error": err.Error(),
			})

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
			b.options.Logger.Debug("command executed", Fields{
				"command": command,
			})
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
		b.options.Logger.Warn("unknown update", Fields{
			"update": update,
		})
	}
}

// SetLogger sets the logger of the bot.
func (b *bot) SetLogger(logger Logger) {
	b.options.Logger = logger
}

func (b *bot) DownloadFile(file entity.File) ([]byte, error) {
	res, err := b.options.Client.Get(fmt.Sprintf(apiFileString, b.apiKey, file.FilePath))
	if err != nil {
		return nil, err
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			b.options.Logger.Error("error while closing response body", Fields{
				"error": err.Error(),
			})
		}
	}()

	return io.ReadAll(res.Body)
}
