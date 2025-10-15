package usecase

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"konverter/internal/msgpack/models"

	jsoniter "github.com/json-iterator/go"
	"github.com/vmihailenco/msgpack/v5"
)

// Encodes JSON data to MessagePack
func Encode(req models.EncodeRequest) (string, error) {
	if err := req.Validate(); err != nil {
		return "", err
	}

	// Parse input data
	var data any
	err := jsoniter.Unmarshal([]byte(req.Data), &data)
	if err != nil {
		return "", errors.New("invalid JSON data: " + err.Error())
	}

	// Encode to msgpack and return as base64 encoded string
	msgpackData, err := msgpack.Marshal(data)
	if err != nil {
		return "", errors.New("failed to encode msgpack: " + err.Error())
	}

	// Return as base64 encoded string or raw bytes
	switch req.Type {
	case "base64":
		encoded := base64.StdEncoding.EncodeToString(msgpackData)
		return encoded, nil
	case "bytes":
		return fmt.Sprintf("%v", msgpackData), nil
	default:
		return "", errors.New("invalid request type: " + req.Type)
	}
}

// Decodes MessagePack data to JSON
func Decode(req models.DecodeRequest) (any, error) {
	if err := req.Validate(); err != nil {
		return "", err
	}

	var data []byte
	var err error

	// Decode from base64 encoded string or raw bytes
	switch req.Type {
	case "base64":
		data, err = base64.StdEncoding.DecodeString(req.Data)
		if err != nil {
			return "", err
		}
	case "bytes":
		// Try to parse as byte array format first, fallback to raw string
		if strings.HasPrefix(strings.TrimSpace(req.Data), "[") {
			data, err = parseByteArray(req.Data)
			if err != nil {
				return "", err
			}
		} else {
			data = []byte(req.Data)
		}
	default:
		return "", errors.New("invalid request type: " + req.Type)
	}

	// Decode from msgpack and return as any type
	var decoded any
	err = msgpack.Unmarshal(data, &decoded)
	if err != nil {
		return "", errors.New("failed to decode msgpack: " + err.Error())
	}

	return decoded, nil
}

// Parses a string like "[123 111 100]" into a byte slice
func parseByteArray(s string) ([]byte, error) {
	// Remove brackets and trim spaces
	s = strings.TrimSpace(s)
	if !strings.HasPrefix(s, "[") || !strings.HasSuffix(s, "]") {
		return nil, errors.New("invalid byte array format: must be [123 111 100]")
	}

	s = strings.Trim(s, "[]")
	s = strings.ReplaceAll(s, ",", "")
	if s == "" {
		return []byte{}, nil
	}

	// Split by spaces
	parts := strings.Fields(s)
	result := make([]byte, len(parts))

	for i, part := range parts {
		val, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("invalid byte value '%s': %v", part, err)
		}
		if val < 0 || val > 255 {
			return nil, fmt.Errorf("byte value %d out of range [0, 255]", val)
		}
		result[i] = byte(val)
	}

	return result, nil
}
