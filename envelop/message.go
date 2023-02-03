package envelop

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
