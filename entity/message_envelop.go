package entity

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

// MessageEnvelop holds the object that is used to send a new message
type MessageEnvelop struct {
	// ChatID is the unique identifier for the target chat
	// or username of the target channel
	//
	// It is a required field.
	ChatID string `json:"chat_id,omitempty"`
	// MessageThreadID is the unique identifier for the target message thread (topic) of the forum;
	// for forum supergroups only
	MessageThreadID int64 `json:"message_thread_id,omitempty"`
	// Text is text of the message to be sent, 1-4096 characters after entities parsing.
	//
	// It is a required field
	Text string `json:"text,omitempty"`
	// ParseMode is the mode for parsing entities in the message text.
	ParseMode string `json:"parse_mode,omitempty"`
	// Entities are A JSON-serialized list of special entities that appear in message text,
	// which can be specified instead of parse_mode.
	Entities []MessageEntity `json:"entities,omitempty"`
	// DisableWebPagePreview disables link previews for links in this message.
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`
	// DisableNotification sends the message silently.
	// Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// ProtectContent protects the contents of the sent message from forwarding and saving.
	ProtectContent bool `json:"protect_content,omitempty"`
	// ReplyToMessageID is, if the message is a reply, ID of the original message.
	ReplyToMessageID int64 `json:"reply_to_message_id,omitempty"`
	// AllowSendingWithoutReply can be true if the message should be sent
	// even if the specified replied-to message is not found.
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	// ReplyMarkup contains additional interface options.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
	// Photo is the photo to send.
	// Pass a file_id as String to send a photo that exists on the Telegram servers (recommended),
	// pass an HTTP URL as a String for Telegram to get a photo from the Internet,
	// or upload a new photo using multipart/form-data.
	//
	// It is a required field for sending a photo.
	Photo *FileEnvelop `json:"photo,omitempty"`
	// Caption is the media caption
	Caption string `json:"caption,omitempty"`
	// CaptionEntities is list of special entities that appear in the caption,
	// which can be specified instead of parse_mode.
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	// Audio is the audio file to send.
	// Pass a file_id as String to send an audio file that exists on the Telegram servers (recommended),
	// pass an HTTP URL as a String for Telegram to get an audio file from the Internet,
	// or upload a new one using multipart/form-data.
	//
	// It is a required field for sending an audio.
	Audio *FileEnvelop `json:"audio,omitempty"`
	// Duration of the media in seconds.
	Duration int64 `json:"duration,omitempty"`
	// Performer of the media
	Performer string `json:"performer,omitempty"`
	// Title is the track name
	Title string `json:"title,omitempty"`
	// Thumb is the thumbnail of the file sent.
	// It should be in JPEG format and less than 200 kB in size.
	// It's width and height should not exceed 320.
	// Ignored if the file is not uploaded using multipart/form-data.
	Thumb *FileEnvelop `json:"thumb,omitempty"`
	// Document is the document to send.
	// Pass a file_id as String to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL as a String for Telegram to get a file from the Internet,
	// or upload a new one using multipart/form-data.
	//
	// It is a required field for sending a document.
	Document *FileEnvelop `json:"document,omitempty"`
	// DisableContentTypeDetection disables automatic server-side content type detection
	// for files uploaded using multipart/form-data.
	DisableContentTypeDetection bool `json:"disable_content_type_detection,omitempty"`
	// Video is the Video to send.
	// Pass a file_id as String to send a video that exists on the Telegram servers (recommended),
	// pass an HTTP URL as a String for Telegram to get a video from the Internet,
	// or upload a new video using multipart/form-data.
	//
	// It is a required field for sending video.
	Video *FileEnvelop `json:"video,omitempty"`
	// Width of the media
	Width int64 `json:"width,omitempty"`
	// Height of the media
	Height int64 `json:"height,omitempty"`
	// SupportsStreaming can be true if the uploaded media is suitable for streaming.
	SupportsStreaming bool `json:"supports_streaming,omitempty"`
	// Animation is the animation to send.
	// Pass a file_id as String to send an animation that exists on the Telegram servers (recommended),
	// pass an HTTP URL as a String for Telegram to get an animation from the Internet,
	// or upload a new animation using multipart/form-data.
	//
	// It is a required field for sending an animation.
	Animation *FileEnvelop `json:"animation,omitempty"`
	// Voice is the audio file to send.
	// Pass a file_id as String to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL as a String for Telegram to get a file from the Internet,
	// or upload a new one using multipart/form-data.
	//
	// It is a required field for sending a voice.
	Voice *FileEnvelop `json:"voice,omitempty"`
	// VideoNote is the video note to send.
	// Pass a file_id as String to send a video note that exists on the Telegram servers (recommended)
	// or upload a new video using multipart/form-data.
	// Sending video notes by a URL is currently unsupported.
	//
	// It is a required field for sending a video note.
	VideoNote *FileEnvelop `json:"video_note,omitempty"`
	// Length is the media width and height, i.e. diameter of the video message.
	Length int64 `json:"length,omitempty"`
	// Media is an array of media messages to be sent, must include 2-10 items.
	Media []InputMedia `json:"media,omitempty"`

	// All Parameters for sendLocation Method
	Location

	// All Parameters for sendContact Method
	Contact

	// All Parameters for sendPoll Method

	// Question is poll question, 1-300 characters
	//
	// It is a required field for sending a poll.
	Question string `json:"question"`
	// Options is  list of poll options
	//
	// It is a required field for sending a poll.
	Options []string `json:"options,omitempty"`
	// IsAnonymous is true, if the poll is anonymous
	//
	// It is a required field for sending a poll.
	IsAnonymous bool `json:"is_anonymous"`
	// Type	is poll type, currently can be “regular” or “quiz”
	//
	// It is a required field for sending a poll.
	Type string `json:"type"`
	// AllowsMultipleAnswers is true, if the poll allows multiple answers
	//
	// It is a required field for sending a poll.
	AllowsMultipleAnswers bool `json:"allows_multiple_answers"`
	// CorrectOptionID is  0-based identifier of the correct answer option.
	//
	// Required for polls in the quiz mode.
	CorrectOptionID int64 `json:"correct_option_id"`
	// Explanation is text that is shown when a user chooses an incorrect answer
	// or taps on the lamp icon in a quiz-style poll, 0-200 characters
	Explanation string `json:"explanation"`
	// ExplanationEntities is special entities like usernames, URLs, bot commands, etc.
	ExplanationEntities []*MessageEntity `json:"explanation_entities"`
	// OpenPeriod is amount of time in seconds the poll will be active after creation
	// can't be used together with close_date
	OpenPeriod int64 `json:"open_period"`
	// CloseDate is point in time (Unix timestamp) when the poll will be automatically closed
	// can't be used together with open_period
	CloseDate int64 `json:"close_date"`
	// IsClosed can be true, if the poll needs to be immediately closed
	IsClosed bool `json:"is_closed"`

	// All Parameters for sendDice Method

	// Emoji on which the dice throw animation is based.
	Emoji string `json:"emoji,omitempty"`

	// Action is type of action to broadcast. Choose one, depending on what the user is about to receive.
	// typing for text messages, upload_photo for photos, record_video or upload_video for videos,
	// record_voice or upload_voice for voice notes, upload_document for general files, choose_sticker for stickers,
	// find_location for location data, record_video_note or upload_video_note for video notes.
	//
	// It is a required field for sending a chat action.
	Action ChatAction `json:"action"`

	// All Parameters for sendVenue method

	// Address of the venue
	Address string `json:"address,omitempty"`
	// FoursquareID of the venue
	FoursquareID string `json:"foursquare_id,omitempty"`
	// FoursquareType of the venue, if known.
	// (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	FoursquareType string `json:"foursquare_type,omitempty"`
	// GooglePlaceID of the venue
	GooglePlaceID string `json:"google_place_id,omitempty"`
	// GooglePlaceType of the venue.
	// See supported types at https://developers.google.com/places/web-service/supported_types
	GooglePlaceType string `json:"google_place_type,omitempty"`
}

// FileEnvelop represents a file to be uploaded to the telegram server.
type FileEnvelop struct {
	// Path is a file described as,
	//
	// a file_id on the telegram server or
	//
	// an http url for the file from the internet or
	//
	// a file on this device (must be prefixed with `file://`)
	Path string
	// Name is the name to append this file with on the request body.
	// It is only used if the *FileEnvelop.SetValue is called with an empty name.
	Name string
}

func (f *FileEnvelop) SetValue(writer *multipart.Writer, name string) error {
	if len(name) == 0 {
		name = f.Name
	}

	if strings.HasPrefix(f.Path, "file://") {
		path := strings.TrimPrefix(f.Path, "file://")

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer func(file *os.File) {
			_ = file.Close()
		}(file)

		fileField, err := writer.CreateFormFile(name, filepath.Base(path))
		if err != nil {
			return err
		}

		_, err = io.Copy(fileField, file)

		return err
	}

	return writer.WriteField(name, f.Path)
}
