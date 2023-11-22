package entity

// LabeledPrice represents a portion of the price for goods or services.
type LabeledPrice struct {
	// Label is a label for this portion of the price.
	Label string `json:"label,omitempty"`
	// Amount is price of the product in the smallest units of the currency
	// (integer, not float/double).
	// For example, for a price of US$ 1.45 pass amount = 145.
	Amount int64 `json:"amount,omitempty"`
}
