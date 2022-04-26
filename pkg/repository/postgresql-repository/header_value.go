package postgresql_repository

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type HeaderValue struct {
	Key string `json:"key,omitempty"`
	Val string `json:"value,omitempty"`
}

type Headers []HeaderValue

func (h Headers) Value() (driver.Value, error) {
	return json.Marshal(h)
}

func (h *Headers) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("Type assertion to []byte failed")
	}

	return json.Unmarshal(b, &h)
}
