package envelop

// SetChatAdministratorCustomTitle is the request body for setting custom title for administrators
type SetChatAdministratorCustomTitle struct {
	// ChatID of the target chat or username of the target super group.
	ChatID string `json:"chat_id,omitempty"`
	// UserID of the target user.
	UserID int `json:"user_id,omitempty"`
	// CustomTitle is the new custom title to set.
	CustomTitle string `json:"custom_title,omitempty"`
}
