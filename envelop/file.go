package envelop

// GetFile contains the identifier of the file to get info about.
type GetFile struct {
	// FileID is the file identifier
	//
	// It is a required field.
	FileID string `json:"file_id,omitempty"`
}
