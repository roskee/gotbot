package envelop

import "github.com/roskee/gotbot/entity"

// DeleteMyCommandsEnvelop represents a request to delete all currently registered commands
type DeleteMyCommandsEnvelop struct {
	// Scope is the scope to which the commands are relevant.
	// Defaults to entity.BotCommandScopeDefault.
	Scope *entity.BotCommandScope `json:"scope,omitempty"`
	// LanguageCode is a two-letter ISO 639-1 language code.
	// If empty, commands will be applied to all users from the given scope,
	// for whose language there are no dedicated commands
	LanguageCode string `json:"language_code,omitempty"`
}
