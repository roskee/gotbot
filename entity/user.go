package entity

import (
	"encoding/json"
	"log"
)

// User holds a user response that is sent from the telegram server
type User struct {
	ID                      int64  `json:"id,omitempty"`
	IsBot                   bool   `json:"is_bot,omitempty"`
	FirstName               string `json:"first_name,omitempty"`
	LastName                string `json:"last_name,omitempty"`
	UserName                string `json:"username,omitempty"`
	LanguageCode            string `json:"language_code,omitempty"`
	CanJoinGroups           bool   `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages,omitempty"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries,omitempty"`
}

// FromJSONBody modifies the current user with matching values in body
func (u *User) FromJSONBody(body []byte) {
	err := json.Unmarshal(body, u)
	if err != nil {
		log.Printf("Error while unmarshalling user object: %v", err)
	}
}
