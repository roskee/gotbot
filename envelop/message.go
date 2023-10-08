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

// EditMessageTextEnvelop is used to edit text and game messages.
type EditMessageTextEnvelop struct {
	// ChatID is the id for the target chat or username.
	//
	// It is required if inline_message_id is not specified.
	ChatID string `json:"chat_id,omitempty"`
	// MessageID is the id of the message to edit.
	//
	// It is required if inline_message_id is not specified.
	MessageID int64 `json:"message_id,omitempty"`
	//InlineMessageID is the id of the inline message to edit.
	//
	// It is required if chat_id and message_id are not specified.
	InlineMessageID string `json:"inline_message_id,omitempty"`
	// Text is the new text of the message.
	//
	// It is a required field.
	Text string `json:"text,omitempty"`
	// ParseMode is mode for parsing entities in the message text.
	ParseMode string `json:"parse_mode,omitempty"`
	// Entities is a JSON-serialized list of special entities
	// that appear in the message text,
	// which can be specified instead of parse_mode.
	Entities []entity.MessageEntity `json:"entities,omitempty"`
	// DisableWebPagePreview is to disable link previews for links in the sent message.
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`
	// ReplyMarkup as an additional interface options.
	ReplyMarkup entity.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// EditMessageCaptionEnvelop is used to edit captions of messages.
type EditMessageCaptionEnvelop struct {
	// ChatID is the id for the target chat or username.
	//
	// It is required if inline_message_id is not specified.
	ChatID string `json:"chat_id,omitempty"`
	// MessageID is the id of the message to edit.
	//
	// It is required if inline_message_id is not specified.
	MessageID int64 `json:"message_id,omitempty"`
	//InlineMessageID is the id of the inline message to edit.
	//
	// It is required if chat_id and message_id are not specified.
	InlineMessageID string `json:"inline_message_id,omitempty"`
	// Caption is the new caption of the message.
	Caption string `json:"caption,omitempty"`
	// ParseMode is mode for parsing entities in the message text.
	ParseMode string `json:"parse_mode,omitempty"`
	// CaptionEntities is a JSON-serialized list of special entities
	// that appear in the message text,
	// which can be specified instead of parse_mode.
	CaptionEntities []entity.MessageEntity `json:"caption_entities,omitempty"`
	// ReplyMarkup as an additional interface options.
	ReplyMarkup entity.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// EditMessageMediaEnvelop is used to edit animation, audio, document, photo, or video messages.
type EditMessageMediaEnvelop struct {
	// ChatID is the id for the target chat or username.
	//
	// It is required if inline_message_id is not specified.
	ChatID string `json:"chat_id,omitempty"`
	// MessageID is the id of the message to edit.
	//
	// It is required if inline_message_id is not specified.
	MessageID int64 `json:"message_id,omitempty"`
	//InlineMessageID is the id of the inline message to edit.
	//
	// It is required if chat_id and message_id are not specified.
	InlineMessageID string `json:"inline_message_id,omitempty"`
	// Media is a JSON-serialized object for a new media content of the message.
	//
	// It is a required field.
	Media entity.InputMedia `json:"media,omitempty"`
	// ReplyMarkup as an additional interface options.
	ReplyMarkup entity.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// EditMessageLiveLocationEnvelop is used to edit live location messages.
type EditMessageLiveLocationEnvelop struct {
	// ChatID is the id for the target chat or username.
	//
	// It is required if inline_message_id is not specified.
	ChatID string `json:"chat_id,omitempty"`
	// MessageID is the id of the message to edit.
	//
	// It is required if inline_message_id is not specified.
	MessageID int64 `json:"message_id,omitempty"`
	//InlineMessageID is the id of the inline message to edit.
	//
	// It is required if chat_id and message_id are not specified.
	InlineMessageID string `json:"inline_message_id,omitempty"`
	// Latitude is the new latitude of the message.
	//
	// It is a required field.
	Latitude float64 `json:"latitude,omitempty"`
	// Longitude is the new longitude of the message.
	//
	// It is a required field.
	Longitude float64 `json:"longitude,omitempty"`
	// HorizontalAccuracy is the new horizontal accuracy of the message.
	HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`
	// Heading is the new direction in which the user is moving, in degrees.
	Heading int64 `json:"heading,omitempty"`
	// ProximityAlertRadius is the new maximum distance for proximity alerts
	// about approaching another chat member, in meters.
	ProximityAlertRadius int64 `json:"proximity_alert_radius,omitempty"`
	// ReplyMarkup as an additional interface options.
	ReplyMarkup entity.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// StopMessageLiveLocationEnvelop is used to stop updating a live location message.
type StopMessageLiveLocationEnvelop struct {
	// ChatID is the id for the target chat or username.
	//
	// It is required if inline_message_id is not specified.
	ChatID string `json:"chat_id,omitempty"`
	// MessageID is the id of the message to edit.
	//
	// It is required if inline_message_id is not specified.
	MessageID int64 `json:"message_id,omitempty"`
	//InlineMessageID is the id of the inline message to edit.
	//
	// It is required if chat_id and message_id are not specified.
	InlineMessageID string `json:"inline_message_id,omitempty"`
	// ReplyMarkup as an additional interface options.
	ReplyMarkup entity.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// EditMessageReplyMarkupEnvelop is used to edit only the reply markup of messages.
type EditMessageReplyMarkupEnvelop struct {
	// ChatID is the id for the target chat or username.
	//
	// It is required if inline_message_id is not specified.
	ChatID string `json:"chat_id,omitempty"`
	// MessageID is the id of the message to edit.
	//
	// It is required if inline_message_id is not specified.
	MessageID int64 `json:"message_id,omitempty"`
	//InlineMessageID is the id of the inline message to edit.
	//
	// It is required if chat_id and message_id are not specified.
	InlineMessageID string `json:"inline_message_id,omitempty"`
	// ReplyMarkup as an additional interface options.
	ReplyMarkup entity.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// StopPollEnvelop is used to stop a poll which was sent by the bot.
type StopPollEnvelop struct {
	// ChatID is the id for the target chat or username.
	//
	// It is required if inline_message_id is not specified.
	ChatID string `json:"chat_id,omitempty"`
	// MessageID is the id of the message to edit.
	//
	// It is required if inline_message_id is not specified.
	MessageID int64 `json:"message_id,omitempty"`
	// ReplyMarkup as an additional interface options.
	ReplyMarkup entity.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// DeleteMessageEnvelop is used to delete a message, including service messages,
type DeleteMessageEnvelop struct {
	// ChatID is the id for the target chat or username.
	//
	// It is required if inline_message_id is not specified.
	ChatID string `json:"chat_id,omitempty"`
	// MessageID is the id of the message to edit.
	//
	// It is required if inline_message_id is not specified.
	MessageID int64 `json:"message_id,omitempty"`
}
