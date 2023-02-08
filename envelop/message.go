package envelop

import "github.com/roskee/gotbot/entity"

// ForwardMessageEnvelop is used to forward a message
type ForwardMessageEnvelop struct {
	// ChatID is the id for the target chat or username.
	//
	// It is a required field.
	ChatID string `json:"chat_id,omitempty"`
	// MessageThreadID is the unique identifier for the target message thread
	// (topic) of the forum;
	// for forum super groups only.
	MessageThreadID int64 `json:"message_thread_id,omitempty"`
	// FromChatID id of the chat where the original message was sent.
	//
	// It is a required field.
	FromChatID string `json:"from_chat_id,omitempty"`
	// DisableNotification is to send the message silently.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// ProtectContent is to protects the contents of the forwarded message
	// from forwarding and saving.
	ProtectContent bool `json:"protect_content,omitempty"`
	// MessageID is the id of the message to forward.
	//
	// It is a required field.
	MessageID int64 `json:"message_id,omitempty"`
}

// CopyMessageEnvelop is used to copy a message
type CopyMessageEnvelop struct {
	// ChatID is the id for the target chat or username.
	//
	// It is a required field.
	ChatID string `json:"chat_id,omitempty"`
	// MessageThreadID is the unique identifier for the target message thread
	// (topic) of the forum;
	// for forum super groups only.
	MessageThreadID int64 `json:"message_thread_id,omitempty"`
	// FromChatID id of the chat where the original message was sent.
	//
	// It is a required field.
	FromChatID string `json:"from_chat_id,omitempty"`
	// MessageID is the id of the message to copy.
	//
	// It is a required field.
	MessageID int64 `json:"message_id,omitempty"`
	// Caption is a new caption for media.
	// If not specified, the original caption is kept.
	Caption string `json:"caption,omitempty"`
	// ParseMode is mode for parsing entities in the new caption.
	ParseMode string `json:"parse_mode,omitempty"`
	// CaptionEntities is a JSON-serialized list of special entities
	// that appear in the new caption.
	CaptionEntities []entity.MessageEntity `json:"caption_entities,omitempty"`
	// DisableNotification is to send the message silently.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// ProtectContent is to protects the contents of the sent message
	// from forwarding and saving.
	ProtectContent bool `json:"protect_content,omitempty"`
	// ReplyToMessageID is the id of the original message,
	// if the message is a reply.
	ReplyToMessageID int64 `json:"reply_to_message_id,omitempty"`
	// AllowSendingWithoutReply can be true
	// if the message should be sent
	// even if the specified replied-to message is not found.
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	// ReplyMarkup as an additional interface options.
	ReplyMarkup entity.ReplyMarkup `json:"reply_markup,omitempty"`
}
