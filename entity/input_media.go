package entity

// InputMedia represents the content of a media message to be sent.
// It should be one of
//
// * InputMediaAnimation
//
// * InputMediaDocument
//
// * InputMediaAudio
//
// * InputMediaPhoto
//
// * InputMediaVideo
//
// TODO: incomplete
type InputMedia struct {
	*InputMediaPhoto
}

// InputMediaPhoto represents a photo to be sent.
type InputMediaPhoto struct {
	// Type is type of the result, must be `photo`.
	//
	// It is a required field
	Type string `json:"type,omitempty"`
	// Media file to send.
	// Pass a file_id to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet,
	// or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data
	// under <file_attach_name> name.
	//
	// It is a required field
	Media string `json:"media,omitempty"`
	// Caption is Caption of the photo to be sent, 0-1024 characters after entities parsing.
	Caption string `json:"caption,omitempty"`
	// ParseMode is Mode for parsing entities in the photo caption.
	ParseMode string `json:"parse_mode,omitempty"`
	// CaptionEntities List of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
}
