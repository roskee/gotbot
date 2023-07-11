package entity

import (
	"encoding/json"
	"log"
)

// Commands is a wrapper around commands to be sent to the telegram server
type Commands struct {
	Commands     []Command `json:"commands"`
	LanguageCode string    `json:"language_code"`
}

// Command holds the command object the telegram server sends
type Command struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

// FromJSONBody modifies the current Command with matching values in body
func (c *Command) FromJSONBody(body []byte) {
	err := json.Unmarshal(body, c)
	if err != nil {
		log.Printf("error while unmarshalling command: %+v", err)
	}
}

// ToJSONBody returns the current Command as a serialized json
func (c *Command) ToJSONBody() []byte {
	bytes, err := json.Marshal(c)
	if err != nil {
		log.Printf("error while marshalling command: %+v", err)
	}
	return bytes
}

// BotCommandScope describes the scope of bot commands,
type BotCommandScope struct {
	// Type is the type of scope.
	//
	// It is a required field.
	Type BotCommandScopeType `json:"type,omitempty"`
	// ChatID is the unique identifier for the target chat
	// or username of the target supergroup.
	//
	// It is a required field if Type is chat, chat_administrators or chat_member.
	ChatID string `json:"chat_id,omitempty"`
	// UserID is the unique identifier of the target user.
	//
	// It is a required field if Type is chat_member.
	UserID int64 `json:"user_id,omitempty"`
}
