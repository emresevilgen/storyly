package go_log

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
	"time"
)

/*
 * Inspired by the https://github.com/sirupsen/logrus/blob/master/json_formatter.go
 * Instead of using json.Marshal(), PrettJSONFormatter uses json.MarshalIndent().
 */

type PrettyJSONFormatter struct {
	// TimestampFormat sets the format used for marshaling timestamps.
	TimestampFormat string
}

func (f *PrettyJSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(logrus.Fields, len(entry.Data)+3)

	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			// Otherwise errors are ignored by `encoding/json`
			// https://github.com/Sirupsen/logrus/issues/137
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}

	prefixFieldClashes(data)

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = time.RFC3339 // default timestamp format
	}

	data["time"] = entry.Time.Format(timestampFormat)
	data["msg"] = entry.Message
	data["level"] = entry.Level.String()

	serialized, err := json.MarshalIndent(data, "", "  ")

	if err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %+v", err)
	}

	return append(serialized, '\n'), nil
}

func prefixFieldClashes(data logrus.Fields) {
	_, ok := data["time"]
	if ok {
		data["fields.time"] = data["time"]
	}

	_, ok = data["msg"]
	if ok {
		data["fields.msg"] = data["msg"]
	}

	_, ok = data["level"]
	if ok {
		data["fields.level"] = data["level"]
	}
}
