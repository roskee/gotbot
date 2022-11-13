package entity

import "strings"

// Message holds a message object the telegram server sends.
//
// TODO: incomplete
type Message struct {
	MessageID      int64    `json:"message_id,omitempty"`
	From           User     `json:"from,omitempty"`
	Date           int64    `json:"date,omitempty"`
	ForwardFrom    User     `json:"forward_from,omitempty"`
	ReplyToMessage *Message `json:"reply_to_message,omitempty"`
	Text           string   `json:"text,omitempty"`
}

// GetCommand checks if this message has Text starting with '/'.
// if so, it returns the immediate word attached to it
func (m *Message) GetCommand() string {
	if len(m.Text) == 0 || m.Text[0] != '/' {
		return ""
	}
	index := strings.Index(m.Text, " ")
	if index == -1 {
		index = len(m.Text)
	}
	return m.Text[1:index]
}
