package gotbot

import (
	"encoding/json"
	"fmt"
	"time"
)

// Level is the level of the log generated.
// It can be one of `Debug`, `Info`, `Warn`, or `Error`
type Level string

type Fields map[string]any

const (
	Debug Level = "debug"
	Info  Level = "info"
	Warn  Level = "warn"
	Error Level = "error"
)

// Logger is used to print any messages that might be generated inside this package.
// The default log is json formatted.
//
// Note: default fields such as `time` are not included on `fields`.
// So if you plan on creating a custom logger, be sure to include those fields as well.
type Logger interface {
	Debug(msg string, fields Fields)
	Info(msg string, fields Fields)
	Warn(msg string, fields Fields)
	Error(msg string, fields Fields)
}

// JSONLogger is the default logger.
type JSONLogger struct {
	// TimeFormat contains the layout of the time field on the log.
	// It can be one of the standard time layouts from `time` package.
	TimeFormat string
}

func (d *JSONLogger) Debug(msg string, fields Fields) {
	d.formatLog(Debug, msg, fields)
}

func (d *JSONLogger) Info(msg string, fields Fields) {
	d.formatLog(Info, msg, fields)
}
func (d *JSONLogger) Warn(msg string, fields Fields) {
	d.formatLog(Warn, msg, fields)
}
func (d *JSONLogger) Error(msg string, fields Fields) {
	d.formatLog(Error, msg, fields)
}

func (d *JSONLogger) formatLog(level Level, msg string, fields map[string]any) {
	fields["time"] = time.Now().Format(d.TimeFormat)
	fields["level"] = level
	fields["msg"] = msg

	js, err := json.Marshal(fields)
	if err != nil {
		fmt.Println(fmt.Sprintf(`{"level":"error","msg":"failed to parse the log message to json","error":"%s","time":"%s","log":"%s"}`,
			err.Error(),
			time.Now().Format(d.TimeFormat),
			fmt.Sprintf("msg: %s, fields: %v, level: %s",
				msg, fields, level)))
	}

	fmt.Println(string(js))
}
