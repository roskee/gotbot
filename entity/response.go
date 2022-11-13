package entity

// Response is a wrapper object that the telegram server sends for every request
type Response struct {
	OK          bool               `json:"ok"`
	Description string             `json:"description"`
	ErrorCode   int64              `json:"error_code"`
	Parameters  ResponseParameters `json:"parameters"`
	Result      interface{}        `json:"result"`
}

// ResponseParameters contains error details for a failed response message
type ResponseParameters struct {
	MigrateToChatID int64 `json:"migrate_to_chat_id"`
	RetryAfter      int64 `json:"retry_after"`
}
