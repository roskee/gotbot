package entity

// Chat represents a chat.
type Chat struct {
	// ID is a unique identifier for this chat.
	//
	// It is a required field
	ID int64 `json:"id,omitempty"`
	// Type is the type of chat, can be either “private”, “group”, “supergroup” or “channel”.
	//
	// It is a required field
	Type string `json:"type,omitempty"`
	// Title is the title for groups, super-groups and channels.
	Title string `json:"title,omitempty"`
	// Username is the username, for private chats, supergroups and channels if available.
	Username string `json:"username,omitempty"`
	// FirstName is the first name of the other party in a private chat.
	FirstName string `json:"first_name,omitempty"`
	// LastName is the last name of the other party in a private chat.
	LastName string `json:"last_name,omitempty"`
	// IsForum is true, if the supergroup chat is a forum (has topics enabled).
	IsForum bool `json:"is_forum,omitempty"`
	// Photo is the chat photo. Returned only in GetChat.
	Photo *ChatPhoto `json:"photo,omitempty"`
	// ActiveUsernames is if non-empty, the list of all active chat usernames;
	// for private chats, supergroups and channels.
	// Returned only in GetChat.
	ActiveUsernames []string `json:"active_usernames,omitempty"`
	// EmojiStatusCustomEmojiID is the custom emoji identifier of emoji status of the
	// other party in a private chat.
	// Returned only in GetChat.
	EmojiStatusCustomEmojiID string `json:"emoji_status_custom_emoji_id,omitempty"`
	// Bio is the bio of the other party in a private chat.
	// Returned only in GetChat
	Bio string `json:"bio,omitempty"`
	// HasPrivateForwards is true, if privacy settings of the other party in the private chat
	// allows to use tg://user?id=<user_id> links only in chats with the user.
	// Returned only in GetChat.
	HasPrivateForwards bool `json:"has_private_forwards,omitempty"`
	// HasRestrictedVoiceAndVideoMessages is true, if the privacy settings of the other party
	// restrict sending voice and video note messages in the private chat.
	// Returned only in GetChat.
	HasRestrictedVoiceAndVideoMessages bool `json:"has_restricted_voice_and_video_messages,omitempty"`
	// JoinToSendMessages is true, if users need to join the supergroup before they can send messages.
	// Returned only in getChat.
	JoinToSendMessages bool `json:"join_to_send_messages,omitempty"`
	// JoinByRequest is true, if all users directly joining the supergroup need to be approved by
	//  supergroup administrators.
	// Returned only in getChat.
	JoinByRequest bool `json:"join_by_request,omitempty"`
	// Description is the description, for groups, supergroups and channel chats.
	// Returned only in getChat.
	Description string `json:"description,omitempty"`
	// InviteLink is the primary invite link, for groups, supergroups and channel chats.
	// Returned only in getChat.
	InviteLink string `json:"invite_link,omitempty"`
	// PinnedMessage is  the most recent pinned message (by sending date).
	// Returned only in getChat.
	PinnedMessage *Message `json:"pinned_message,omitempty"`
	// Permissions is the default chat member permissions, for groups and supergroups.
	// Returned only in getChat.
	Permissions *ChatPermissions `json:"permissions,omitempty"`
	// SlowModeDelay is for supergroups, the minimum allowed delay between consecutive
	// messages sent by each unpriviledged user; in seconds.
	// Returned only in GetChat.
	SlowModeDelay int64 `json:"slow_mode_delay,omitempty"`
	// MessageAutoDeleteTime is the time after which all messages sent
	// to the chat will be automatically deleted; in seconds.
	// Returned only in GetChat.
	MessageAutoDeleteTime int64 `json:"message_auto_delete_time,omitempty"`
	// HasProtectedContent is true, if messages from the chat can't be forwarded to other chats.
	// Returned only in GetChat.
	HasProtectedContent bool `json:"has_protected_content,omitempty"`
	// StickerSetName is, for supergroups, name of group sticker set.
	// Returned only in GetChat.
	StickerSetName string `json:"sticker_set_name,omitempty"`
	// CanSetStickerSet is true, if the bot can change the group sticker set.
	// Returned only in GetChat.
	CanSetStickerSet bool `json:"can_set_sticker_set"`
	// LinkedChatID is the unique identifier for the linked chat.
	// Returned only in GetChat.
	LinkedChatID string `json:"linked_chat_id"`
	// Location is, for supergroups, the location to which the supergroup is connected.
	// Returned only in GetChat.
	Location *ChatLocation `json:"location,omitempty"`
}

// ChatPhoto represents a chat photo.
type ChatPhoto struct {
	// SmallFileID is the file identifier of small (160*160) chat photo.
	SmallFileID string `json:"small_file_id,omitempty"`
	// SmallFileUniqueID is the unique file identifier of small (160x160) chat photo,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	SmallFileUniqueID string `json:"small_file_unique_id,omitempty"`
	// BigFileID is the file identifier of big (640x640) chat photo.
	BigFileID string `json:"big_file_id,omitempty"`
	// BigFileUniqueID is the unique file identifier of big (640x640) chat photo,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	BigFileUniqueID string `json:"big_file_unique_id,omitempty"`
}

// ChatPermissions describes actions that a non-administrator user is allowed to take in a chat.
type ChatPermissions struct {
	// CanSendMessages is true, if the user is allowed to send text messages, contacts, locations and venues.
	CanSendMessages bool `json:"can_send_messages,omitempty"`
	// CanSendMediaMessages is true, if  the user is allowed to send audios, documents, photos, videos,
	// video notes and voice notes, implies can_send_messages.
	CanSendMediaMessages bool `json:"can_send_media_messages,omitempty"`
	// CanSendPolls is true, if the user is allowed to send polls, implies can_send_messages.
	CanSendPolls bool `json:"can_send_polls,omitempty"`
	// CanSendOtherMessages is true, if the user is allowed to send animations, games, stickers
	// and use inline bots, implies can_send_media_messages.
	CanSendOtherMessages bool `json:"can_send_other_messages,omitempty"`
	// CanAddWebPagesPreviews is true, if the user is allowed to add web page previews to their messages,
	// implies can_send_media_messages
	CanAddWebPagesPreviews bool `json:"can_add_web_page_previews,omitempty"`
	// CanChangeInfo is true, if the user is allowed to change the chat title, photo and other settings.
	// Ignored in public supergroups
	CanChangeInfo bool `json:"can_change_info,omitempty"`
	// CanInviteUsers is true, if the user is allowed to invite new users to the chat.
	CanInviteUsers bool `json:"can_invite_users,omitempty"`
	// CanPinMessages is true, if the user is allowed to pin messages.
	// Ignored in public supergroups.
	CanPinMessages bool `json:"can_pin_messages,omitempty"`
	// CanManageTopics is true, if the user is allowed to create forum topics.
	// If omitted defaults to the value of can_pin_messages
	CanManageTopics bool `json:"can_manage_topics,omitempty"`
}

// ChatLocation represents a location to which a chat is connected.
type ChatLocation struct {
	// Location is the the location to which the supergroup is connected.
	//
	// It is a required field
	Location Location `json:"location,omitempty"`
	// Address is location address; 1-64 characters, as defined by the chat owner
	//
	// It is a required field
	Address string `json:"address,omitempty"`
}
