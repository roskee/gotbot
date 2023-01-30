package gotbot

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/roskee/gotbot/entity"
)

// GetMe is the implementation of the builtin getMe function of the bot
func (b *bot) GetMe() (entity.User, error) {
	body, err := b.SendRawRequest(http.MethodGet, "getMe", nil, nil)
	if err != nil {
		return entity.User{}, err
	}
	user := entity.User{}
	user.FromJSONBody(body)
	return user, nil
}

// GetMyCommands is the implementation of the builtin getMyCommands function of the bot.
// It returns the list of all currently registered commands
func (b *bot) GetMyCommands() ([]entity.Command, error) {
	res, err := b.SendRawRequest(http.MethodGet, "getMyCommands", nil, nil)
	if err != nil {
		return nil, err
	}
	var commands []entity.Command
	if err := json.Unmarshal(res, &commands); err != nil {
		return nil, err
	}
	return commands, nil
}

// setMyCommands is the implementation of the builtin setMyCommands function of the bot.
// It sets the given commands as the bot's command
func (b *bot) setMyCommands(commands []entity.Command) error {
	_, err := b.SendRawRequest(http.MethodPost, "setMyCommands", func() (io.Reader, BodyOptions, error) {
		return GetJSONBody(entity.Commands{
			Commands: commands,
		})
	}, SetApplicationJSON)
	if err != nil {
		return err
	}
	return nil
}

func (b *bot) sendMessage(messageType MessageType, message entity.MessageEnvelop, files ...entity.FileEnvelop) (entity.Message, error) {
	var res []byte
	var err error

	res, err = b.SendRawRequest(http.MethodPost, string(messageType), func() (io.Reader, BodyOptions, error) {
		return GetMultipartBody(message, files...)
	}, nil)

	var msg entity.Message
	err = json.Unmarshal(res, &msg)
	return msg, err
}

// SendMessage is the implementation of the builtin sendMessage function of the bot.
// It sends the given message to the sender user
func (b *bot) SendMessage(message entity.MessageEnvelop) (entity.Message, error) {
	msg, err := b.sendMessage(MessageText, message)
	if err != nil {
		return entity.Message{}, err
	}

	return msg, err
}

func (b *bot) AnswerCallbackQuery(options entity.AnswerCallbackQueryEntity) error {
	_, err := b.SendRawRequest(http.MethodPost, "answerCallbackQuery", func() (io.Reader, BodyOptions, error) {
		return GetJSONBody(options)
	}, SetApplicationJSON)
	return err
}
