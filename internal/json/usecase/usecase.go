package usecase

import (
	"bytes"
	stdjson "encoding/json"
	"errors"
	"strconv"

	jsonmodels "konverter/internal/json/models"

	jsoniter "github.com/json-iterator/go"
)

// Validates that input is JSON and returns an escaped JSON string
func Escape(req jsonmodels.EscapeRequest) (string, error) {
	if err := req.Validate(); err != nil {
		return "", err
	}

	// Escape as valid Go string literal which matches JSON escaping
	escaped := strconv.Quote(req.Data)
	// Drop the surrounding quotes to return the inner escaped content
	if len(escaped) >= 2 {
		escaped = escaped[1 : len(escaped)-1]
	}
	return escaped, nil
}

// Reverses escaping and validates the result is valid JSON
func Unescape(req jsonmodels.UnescapeRequest) (string, error) {
	if err := req.Validate(); err != nil {
		return "", err
	}

	// Add quotes so Unquote can interpret escapes
	unquoted, err := strconv.Unquote("\"" + req.Data + "\"")
	if err != nil {
		return "", errors.New("invalid escaped JSON string: " + err.Error())
	}

	// Validate that the unescaped content is valid JSON
	var v any
	if err := jsoniter.Unmarshal([]byte(unquoted), &v); err != nil {
		return "", errors.New("invalid JSON data after unescape: " + err.Error())
	}

	return unquoted, nil
}

// Formats/pretty-prints JSON string with proper indentation
func Format(req jsonmodels.FormatRequest) (string, error) {
	if err := req.Validate(); err != nil {
		return "", err
	}

	// Validate JSON without altering key order
	if !jsoniter.Valid([]byte(req.Data)) {
		return "", errors.New("invalid JSON data")
	}

	// Pretty-print while preserving the original key order
	var buf bytes.Buffer
	if err := stdjson.Indent(&buf, []byte(req.Data), "", "  "); err != nil {
		return "", errors.New("failed to format JSON: " + err.Error())
	}

	return buf.String(), nil
}

// Minifies JSON string by removing unnecessary whitespace while preserving key order
func Minify(req jsonmodels.MinifyRequest) (string, error) {
	if err := req.Validate(); err != nil {
		return "", err
	}

	// Validate JSON without altering key order
	if !jsoniter.Valid([]byte(req.Data)) {
		return "", errors.New("invalid JSON data")
	}

	// Minify by compacting the JSON (removing whitespace)
	var buf bytes.Buffer
	if err := stdjson.Compact(&buf, []byte(req.Data)); err != nil {
		return "", errors.New("failed to minify JSON: " + err.Error())
	}

	return buf.String(), nil
}
