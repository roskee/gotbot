package entity

// KeyboardButtonPollType represents a type of poll, which is allowed to be created and sent when the corresponding button is pressed.
type KeyboardButtonPollType struct {
	// Type is the type of this poll
	//
	// If it is set to `quiz`, the user will be allowed to create only polls in the quiz mode.
	// If it is set to `regular`, only regular polls will be allowed.
	// Otherwise, the user will be allowed to create a poll of any type
	Type string `json:"type,omitempty"`
}

// WebAppInfo describes a web app
type WebAppInfo struct {
	// URL is an HTTPS url of a Web App to be opened
	//
	// It is a required field
	URL string `json:"url,omitempty"`
}

// KeyboardButton represents one button of the reply keyboard
type KeyboardButton struct {
	// Text is text of the button
	//
	// It is a required field
	Text string `json:"text,omitempty"`
	// RequestContact can be true to attach the user's phone number when this button is pressed
	RequestContact bool `json:"request_contact,omitempty"`
	// RequestLocation can be true to attach the user's location when this button is pressed
	RequestLocation bool `json:"request_location,omitempty"`
	// RequestPoll if specified, the user will be asked to create a poll and send it to the bot when the button is pressed
	RequestPoll *KeyboardButtonPollType `json:"request_poll,omitempty,omitempty"`
	// WebApp can be used to launch the specified web app
	WebApp *WebAppInfo `json:"web_app,omitempty,omitempty"`
}

// ReplyKeyboardMarkup represents a custom keyboard with reply options
type ReplyKeyboardMarkup struct {
	// Keyboard is an array of button rows, each represented by an array of KeyboardButton objects
	//
	// it is a required field
	Keyboard [][]KeyboardButton `json:"keyboard,omitempty"`
	// ResizeKeyboard requests clients to resize the keyboard vertically for optimal fit
	ResizeKeyboard bool `json:"resize_keyboard,omitempty"`
	// OneTimeKeyboard requests clients to hide the keyboard as soon as it's been used
	OneTimeKeyboard bool `json:"one_time_keyboard,omitempty"`
	// InputFieldPlaceholder is the placeholder to be shown in the input field when the keyboard is active
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`
	// Selective use this parameter if you want to show the keyboard to specific users only
	Selective bool `json:"selective,omitempty"`
}

// InlineKeyboardMarkup represents an inline keyboard that appears right next to the message it belongs to.
type InlineKeyboardMarkup struct {
	// InlineKeyboard is an array of button rows, each represented by an Array of InlineKeyboardButton objects
	//
	// It is a required field
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard,omitempty"`
}

// InlineKeyboardButton represents one button of an inline keyboard.
// You must use exactly one of the optional fields.
type InlineKeyboardButton struct {
	// Text is the label text on the button.
	//
	// It is a required field.
	Text string `json:"text,omitempty"`
	// URL is an HTTP or tg:// url to be opened when the button is pressed.
	URL string `json:"url,omitempty"`
	// CallbackData is a data to be sent in a callback query to the bot when button is pressed.
	CallbackData string `json:"callback_data,omitempty"`
	// WebApp is a description of the Web App that will be launched when the user presses the button.
	WebApp *WebAppInfo `json:"web_app,omitempty,omitempty"`
	// LoginUrl is an HTTPS url used to automatically authorize the user.
	LoginUrl *LoginUrl `json:"login_url,omitempty,omitempty"`
	// SwitchInlineQuery if set, pressing the button will prompt the user to select one of their chats,
	// open that chat and insert the bot's username and the specified inline query in the input field.
	// May be empty, in which case just the bot's username will be inserted.
	SwitchInlineQuery string `json:"switch_inline_query,omitempty"`
	// SwitchInlineQueryCurrentChat If set, pressing the button will insert the bot's username
	// and the specified inline query in the current chat's input field.
	// May be empty, in which case only the bot's username will be inserted.
	SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat,omitempty"`
	// CallbackGame is the description of the game that will be launched when the user presses the button.
	//
	// NOTE: This type of button must always be the first button in the first row.
	CallbackGame string `json:"callback_game,omitempty"`
	// Pay can be true to send a pay button.
	//
	// NOTE: This type of button must always be the first button in the first row
	// and can only be used in invoice messages.
	Pay bool `json:"pay,omitempty"`
}

// LoginUrl represents a parameter of the inline keyboard button used to automatically authorize a user.
type LoginUrl struct {
	// URL is an HTTPS URL to be opened with user authorization data added to the query string when the button is pressed.
	//
	// it is a required field
	URL string `json:"url,omitempty"`
	// ForwardText new text of the button in forwarded messages.
	ForwardText string `json:"forward_text,omitempty"`
	// BotUsername is username of a bot, which will be used for user authorization.
	BotUsername string `json:"bot_username,omitempty"`
	// RequestWriteAccess can be true to request permission for your bot to send messages to the user.
	RequestWriteAccess bool `json:"request_write_access,omitempty"`
}

// CallbackGame is a placeholder, currently holds no information.
type CallbackGame struct {
	// UserID is the user identifier.
	//
	// it is a required field
	UserID int64 `json:"user_id,omitempty"`
	// Score is the new score. must be non-negative.
	//
	// it is a required field
	Score int64 `json:"score,omitempty"`
	// Force can be true if the high score is allowed to decrease.
	Force bool `json:"force,omitempty"`
	// DisableEditMessage can be true if the game message should not be automatically edited
	// to include the current scoreboard
	DisableEditMessage bool `json:"disable_edit_message,omitempty"`
	// ChatID is Unique identifier for the target chat
	//
	// it is required if InlineMessageID is not specified.
	ChatID int64 `json:"chat_id,omitempty"`
	// MessageID is identifier of the sent message.
	//
	// it is required if InlineMessageID is not specified
	MessageID int64 `json:"message_id,omitempty"`
	// InlineMessageID is identifier of the inline message
	//
	// it is required if ChatID and MessageID are not specified.
	InlineMessageID int64 `json:"inline_message_id,omitempty"`
}
