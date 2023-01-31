package entity

// MessageEnvelop holds the object that is used to send a new message
type MessageEnvelop struct {
	// ChatID is the unique identifier for the target chat
	// or username of the target channel
	//
	// It is a required field.
	ChatID string `json:"chat_id,omitempty"`
	// MessageThreadID is the unique identifier for the target message thread (topic) of the forum;
	// for forum supergroups only
	MessageThreadID int64 `json:"message_thread_id,omitempty"`
	// Text is text of the message to be sent, 1-4096 characters after entities parsing.
	//
	// It is a required field
	Text string `json:"text,omitempty"`
	// ParseMode is the mode for parsing entities in the message text.
	ParseMode string `json:"parse_mode,omitempty"`
	// Entities are A JSON-serialized list of special entities that appear in message text,
	// which can be specified instead of parse_mode.
	Entities []MessageEntity `json:"entities,omitempty"`
	// DisableWebPagePreview disables link previews for links in this message.
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`
	// DisableNotification sends the message silently.
	// Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// ProtectContent protects the contents of the sent message from forwarding and saving.
	ProtectContent bool `json:"protect_content,omitempty"`
	// ReplyToMessageID is, if the message is a reply, ID of the original message.
	ReplyToMessageID int64 `json:"reply_to_message_id,omitempty"`
	// AllowSendingWithoutReply can be true if the message should be sent
	// even if the specified replied-to message is not found.
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	// ReplyMarkup contains additional interface options.
	ReplyMarkup struct {
		*InlineKeyboardMarkup
		*ReplyKeyboardMarkup
		*ReplyKeyboardRemove
		*ForceReply
	} `json:"reply_markup,omitempty"`
}

// FileEnvelop represents a file to be uploaded to the telegram server.
type FileEnvelop struct {
	// Name is the field name for this file for the upload request.
	Name string
	// Path is the location of this file on this device.
	Path string
}
