package entity

// PhotoSize represents one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {
	// FileID is identifier for this file, which can be used to download or reuse the file.
	//
	// It is a required field.
	FileID string `json:"file_id,omitempty"`
	// FileUniqueID is a unique identifier for this file,
	// which is supposed to be the same over time and for different bots.
	//
	// It is a required field.
	FileUniqueID string `json:"file_unique_id,omitempty"`
	// Width is the photo's width.
	//
	// It is a required field.
	Width int64 `json:"width,omitempty"`
	// Height is the photo's height.
	//
	// It is a required field.
	Height int64 `json:"height,omitempty"`
	// FileSize is the file size in bytes.
	FileSize int64 `json:"file_size,omitempty"`
}

// UserProfilePhotos represent a user's profile pictures.
type UserProfilePhotos struct {
	// TotalCount is the total number of profile pictures the target user has.
	TotalCount int64 `json:"total_count,omitempty"`
	// Photos is requested profile pictures (in up to 4 sizes each).
	Photos [][]PhotoSize `json:"photos,omitempty"`
}
