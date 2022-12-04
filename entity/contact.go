package entity

//Contact represents a phone contact.
type Contact struct {
	// PhoneNumber is Contact's phone number
	//
	// It is a required field
	PhoneNumber string `json:"phone_number"`
	// FirstName is Contact's first name
	//
	// It is a required field
	FirstName string `json:"first_name"`
	// LastName is Contact's last name
	LastName string `json:"last_name"`
	// UserID is Contact's user identifier in Telegram
	UserID string `json:"user_id"`
	// Vcard is Additional data about the contact in the form of a vCard
	Vcard string `json:"vcard"`
}
