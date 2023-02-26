package gotbot

import (
	"encoding/json"
	"fmt"
	"time"
)

// Level is the level of the log generated.
// It can be one of `Debug`, `Info`, `Warn`, or `Error`
type Level string

const (
	Debug Level = "debug"
	Info  Level = "info"
	Warn  Level = "warn"
	Error Level = "error"
)

// LogMessage contains the details of the log to be printed.
type LogMessage struct {
	Level Level     `json:"level"`
	Msg   string    `json:"msg"`
	Time  time.Time `json:"time"`
}

// Logger is used to print any messages that might be generated inside this package.
// The default log is json formatted.
type Logger interface {
	Log(msg LogMessage)
}

type defaultLogger struct{}

func (d *defaultLogger) Log(msg LogMessage) {
	js, _ := json.Marshal(msg) //nolint:errcheck // since LogMessage is safe to marshal

	fmt.Println(string(js))
}

func logDebug(msg string, fields ...any) LogMessage {
	return LogMessage{
		Level: Debug,
		Msg:   fmt.Sprintf(msg, fields...),
		Time:  time.Now(),
	}
}

func logInfo(msg string, fields ...any) LogMessage {
	return LogMessage{
		Level: Info,
		Msg:   fmt.Sprintf(msg, fields...),
		Time:  time.Now(),
	}
}

func logWarn(msg string, fields ...any) LogMessage {
	return LogMessage{
		Level: Warn,
		Msg:   fmt.Sprintf(msg, fields...),
		Time:  time.Now(),
	}
}

func logError(msg string, fields ...any) LogMessage {
	return LogMessage{
		Level: Error,
		Msg:   fmt.Sprintf(msg, fields...),
		Time:  time.Now(),
	}
}
