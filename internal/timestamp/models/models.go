package models

import (
	"errors"
)

type ConvertHumanizeRequest struct {
	// Timestamp is the unix timestamp input
	Timestamp int64 `json:"timestamp"`
	// Timezone is optional timezone for timezone-specific output (e.g., "America/New_York", "Asia/Tokyo")
	Timezone string `json:"timezone,omitempty"`
}

func (r *ConvertHumanizeRequest) Validate() error {
	if r.Timestamp == 0 {
		return errors.New("timestamp is required")
	}

	// Check if timestamp is within reasonable bounds
	// Earliest reasonable timestamp: 1970-01-01 (0)
	// Latest reasonable timestamp: year 3000 (32503680000 seconds)
	if r.Timestamp < 0 {
		return errors.New("timestamp cannot be negative")
	}

	// If it's in microseconds, check it's not too far in the future
	if r.Timestamp >= 10000000000000000 { // 10^16 microseconds = year ~316,000
		return errors.New("timestamp is too far in the future")
	}

	return nil
}

type ConvertHumanizeResponse struct {
	// InputTimestamp is the original input timestamp
	InputTimestamp int64 `json:"input_timestamp"`
	// DetectedUnit indicates what unit was detected: "seconds", "milliseconds", or "microseconds"
	DetectedUnit string `json:"detected_unit"`
	// Seconds is the timestamp normalized to seconds
	Seconds int64 `json:"seconds"`
	// Milliseconds is the timestamp normalized to milliseconds
	Milliseconds int64 `json:"milliseconds"`
	// Microseconds is the timestamp normalized to microseconds
	Microseconds int64 `json:"microseconds"`
	// Nanoseconds is the timestamp normalized to nanoseconds
	Nanoseconds int64 `json:"nanoseconds"`
	// GMT is the time in RFC3339 format in GMT/UTC
	GMT string `json:"gmt"`
	// TimezoneTime is the time in the specified timezone (if provided)
	TimezoneTime string `json:"timezone_time,omitempty"`
	// Relative is human-readable relative time (e.g., "2 hours ago", "in 5 minutes")
	Relative string `json:"relative"`
}

type DateToUnixRequest struct {
	// DateString is the input date string in various formats
	DateString string `json:"date_string"`
	// Timezone is optional timezone for timezone-specific parsing (e.g., "America/New_York", "Asia/Tokyo")
	Timezone string `json:"timezone,omitempty"`
}

func (r *DateToUnixRequest) Validate() error {
	if r.DateString == "" {
		return errors.New("date_string is required")
	}
	return nil
}

type DateToUnixResponse struct {
	// InputDateString is the original input date string
	InputDateString string `json:"input_date_string"`
	// DetectedFormat indicates what format was detected and used for parsing
	DetectedFormat string `json:"detected_format"`
	// Seconds is the timestamp in seconds
	Seconds int64 `json:"seconds"`
	// Milliseconds is the timestamp in milliseconds
	Milliseconds int64 `json:"milliseconds"`
	// Microseconds is the timestamp in microseconds
	Microseconds int64 `json:"microseconds"`
	// Nanoseconds is the timestamp in nanoseconds
	Nanoseconds int64 `json:"nanoseconds"`
	// GMT is the time in RFC3339 format in GMT/UTC
	GMT string `json:"gmt"`
	// TimezoneTime is the time in the specified timezone (if provided)
	TimezoneTime string `json:"timezone_time,omitempty"`
}
