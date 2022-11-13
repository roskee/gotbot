package entity

import (
	"encoding/json"
	"log"
)

// User holds a user response that is sent from the telegram server
type User struct {
	ID                      int64  `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name"`
	UserName                string `json:"username"`
	LanguageCode            string `json:"language_code"`
	CanJoinGroups           bool   `json:"can_join_groups"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries"`
}

// FromJSONBody modifies the current user with matching values in body
func (u *User) FromJSONBody(body []byte) {
	err := json.Unmarshal(body, u)
	if err != nil {
		log.Printf("Error while unmarshalling user object: %v", err)
	}
}
