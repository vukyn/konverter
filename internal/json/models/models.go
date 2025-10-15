package models

import (
	"errors"

	jsoniter "github.com/json-iterator/go"
)

type EscapeRequest struct {
	// Data is the JSON text to be escaped (e.g., {"a":"b"})
	Data string `json:"data"`
}

func (r *EscapeRequest) Validate() error {
	if r.Data == "" {
		return errors.New("data is required")
	}

	// Ensure input is valid JSON
	var v any
	if err := jsoniter.Unmarshal([]byte(r.Data), &v); err != nil {
		return errors.New("invalid JSON data: " + err.Error())
	}
	return nil
}

type UnescapeRequest struct {
	// Data is the escaped JSON text (e.g., {\"a\":\"b\"})
	Data string `json:"data"`
}

func (r *UnescapeRequest) Validate() error {
	if r.Data == "" {
		return errors.New("data is required")
	}
	return nil
}
