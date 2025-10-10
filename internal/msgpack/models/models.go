package models

import "errors"

type EncodeRequest struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

func (r *EncodeRequest) Validate() error {
	if r.Type != "base64" && r.Type != "bytes" {
		return errors.New("type must be either 'base64' or 'bytes'")
	}
	if r.Data == "" {
		return errors.New("data is required")
	}
	return nil
}

type DecodeRequest struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

func (r *DecodeRequest) Validate() error {
	if r.Type != "base64" && r.Type != "bytes" {
		return errors.New("type must be either 'base64' or 'bytes'")
	}
	if r.Data == "" {
		return errors.New("data is required")
	}
	return nil
}
