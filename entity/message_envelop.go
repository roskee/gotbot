package entity

// MessageEnvelop holds the object that is used to send a new message
type MessageEnvelop struct {
	ChatID                   string `json:"chat_id"`
	Text                     string `json:"text"`
	ReplyToMessageID         int64  `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool   `json:"allow_sending_without_reply"`
}
