package entity

import (
	"encoding/json"
	"log"
)

// Update holds values from an update sent by the telegram server.
// At most one of the optional parameters can be present in any given update.
// // incomplete //
type Update struct {
	// UpdateID is the update's unique identifier.
	//
	// It is a required field.
	UpdateID int64 `json:"update_id,omitempty"`
	// Message is a new incoming message of any kind - text, photo, sticker, etc.
	Message *Message `json:"message,omitempty"`
	// EditedMessage is new version of a message that is known to the bot and was edited.
	EditedMessage *Message `json:"edited_message,omitempty"`
	// ChannelPost is new incoming channel post of any kind - text, photo, sticker, etc.
	ChannelPost *Message `json:"channel_post,omitempty"`
	// EditedChannelPost is new version of a channel post that is known to the bot and was edited.
	EditedChannelPost *Message `json:"edited_channel_post,omitempty"`
	// InlineQuery is new incoming InlineQuery.
	InlineQuery *InlineQuery `json:"inline_query,omitempty"`
	// ChosenInlineResult is the result of an inline query that was chosen by a user and sent to their chat partner.
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	// CallbackQuery is a new incoming callback query.
	CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`
}

// UpdateConfig holds methods for each kind of update
// only one of the methods is called for an update.
//
// TODO: incomplete
type UpdateConfig struct {
	// OnMessage is called if this update holds a new message.
	OnMessage func(message Message)
	// OnEditedMessage is called if this update holds an existing but edited message.
	OnEditedMessage func(editedMessage Message)
	// OnChannelPost is called if this update holds a new channel post.
	OnChannelPost func(channelPost Message)
	// OnEditedChannelPost is called if this update holds an existing but edited channel post.
	OnEditedChannelPost func(editedChannelPost Message)
	// OnInlineQuery is called if this update holds a new incoming inline query.
	OnInlineQuery func(inlineQuery InlineQuery)
	// OnChosenInlineResult is called if this update holds a result of an inline query
	// that was chosen by a user.
	OnChosenInlineResult func(chosenInlineResult ChosenInlineResult)
	// OnCallbackQuery is called if this update holds a new incoming callback query.
	OnCallbackQuery func(callbackQuery CallbackQuery)
}

// FromJSONBody modifies the current Update with matching values in body
func (u *Update) FromJSONBody(body []byte) {
	err := json.Unmarshal(body, u)
	if err != nil {
		log.Printf("error while unmarshalling update: %+v", err)
	}
}
