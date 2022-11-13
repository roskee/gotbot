package gotbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gotbot/entity"
	"gotbot/router"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var apiString = "https://api.telegram.org/bot%s/%s"

// Bot is a gateway to manage a telegram bot
type Bot interface {

	// SendRawRequest sends a request to the telegram server and returns the result part of the response as a serialized json body
	SendRawRequest(httpMethod, function string, body interface{}) ([]byte, error)

	// RegisterMethod registers a new bot command with its name, description and implementation to the telegram server
	RegisterMethod(name, description string, function func(update entity.Update)) error

	// SendMessage is the implementation of the builtin sendMessage function of the bot.
	// It sends the given message to the sender user
	SendMessage(message entity.MessageEnvelop) (entity.Message, error)

	// GetMyCommands is the implementation of the builtin getMyCommands function of the bot.
	// It returns the list of all currently registered commands
	GetMyCommands() ([]entity.Command, error)

	// GetMe is the implementation of the builtin getMe function of the bot
	GetMe() (entity.User, error)

	// Listen creates a http server to listen for updates as a webhook handler.
	// It returns on failure only
	Listen(webhook entity.Webhook) error

	// Poll initiates a manual poll to get updates from the telegram server.
	// It returns on failure only
	Poll(duration time.Duration) error
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
func (b *bot) SendRawRequest(httpMethod, function string, body interface{}) ([]byte, error) {
	client := http.Client{}
	js, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(httpMethod, fmt.Sprintf(apiString, b.apiKey, function), bytes.NewBuffer(js))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
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
func (b *bot) Listen(webhook entity.Webhook) error {
	_, err := b.SendRawRequest("POST", "setWebhook", webhook)
	if err != nil {
		return err
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	return http.ListenAndServe(fmt.Sprintf(":%s", port), http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		path := req.URL.Path
		if !strings.Contains(path, b.apiKey) {
			res.WriteHeader(401)
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
		log.Printf("New Text Message: %+v\n", update.Message.Text)
		b.executeMethod(update.Message.GetCommand(), update)
		res.WriteHeader(200)
	}))
}

// Poll initiates a manual poll to get updates from the telegram server.
// It returns on failure only
func (b *bot) Poll(duration time.Duration) error {
	log.Printf("Polling with duration %v", duration)
	var lastUpdate entity.Update
	for {
		time.Sleep(duration)
		updatesJSON, err := b.SendRawRequest("GET", fmt.Sprintf("getUpdates?offset=%d", lastUpdate.UpdateID+1), nil)
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
			fmt.Printf("New Text Message: %+v\n", update.Message.Text)
			lastUpdate = update
			b.executeMethod(update.Message.GetCommand(), update)
		}
	}
}
