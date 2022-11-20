package entity

// InlineQuery represents an incoming inline query.
// When the user sends an empty query, your bot could return some default or trending results.
type InlineQuery struct {
	// ID is a unique identifier for this query.
	//
	// It is a required field.
	ID string `json:"id,omitempty"`
	// From is the sender of the query.
	//
	// It is a required field.
	From *User `json:"user,omitempty"`
	// Query is the text of the query.
	//
	// It is a required field.
	Query string `json:"query,omitempty"`
	// Offset is offset of the results to be returned, can be controlled by the bot.
	//
	// It is a required field.
	Offset string `json:"offset,omitempty"`
	// ChatType is the type of the chat from which the inline query was sent.
	ChatType string `json:"chat_type,omitempty"`
	// Location is the sender location, only for bots that request user location
	Location *Location `json:"location,omitempty"`
}

// ChosenInlineResult represents a result of an inline query that was chosen by the user and sent to their chat partner.
type ChosenInlineResult struct {
	// ResultID is the unique identifier for the result that was chosen.
	//
	// It is a required field.
	ResultID string `json:"result_id,omitempty"`
	// From is the user that chose the result.
	//
	// It is a required field.
	From *User `json:"from,omitempty"`
	// Location is sender location, only for bots that require user location.
	Location *Location `json:"location,omitempty"`
	// InlineMessageID is the identifier of the sent inline message.
	// Available only if there is an InlineKeyboardMarkup attached to the message
	InlineMessageID string `json:"inline_message_id,omitempty"`
	// Query is the query that was used to obtain the result.
	Query string `json:"query,omitempty"`
}

// CallbackQuery represents an incoming callback query from a callback button in an inline keyboard.
//
// After the user presses a callback button,
// Telegram clients will display a progress bar until you call AnswerCallbackQuery.
type CallbackQuery struct {
	// ID is the unique identifier for this query.
	//
	// It is a required field.
	ID string `json:"id,omitempty"`
	// From is the sender of the query.
	//
	// It is a required field.
	From *User `json:"from,omitempty"`
	// Message is the message with the callback button that originated the query.
	Message *Message `json:"message,omitempty"`
	// InlineMessageID is the identifier of the message sent via the bot in inline mode, that originated the query.
	InlineMessageID string `json:"inline_message_id,omitempty"`
	// ChatInstance Global identifier,
	// uniquely corresponding to the chat to which the message with the callback button was sent.
	// Useful for high scores in games.
	ChatInstance string `json:"chat_instance,omitempty"`
	// Data is the data associated with the callback button.
	// Be aware that the message originated the query can contain no callback buttons with this data.
	Data string `json:"data,omitempty"`
	// GameShortName is the short name of a Game to be returned, serves as the unique identifier for the game.
	GameShortName string `json:"game_short_name,omitempty"`
}

// AnswerCallbackQueryEntity is the model used when sending AnswerCallbackQuery.
type AnswerCallbackQueryEntity struct {
	// CallbackQueryID is the unique identifier for the query to be answered.
	//
	// It is a required field.
	CallbackQueryID string `json:"callback_query_id"`
	// Text is the text of the notification.
	// If not specified, nothing will be shown to the user, 0-200 characters.
	Text string `json:"text,omitempty"`
	// ShowAlert If True, an alert will be shown by the client
	// instead of a notification at the top of the chat screen.
	// Defaults to false.
	ShowAlert bool `json:"show_alert,omitempty"`
	// URL is the url that will be opened by the user's client.
	URL string `json:"url,omitempty"`
	// CacheTime is the maximum amount of time in seconds
	// that the result of the callback query may be cached client-side.
	CacheTime int64 `json:"cache_time,omitempty"`
}
