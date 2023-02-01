package entity

// InputMedia represents the content of a media message to be sent.
// It can be one of the f/f types from the telegram api.
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
// Note that only a subset of the fields are supported for every type.
type InputMedia struct {
	// Type is type of the result, must be one of,
	// `photo`, `video`, `animation`, `audio`, `document`,
	//
	// It is a required field
	Type string `json:"type,omitempty"`
	// Media is the media to send
	//
	// pass a file_id to send a media that exists on the Telegram servers,
	// pass an HTTP URL for Telegram to get a file from the Internet,
	// or pass “attach://<file_attach_name>” to upload a new one
	// using multipart/form-data under <file_attach_name> name
	//
	// It is a required field
	Media string `json:"media,omitempty"`
	// Caption of the media to be sent.
	// 0-1024 characters.
	Caption string `json:"caption,omitempty"`
	// ParseMode is mode for parsing entities in the caption.
	ParseMode string `json:"parse_mode,omitempty"`
	// CaptionEntities is the list of special entities that appear in the caption,
	// which can be specified instead of parse_mode.
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	// HasSpoiler can be true to cover the media with spoiler animation
	HasSpoiler bool `json:"has_spoiler,omitempty"`
	// Thumb is the thumbnail of the media to be sent.
	Thumb any `json:"thumb,omitempty"`
	// Width of the media.
	Width float64 `json:"width,omitempty"`
	// Height of the media.
	Height float64 `json:"height,omitempty"`
	// Duration of the medis.
	Duration float64 `json:"duration,omitempty"`
	// SupportsStreaming can be true if the media is suitable for streaming.
	SupportsStreaming bool `json:"supports_streaming,omitempty"`
	// Performer of the audio.
	Performer string `json:"performer,omitempty"`
	// Title of the audio.
	Title string `json:"title,omitempty"`
	// DisableContentTypeDetection disables server side file type detection.
	DisableContentTypeDetection bool `json:"disable_content_type_detection,omitempty"`
}
