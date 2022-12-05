package gotbot

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// request setups
var (
	// SetApplicationJSON is a function that sets the content-type of a request to application/json
	SetApplicationJSON = func(req *http.Request) error {
		req.Header.Set("Content-Type", "application/json")
		return nil
	}
	// SetMultipartFormData is a function that sets the content-type of a request to multipart/form-data
	SetMultipartFormData = func(req *http.Request) error {
		req.Header.Set("Content-Type", "multipart/form-data")
		return nil
	}
)

// body setters
var (
	// GetJSONBody marshals a given object to a json serialized string
	GetJSONBody = func(value any) (io.Reader, error) {
		body, err := json.Marshal(value)
		return bytes.NewBuffer(body), err
	}
	// GetMultipartBody creates a form data with the given fields and files
	GetMultipartBody = func(fields []FormField, files []FormFile) (io.Reader, error) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		for i := 0; i < len(fields); i++ {
			if err := writer.WriteField(fields[i].Name, fields[i].Value); err != nil {
				return nil, err
			}
		}

		for i := 0; i < len(files); i++ {
			if err := func() error {
				file, err := os.Open(files[i].Path)
				if err != nil {
					return err
				}
				defer func(file *os.File) {
					_ = file.Close()
				}(file)

				fileField, err := writer.CreateFormFile(files[i].Name, filepath.Base(files[i].Path))
				if err != nil {
					return err
				}
				_, err = io.Copy(fileField, file)
				if err != nil {
					return err
				}
				return writer.Close()
			}(); err != nil {
				return nil, err
			}
		}

		return body, nil
	}
)

// FormFile describes a file to be added to a form data
type FormFile struct {
	Name string
	Path string
}

// FormField describes a non-file field value to be added to a form data
type FormField struct {
	Name  string
	Value string
}

type MessageType string

const (
	// MessageText is for text message.
	MessageText MessageType = "sendMessage"
	// MessagePhoto is for single photo message.
	MessagePhoto = "sendPhoto"
	// MessageAudio is for single audio message.
	MessageAudio = "sendAudio"
	// MessageDocument is for single document message.
	MessageDocument = "sendDocument"
	// MessageVideo is for single video message.
	MessageVideo = "sendVideo"
	// MessageAnimation is for single animation message.
	MessageAnimation = "sendAnimation"
	// MessageVoice is for single voice message.
	MessageVoice = "sendVoice"
	// MessageVideoNote is for rounded, 1 minute MPEG4 videos.
	MessageVideoNote = "sendVideoNote"
	// MessageMediaGroup is for a group of media messages including photos, audios, videos, documents.
	MessageMediaGroup = "sendMediaGroup"
	// MessageLocation is for  location message.
	MessageLocation = "sendLocation"
	// MessageVenue is for venue message.
	MessageVenue = "sendVenue"
	// MessageContact is for phone contact message.
	MessageContact = "sendContact"
	// MessagePoll is for poll message.
	MessagePoll = "sendPoll"
	// MessageDice is for dice message.
	MessageDice = "sendDice"
	// MessageChatAction is for chat action message.
	MessageChatAction = "sendChatAction"
)
