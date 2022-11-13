package entity

import (
	"encoding/json"
	"log"
)

// Update holds values from an update sent by the telegram server.
// // incomplete //
type Update struct {
	UpdateID int64    `json:"update_id"`
	Message  *Message `json:"message"`
}

// FromJSONBody modifies the current Update with matching values in body
func (u *Update) FromJSONBody(body []byte) {
	err := json.Unmarshal(body, u)
	if err != nil {
		log.Printf("error while unmarshalling update: %+v", err)
	}
}
