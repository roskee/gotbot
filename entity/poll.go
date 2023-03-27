package entity

// Poll contains information about a poll.
type Poll struct {
	// ID is unique poll identifier
	//
	// It is a required field
	ID string `json:"id"`
	// Question is poll question, 1-300 characters
	//
	// It is a required field
	Question string `json:"question"`
	// Options is  list of poll options
	//
	// It is a required field
	Options []*PollOption `json:"options"`
	// TotalVoterCount is total number of users that voted in the poll
	//
	// It is a required field
	TotalVoterCount int64 `json:"total_voter_count"`
	// IsClosed is true, if the poll is closed
	//
	// It is a required field
	IsClosed bool `json:"is_closed"`
	// IsAnonymous is true, if the poll is anonymous
	//
	// It is a required field
	IsAnonymous bool `json:"is_anonymous"`
	// Type	is poll type, currently can be “regular” or “quiz”
	//
	// It is a required field
	Type string `json:"type"`
	// AllowsMultipleAnswers is true, if the poll allows multiple answers
	//
	// It is a required field
	AllowsMultipleAnswers bool `json:"allows_multiple_answers"`
	// CorrectOptionID is  0-based identifier of the correct answer option.
	CorrectOptionID int64 `json:"correct_option_id"`
	// Explanation is text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a quiz-style poll, 0-200 characters
	Explanation string `json:"explanation"`
	// ExplanationEntities is special entities like usernames, URLs, bot commands, etc.
	ExplanationEntities []*MessageEntity `json:"explanation_entities"`
	// OpenPeriod is amount of time in seconds the poll will be active after creation
	OpenPeriod int64 `json:"open_period"`
	// CloseDate is point in time (Unix timestamp) when the poll will be automatically closed
	CloseDate int64 `json:"close_date"`
}

// PollOption contains information about one answer option in a poll.
type PollOption struct {
	// Text is option text, 1-100 characters
	//
	// It is q required field
	Text string `json:"text"`
	// VoterCount is  number of users that voted for this option
	//
	// It is q required field
	VoterCount int64 `json:"voter_count"`
}
