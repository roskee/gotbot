package entity

// Response is a wrapper object that the telegram server sends for every request
type Response struct {
	OK          bool                `json:"ok,omitempty"`
	Description string              `json:"description,omitempty"`
	ErrorCode   int64               `json:"error_code,omitempty"`
	Parameters  *ResponseParameters `json:"parameters,omitempty"`
	Result      interface{}         `json:"result,omitempty"`
}

// ResponseParameters contains error details for a failed response message
type ResponseParameters struct {
	MigrateToChatID int64 `json:"migrate_to_chat_id,omitempty"`
	RetryAfter      int64 `json:"retry_after,omitempty"`
}
