package envelop

// GetUserProfilePhotos represents a request to get a list of profile pictures for a user.
type GetUserProfilePhotos struct {
	// UserID is unique identifier of the target user.
	//
	// It is a required field.
	UserID int64 `json:"user_id,omitempty"`
	// Offset is sequential number of the first photo to be returned.
	// By default, all photos are returned.
	Offset int64 `json:"offset,omitempty"`
	// Limit limits the number of photos to be retrieved.
	// Values between 1—100 are accepted. Defaults to 100.
	Limit int64 `json:"limit,omitempty"`
}
