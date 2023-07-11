package entity

// ChatAction is the type of chat action.
type ChatAction string

const (
	// ChatActionTyping is for typing action.
	ChatActionTyping ChatAction = "typing"
	// ChatActionUploadPhoto is for upload photo action.
	ChatActionUploadPhoto ChatAction = "upload_photo"
	// ChatActionRecordVideo is for record video action.
	ChatActionRecordVideo ChatAction = "record_video"
	// ChatActionUploadVideo is for upload video action.
	ChatActionUploadVideo ChatAction = "upload_video"
	// ChatActionRecordVoice is for record voice action.
	ChatActionRecordVoice ChatAction = "record_voice"
	// ChatActionUploadVoice is for upload voice action.
	ChatActionUploadVoice ChatAction = "upload_voice"
	// ChatActionUploadDocument is for upload document action.
	ChatActionUploadDocument ChatAction = "upload_document"
	// ChatActionChooseSticker is for choose sticker action.
	ChatActionChooseSticker ChatAction = "choose_sticker"
	// ChatActionFindLocation is for find location action.
	ChatActionFindLocation ChatAction = "find_location"
	// ChatActionRecordVideoNote is for record video note action.
	ChatActionRecordVideoNote ChatAction = "record_video_note"
	// ChatActionUploadVideoNote is for upload video note action.
	ChatActionUploadVideoNote ChatAction = "upload_video_note"
)

type BotCommandScopeType string

const (
	// BotCommandScopeDefault is the default scope.
	BotCommandScopeDefault BotCommandScopeType = "default"
	// BotCommandScopeAllPrivateChats is the all private chats scope.
	BotCommandScopeAllPrivateChats BotCommandScopeType = "all_private_chats"
	// BotCommandScopeAllGroupChats is the all group chats scope.
	BotCommandScopeAllGroupChats BotCommandScopeType = "all_group_chats"
	// BotCommandScopeAllChatAdministrators is the all chat administrators scope.
	BotCommandScopeAllChatAdministrators BotCommandScopeType = "all_chat_administrators"
	// BotCommandScopeChat is the chat scope.
	BotCommandScopeChat BotCommandScopeType = "chat"
	// BotCommandScopeChatAdministrators is the chat administrators scope.
	BotCommandScopeChatAdministrators BotCommandScopeType = "chat_administrators"
	// BotCommandScopeChatMember is the chat member scope.
	BotCommandScopeChatMember BotCommandScopeType = "chat_member"
)
