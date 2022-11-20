package entity

import (
	"encoding/json"
	"log"
)

// Webhook contains options to set up a webhook for updates.
//
// TODO: Incomplete:
type Webhook struct {
	// URL is the https url to send updates to.
	// Use an empty string to remove webhook integration.
	//
	// It is a required field
	URL string `json:"url,omitempty"`
	// IPAddress is a fixed IP address which will be used to send webhook requests
	// instead of the IP address resolved through DNS.
	IPAddress string `json:"ip_address,omitempty"`
	// MaxConnections is the maximum allowed number of simultaneous HTTPS connections
	// to the webhook for update delivery, 1-100.
	// Defaults to 40.
	MaxConnections int64 `json:"max_connections,omitempty"`
	// AllowedUpdates the list of the update types you want your bot to receive.
	//
	// eg: UpdateMessage, UpdateEditedMessage, etc
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
	// DropPendingUpdates can be true to drop all pending updates.
	DropPendingUpdates bool `json:"drop_pending_updates,omitempty"`
	// SecretToken is a secret token to be sent in a header “X-Telegram-Bot-Api-Secret-Token” in every webhook request,
	// 1-256 characters. Only characters A-Z, a-z, 0-9, _ and - are allowed.
	SecretToken string `json:"secret_token,omitempty"`
}

// FromJSONBody modifies the current Webhook with matching values in body
func (u *Webhook) FromJSONBody(body []byte) {
	err := json.Unmarshal(body, u)
	if err != nil {
		log.Printf("error while unmarshalling webhook: %+v", err)
	}
}
