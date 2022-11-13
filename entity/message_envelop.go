package entity

// MessageEnvelop holds the object that is used to send a new message
//
// TODO: incomplete
type MessageEnvelop struct {
	ChatID                   string `json:"chat_id,omitempty"`
	Text                     string `json:"text,omitempty"`
	ReplyToMessageID         int64  `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool   `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              struct {
		*InlineKeyboardMarkup
		*ReplyKeyboardMarkup
		// ReplyKeyboardRemove
		// ForceReply
	} `json:"reply_markup,omitempty"`
}
