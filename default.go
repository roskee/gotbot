package gotbot

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/roskee/gotbot/entity"
	"github.com/roskee/gotbot/envelop"
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

// SetMyCommands is the implementation of the builtin setMyCommands function of the bot.
// It sets the given commands as the bot's command
func (b *bot) SetMyCommands(commands []entity.Command) error {
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

// SendMessage is the implementation of the builtin sendMessage function of the bot.
// It sends the given message to the sender user
func (b *bot) SendMessage(message entity.MessageEnvelop) (entity.Message, error) {
	var res entity.Message

	err := b.SendMessageAny(MessageText, message, &res)

	return res, err
}

func (b *bot) AnswerCallbackQuery(options entity.AnswerCallbackQueryEntity) error {
	_, err := b.SendRawRequest(http.MethodPost, "answerCallbackQuery", func() (io.Reader, BodyOptions, error) {
		return GetJSONBody(options)
	}, SetApplicationJSON)
	return err
}

func (b *bot) ForwardMessage(msgEnvelop envelop.ForwardMessageEnvelop) (entity.Message, error) {
	res, err := b.SendRawRequest(http.MethodPost, "forwardMessage", func() (io.Reader, BodyOptions, error) {
		return GetJSONBody(msgEnvelop)
	}, SetApplicationJSON)
	if err != nil {
		return entity.Message{}, err
	}

	var msg entity.Message

	return msg, json.Unmarshal(res, &msg)
}

func (b *bot) CopyMessage(msgEnvelop envelop.CopyMessageEnvelop) (int64, error) {
	res, err := b.SendRawRequest(http.MethodPost, "copyMessage", func() (io.Reader, BodyOptions, error) {
		return GetJSONBody(msgEnvelop)
	}, SetApplicationJSON)
	if err != nil {
		return 0, err
	}

	var msgID entity.Message

	return msgID.MessageID, json.Unmarshal(res, &msgID)
}

func (b *bot) SendPhoto(msg entity.MessageEnvelop) (entity.Message, error) {
	var res entity.Message

	err := b.SendMessageAny(MessagePhoto, msg, &res)

	return res, err
}

func (b *bot) SendAudio(msg entity.MessageEnvelop) (entity.Message, error) {
	var res entity.Message

	err := b.SendMessageAny(MessageAudio, msg, &res)

	return res, err
}

func (b *bot) SendVideo(msg entity.MessageEnvelop) (entity.Message, error) {
	var res entity.Message

	err := b.SendMessageAny(MessageVideo, msg, &res)

	return res, err
}

func (b *bot) SendLocation(msg entity.MessageEnvelop) (entity.Message, error) {
	var res entity.Message

	err := b.SendMessageAny(MessageLocation, msg, &res)
	return res, err
}

func (b *bot) SendDocument(msg entity.MessageEnvelop) (entity.Message, error) {
	var res entity.Message

	err := b.SendMessageAny(MessageDocument, msg, &res)

	return res, err
}

func (b *bot) SendVoice(msg entity.MessageEnvelop) (entity.Message, error) {
	var res entity.Message

	err := b.SendMessageAny(MessageVoice, msg, &res)

	return res, err
}

func (b *bot) SendMediaGroup(msg entity.MessageEnvelop) ([]entity.Message, error) {
	var res []entity.Message

	err := b.SendMessageAny(MessageMediaGroup, msg, &res)

	return res, err
}

func (b *bot) SendVideoNote(msg entity.MessageEnvelop) (entity.Message, error) {
	var res entity.Message

	err := b.SendMessageAny(MessageVideoNote, msg, &res)

	return res, err
}

func (b *bot) SendContact(msg entity.MessageEnvelop) (entity.Message, error) {
	var res entity.Message

	err := b.SendMessageAny(MessageContact, msg, &res)

	return res, err
}

func (b *bot) GetUserProfilePhotos(options envelop.GetUserProfilePhotos) (entity.UserProfilePhotos, error) {
	res, err := b.SendRawRequest(http.MethodPost, "getUserProfilePhotos", func() (io.Reader, BodyOptions, error) {
		return GetJSONBody(options)
	}, SetApplicationJSON)
	if err != nil {
		return entity.UserProfilePhotos{}, err
	}

	var photos entity.UserProfilePhotos

	return photos, json.Unmarshal(res, &photos)
}

func (b *bot) GetFile(getFile envelop.GetFile) (entity.File, error) {
	res, err := b.SendRawRequest(http.MethodPost, "getFile", func() (io.Reader, BodyOptions, error) {
		return GetJSONBody(getFile)
	}, SetApplicationJSON)
	if err != nil {
		return entity.File{}, err
	}

	var file entity.File

	return file, json.Unmarshal(res, &file)
}

func (b *bot) SendPoll(msg entity.MessageEnvelop) (entity.Message, error) {
	var res entity.Message

	err := b.SendMessageAny(MessagePoll, msg, &res)

	return res, err
}

func (b *bot) SendChatAction(msg entity.MessageEnvelop) (bool, error) {
	var res bool

	err := b.SendMessageAny(MessageChatAction, msg, &res)

	return res, err
}

func (b *bot) SendAnimation(msg entity.MessageEnvelop) (entity.Message, error) {
	var res entity.Message

	err := b.SendMessageAny(MessageAnimation, msg, &res)

	return res, err
}

func (b *bot) SendDice(msg entity.MessageEnvelop) (entity.Message, error) {
	var res entity.Message

	err := b.SendMessageAny(MessageDice, msg, &res)

	return res, err
}

func (b *bot) SendVenue(msg entity.MessageEnvelop) (entity.Message, error) {
	var res entity.Message

	err := b.SendMessageAny(MessageVenue, msg, &res)

	return res, err
}
