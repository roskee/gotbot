package envelop

import "github.com/roskee/gotbot/entity"

// SendInvoiceEnvelop is used to send invoices.
type SendInvoiceEnvelop struct {
	// ChatID is unique identifier for the target chat.
	//
	// It is a required field.
	ChatID int64 `json:"chat_id,omitempty"`
	// MessageThreadID is unique identifier of the target message thread (topic).
	MessageThreadID int64 `json:"message_thread_id,omitempty"`
	// Title is product name, 1-32 characters.
	//
	// It is a required field.
	Title string `json:"title,omitempty"`
	// Description is product description, 1-255 characters.
	//
	// It is a required field.
	Description string `json:"description,omitempty"`
	// Payload is bot-defined invoice payload, 1-128 bytes.
	//
	// It is a required field.
	Payload string `json:"payload,omitempty"`
	// ProviderToken is payments provider token, obtained via Botfather.
	//
	// It is a required field.
	ProviderToken string `json:"provider_token,omitempty"`
	// Currency is three-letter ISO 4217 currency code.
	//
	// It is a required field.
	Currency string `json:"currency,omitempty"`
	// Prices is price breakdown, a list of components
	// (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.).
	//
	// It is a required field.
	Prices []entity.LabeledPrice `json:"prices,omitempty"`
	// MaxTipAmount is the maximum accepted amount for tips
	// in the smallest units of the currency (integer, not float/double).
	MaxTipAmount int64 `json:"max_tip_amount,omitempty"`
	// SuggestedTipAmounts is a JSON-serialized array of suggested amounts of tip
	// in the smallest units of the currency (integer, not float/double).
	// At most 4 suggested tip amounts can be specified.
	// The suggested tip amounts must be positive, passed in a strictly increased order
	// and must not exceed max_tip_amount.
	SuggestedTipAmounts []int64 `json:"suggested_tip_amounts,omitempty"`
	// StartParameter is unique deep-linking parameter.
	// If left empty, forwarded copies of the sent message will have a Pay button,
	// allowing multiple users to pay directly from the forwarded message,
	// using the same invoice.
	StartParameter string `json:"start_parameter,omitempty"`
	// ProviderData is a JSON-serialized data about the invoice,
	// which will be shared with the payment provider.
	// A detailed description of required fields should be provided by the payment provider.
	// See https://core.telegram.org/bots/api#payments
	ProviderData string `json:"provider_data,omitempty"`
	// PhotoURL is URL of the product photo for the invoice.
	// Can be a photo of the goods or a marketing image for a service.
	// People like it better when they see what they are paying for.
	PhotoURL string `json:"photo_url,omitempty"`
	// PhotoSize is photo size in bytes.
	PhotoSize int64 `json:"photo_size,omitempty"`
	// PhotoWidth is photo width.
	PhotoWidth int64 `json:"photo_width,omitempty"`
	// PhotoHeight is photo height.
	PhotoHeight int64 `json:"photo_height,omitempty"`
	// NeedName is pass True, if you require the user's full name to complete the order.
	NeedName bool `json:"need_name,omitempty"`
	// NeedPhoneNumber is pass True, if you require the user's phone number to complete the order.
	NeedPhoneNumber bool `json:"need_phone_number,omitempty"`
	// NeedEmail is pass True, if you require the user's email address to complete the order.
	NeedEmail bool `json:"need_email,omitempty"`
	// NeedShippingAddress is pass True, if you require the user's shipping address to complete the order.
	NeedShippingAddress bool `json:"need_shipping_address,omitempty"`
	// SendPhoneNumberToProvider is pass True, if user's phone number should be sent to provider.
	SendPhoneNumberToProvider bool `json:"send_phone_number_to_provider,omitempty"`
	// SendEmailToProvider is pass True, if user's email address should be sent to provider.
	SendEmailToProvider bool `json:"send_email_to_provider,omitempty"`
	// IsFlexible is pass True, if the final price depends on the shipping method.
	IsFlexible bool `json:"is_flexible,omitempty"`
	// DisableNotification is sends the message silently.
	// Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// ProtectContent is pass True, if the content of the message needs to be protected
	// from forwarding and saving.
	ProtectContent bool `json:"protect_content,omitempty"`
	// ReplyToMessageID is if the message is a reply, ID of the original message.
	ReplyToMessageID int64 `json:"reply_to_message_id,omitempty"`
	// AllowSendingWithoutReply is pass True, if the message should be sent
	// even if the specified replied-to message is not found.
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	// ReplyMarkup is additional interface options.
	ReplyMarkup *entity.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}
