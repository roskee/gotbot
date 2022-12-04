package entity

// Dice represents an animated emoji that displays a random value.
type Dice struct {
	// Emoji is emoji on which the dice throw animation is based
	//
	// It is required field
	Emoji string `json:"emoji"`
	// Value is value of the dice, 1-6 for “🎲”, “🎯” and “🎳” base emoji, 1-5 for “🏀” and “⚽” base emoji, 1-64 for “🎰” base emoji
	//
	// It is required field, but optional on sending dice using sendDice method
	Value int64 `json:"value"`
}
