package gotbot

import (
	"encoding/json"
	"github.com/roskee/gotbot/entity"
)

// GetMe is the implementation of the builtin getMe function of the bot
func (b *bot) GetMe() (entity.User, error) {
	body, err := b.SendRawRequest("GET", "getMe", nil)
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
	cmdsJSON, err := b.SendRawRequest("GET", "getMyCommands", nil)
	if err != nil {
		return nil, err
	}
	var commands []entity.Command
	if err := json.Unmarshal(cmdsJSON, &commands); err != nil {
		return nil, err
	}
	return commands, nil
}

// setMyCommands is the implementation of the builtin setMyCommands function of the bot.
// It sets the given commands as the bot's command
func (b *bot) setMyCommands(commands []entity.Command) error {
	_, err := b.SendRawRequest("POST", "setMyCommands", entity.Commands{
		Commands: commands,
	})
	if err != nil {
		return err
	}
	return nil
}

// SendMessage is the implementation of the builtin sendMessage function of the bot.
// It sends the given message to the sender user
func (b *bot) SendMessage(message entity.MessageEnvelop) (entity.Message, error) {
	body, err := b.SendRawRequest("POST", "sendMessage", message)
	if err != nil {
		return entity.Message{}, err
	}
	var msg entity.Message
	err = json.Unmarshal(body, &msg)
	return msg, err
}
