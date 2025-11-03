package models

import (
	"errors"
	"unicode/utf8"
)

const (
	MaxTextSize     = 10 * 1024 * 1024 // 10MB in bytes
	MinSecretLength = 8
)

type EncryptRequest struct {
	Text    string `json:"text"`    // Plain text to encrypt (max 10MB)
	Secret  string `json:"secret"`  // Encryption secret (min 8 characters)
	Salt    string `json:"salt"`    // Salt for HKDF key derivation (optional)
	CtxInfo string `json:"ctx_info"` // Context info for HKDF key derivation (optional)
}

func (r *EncryptRequest) Validate() error {
	if r.Text == "" {
		return errors.New("text is required")
	}

	if r.Secret == "" {
		return errors.New("secret is required")
	}

	if len(r.Secret) < MinSecretLength {
		return errors.New("secret must be at least 8 characters")
	}

	// Check text size in bytes (UTF-8 encoding)
	textSize := utf8.RuneCountInString(r.Text)
	if textSize > MaxTextSize {
		return errors.New("text size exceeds maximum limit of 10MB")
	}

	return nil
}

type DecryptRequest struct {
	Text    string `json:"text"`    // Base64 encoded encrypted text
	Secret  string `json:"secret"`  // Decryption secret (min 8 characters)
	Salt    string `json:"salt"`    // Salt used during encryption (optional)
	CtxInfo string `json:"ctx_info"` // Context info used during encryption (optional)
}

func (r *DecryptRequest) Validate() error {
	if r.Text == "" {
		return errors.New("text is required")
	}

	if r.Secret == "" {
		return errors.New("secret is required")
	}

	if len(r.Secret) < MinSecretLength {
		return errors.New("secret must be at least 8 characters")
	}

	return nil
}

type EncryptResponse struct {
	EncryptedText string `json:"encrypted_text"`
}

type DecryptResponse struct {
	DecryptedText string `json:"decrypted_text"`
}

