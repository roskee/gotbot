package entity

// Audio represents an audio file to be treated as music by the Telegram clients.
type Audio struct {
	// FileID is identifier for this file, which can be used to download or reuse the file.
	//
	// It is a required field
	FileID string `json:"file_id,omitempty"`
	// FileUniqueID is unique identifier for this file, which is supposed to be the same over time and for different bots.
	//
	// It is a required field
	FileUniqueID string `json:"file_unique_id,omitempty"`
	// Duration is duration of the audio in seconds as defined by sender.
	//
	// It is a required field
	Duration int64 `json:"duration,omitempty"`
	// Performer is performer of the audio as defined by sender or by audio tags.
	Performer string `json:"performer,omitempty"`
	// Title is title of the audio as defined by sender or by audio tags.
	Title string `json:"title,omitempty"`
	// FileName is original flename as defined by sender.
	FileName string `json:"file_name,omitempty"`
	// MimeType is the MIME type of the file as defined by sender.
	MimeType string `json:"mime_type,omitempty"`
	// FileSize is the file size in bytes.
	FileSize int64 `json:"file_size,omitempty"`
	// Thumb is thumbnail of the album cover to which the music file belongs.
	Thumb *PhotoSize `json:"thumb,omitempty"`
}

// Documet represents a general file
// (as opposed to photos, voice messages and audio files).
type Document struct {
	// FileID is identifier for this file, which can be used to download or reuse the file.
	//
	// It is a required field
	FileID string `json:"file_id,omitempty"`
	// FileUniqueID is unique identifier for this file, which is supposed to be the same over time and for different bots.
	//
	// It is a required field
	FileUniqueID string `json:"file_unique_id,omitempty"`
	// Thumb is document thumbnail as defined by sender.
	Thumb *PhotoSize `json:"thumb,omitempty"`
	// FileName is original animation filename as defined by sender.
	FileName string `json:"file_name,omitempty"`
	// MimeType is the MIME type of the file as defined by sender.
	MimeType string `json:"mime_type,omitempty"`
	// FileSize is the file size in bytes.
	FileSize int64 `json:"file_size,omitempty"`
}

// Animation represents an animation file (GIF or H.264/MPEG-4 AVC video without sound).
type Animation struct {
	// FileID is identifier for this file, which can be used to download or reuse the file.
	//
	// It is a required field
	FileID string `json:"file_id,omitempty"`
	// FileUniqueID is unique identifier for this file, which is supposed to be the same over time and for different bots.
	//
	// It is a required field
	FileUniqueID string `json:"file_unique_id,omitempty"`
	// Width is video width as defined by sender.
	//
	// It is a required field
	Width int64 `json:"width,omitempty"`
	// Height is video height as defined by sender.
	//
	// It is a required field
	Height int64 `json:"height,omitempty"`
	// Duration is duration of the video in seconds as defined by sender.
	//
	// It is a required field
	Duration int64 `json:"duration,omitempty"`
	// Thumb is animation thumbnail as defined by sender.
	Thumb *PhotoSize `json:"thumb,omitempty"`
	// FileName is original animation filename as defined by sender.
	FileName string `json:"file_name,omitempty"`
	// MimeType is the MIME type of the file as defined by sender.
	MimeType string `json:"mime_type,omitempty"`
	// FileSize is the file size in bytes.
	FileSize int64 `json:"file_size,omitempty"`
}

// Sticker represents a sticker.
type Sticker struct {
	// FileID is identifier for this file, which can be used to download or reuse the file.
	//
	// It is a required field
	FileID string `json:"file_id,omitempty"`
	// FileUniqueID is unique identifier for this file, which is supposed to be the same over time and for different bots.
	//
	// It is a required field
	FileUniqueID string `json:"file_unique_id,omitempty"`
	// Type is type of the sticker.
	// currently one of “regular”, “mask”, “custom_emoji”.
	//
	// It is a required field
	Type string `json:"type,omitempty"`
	// Width is the sticker's width.
	//
	// It is a required field
	Width int64 `json:"width,omitempty"`
	// Height is the sticker's height.
	//
	// It is a required field
	Height int64 `json:"height,omitempty"`
	// IsAnimated is true, if the sticker is animated.
	//
	// It is a required field
	IsAnimated bool `json:"is_animated,omitempty"`
	// IsVideo is true, if the sticker is a video sticker.
	//
	// It is a required field
	IsVideo bool `json:"is_video,omitempty"`
	// Thumb is sticker's thumbnail in the .WEBP or .JPG format.
	Thumb *PhotoSize `json:"thumb,omitempty"`
	// Emoji is the emoji associated with the sticker.
	Emoji string `json:"emoji,omitempty"`
	// SetName is the name of the sticker set to which the sticker belongs.
	SetName string `json:"set_name,omitempty"`
	// PremiumAnimation is, for premium regular stickers, premium animation for the sticker.
	PremiumAnimation *File `json:"premium_animation,omitempty"`
	// MaskPosition is, for mask stickers, the position where the mask should be placed.
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	// CustomEmojiID is, for custom emoji stickers, unique identifier of the custom emoji.
	CustomEmojiID string `json:"custom_emoji_id,omitempty"`
	// FileSize is the file size in bytes.
	FileSize int64 `json:"file_size,omitempty"`
}

// Video represents a video file.
type Video struct {
	// FileID is identifier for this file, which can be used to download or reuse the file.
	//
	// It is a required field
	FileID string `json:"file_id,omitempty"`
	// FileUniqueID is unique identifier for this file, which is supposed to be the same over time and for different bots.
	//
	// It is a required field
	FileUniqueID string `json:"file_unique_id,omitempty"`
	// Width is video width as defined by sender.
	//
	// It is a required field
	Width int64 `json:"width,omitempty"`
	// Height is video height as defined by sender.
	//
	// It is a required field
	Height int64 `json:"height,omitempty"`
	// Duration is duration of the video in seconds as defined by sender.
	//
	// It is a required field
	Duration int64 `json:"duration,omitempty"`
	// Thumb is video thumbnail
	Thumb *PhotoSize `json:"thumb,omitempty"`
	// FileName is original animation filename as defined by sender.
	FileName string `json:"file_name,omitempty"`
	// MimeType is the MIME type of the file as defined by sender.
	MimeType string `json:"mime_type,omitempty"`
	// FileSize is the file size in bytes.
	FileSize int64 `json:"file_size,omitempty"`
}

// VideoNote represents a video message (available in Telegram apps as of v.4.0).
type VideoNote struct {
	// FileID is identifier for this file, which can be used to download or reuse the file.
	//
	// It is a required field
	FileID string `json:"file_id,omitempty"`
	// FileUniqueID is unique identifier for this file, which is supposed to be the same over time and for different bots.
	//
	// It is a required field
	FileUniqueID string `json:"file_unique_id,omitempty"`
	// Length is video width and height (diameter of the video message) as defined by sender.
	Length int64 `json:"length,omitempty"`
	// Duration is duration of the video in seconds as defined by sender.
	//
	// It is a required field
	Duration int64 `json:"duration,omitempty"`
	// Thumb is video thumbnail
	Thumb *PhotoSize `json:"thumb,omitempty"`
	// FileSize is the file size in bytes.
	FileSize int64 `json:"file_size,omitempty"`
}

// Voice
type Voice struct {
	// FileID is identifier for this file, which can be used to download or reuse the file.
	//
	// It is a required field
	FileID string `json:"file_id,omitempty"`
	// FileUniqueID is unique identifier for this file, which is supposed to be the same over time and for different bots.
	//
	// It is a required field
	FileUniqueID string `json:"file_unique_id,omitempty"`
	// Duration is duration of the audio in seconds as defined by sender.
	//
	// It is a required field
	Duration int64 `json:"duration,omitempty"`
	// MimeType is the MIME type of the file as defined by sender.
	MimeType string `json:"mime_type,omitempty"`
	// FileSize is the file size in bytes.
	FileSize int64 `json:"file_size,omitempty"`
}

// File represents a file ready to be downloaded.
type File struct {
	// FileID is identifier for this file, which can be used to download or reuse the file.
	//
	// It is a required field
	FileID string `json:"file_id,omitempty"`
	// FileUniqueID is unique identifier for this file, which is supposed to be the same over time and for different bots.
	//
	// It is a required field
	FileUniqueID string `json:"file_unique_id,omitempty"`
	// FileSize is the file size in bytes.
	FileSize int64 `json:"file_size,omitempty"`
	// FilePath is the file path
	FilePath string `json:"file_path,omitempty"`
}

// MaskPosition describes the position on faces where a mask should be placed by default..
type MaskPosition struct {
	// Point is the part of the face relative to which the mask should be placed.
	// One of “forehead”, “eyes”, “mouth”, or “chin”.
	Point string `json:"point,omitempty"`
	// XShift is shift by X-axis measured in widths of the mask scaled to the face size, from left to right.
	XShift float64 `json:"x_shift,omitempty"`
	// YShift is shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom.
	YShift float64 `json:"y_shift,omitempty"`
	// Scale is the Mask scaling coefficient.
	Scale float64 `json:"scale,omitempty"`
}
