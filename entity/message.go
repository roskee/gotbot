package entity

import "strings"

// Message holds a message object the telegram server sends.
//
// TODO: incomplete
type Message struct {
	// MessageID is a unique message identifier inside this chat.
	//
	// it is a required field.
	MessageID int64 `json:"message_id,omitempty"`
	// MessageThreadID is a unique identifier of a message thread to which the message belongs.
	// it is for supergroups only
	MessageThreadID int64 `json:"message_thread_id,omitempty"`
	// From is the sender of the message.
	// It is empty for messages sent to channels.
	From *User `json:"from,omitempty"`
	// SenderChat is the sender of the message, sent on behalf of a chat.
	SenderChat *Chat `json:"sender_chat,omitempty"`
	// Date is the date the message was sent in Unix time.
	//
	// It is a required field
	Date int64 `json:"date,omitempty"`
	// Chat is the conversation this message belongs to.
	//
	// It is a required field
	Chat *Chat `json:"chat,omitempty"`
	// ForwardFrom is, for forwarded messages, sender of the original message.
	ForwardFrom *User `json:"forward_from,omitempty"`
	// ForwardFromChat is,  For messages forwarded from channels or from anonymous administrators,
	// information about the original sender chat
	ForwardFromChat *Chat    `json:"forward_from_chat,omitempty"`
	ReplyToMessage  *Message `json:"reply_to_message,omitempty"`
	Text            string   `json:"text,omitempty"`
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

// MessageEntity  represents one special entity in a text message.
// For example, hashtags, usernames, URLs, etc.
type MessageEntity struct {
	// Type is Type of the entity.
	// Currently, can be “mention” (@username), “hashtag” (#hashtag), “cashtag” ($USD),
	// “bot_command” (/start@jobs_bot), “url” (https://telegram.org), “email” (do-not-reply@telegram.org),
	// “phone_number” (+1-212-555-0123), “bold” (bold text), “italic” (italic text),
	// “underline” (underlined text), “strikethrough” (strikethrough text), “spoiler” (spoiler message),
	// “code” (monowidth string), “pre” (monowidth block), “text_link” (for clickable text URLs),
	// “text_mention” (for users without usernames), “custom_emoji” (for inline custom emoji stickers)
	//
	// It is a required field
	Type string `json:"type,omitempty"`
	// Offset is offset in UTF-16 code units to the start of the entity.
	//
	// It is a required field
	Offset int64 `json:"offset,omitempty"`
	// Length is length of the entity in UTF-16 code units.
	//
	// It is a required field
	Length int64 `json:"length,omitempty"`
	// URL is, for “text_link” only, URL that will be opened after user taps on the text.
	URL string `json:"url,omitempty"`
	// User is, for “text_mention” only, the mentioned user.
	User *User `json:"user,omitempty"`
	// Language is, for “pre” only, the programming language of the entity text.
	Language string `json:"language,omitempty"`
	// CustomEmojiID is, for “custom_emoji” only, unique identifier of the custom emoji.
	CustomEmojiID string `json:"custom_emoji_id,omitempty"`
}
