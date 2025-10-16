package usecase

import (
	"errors"
	"time"
)

// Determines the unit of the timestamp based on its magnitude
func detectUnit(timestamp int64) string {
	if timestamp < 10000000000 { // 10^10
		return "seconds"
	} else if timestamp < 10000000000000 { // 10^13
		return "milliseconds"
	}
	return "microseconds"
}

// Converts any timestamp unit to seconds
func normalizeToSeconds(timestamp int64, unit string) int64 {
	switch unit {
	case "seconds":
		return timestamp
	case "milliseconds":
		return timestamp / 1000
	case "microseconds":
		return timestamp / 1000000
	case "nanoseconds":
		return timestamp / 1000000000
	default:
		return timestamp
	}
}

// Converts any timestamp unit to milliseconds
func normalizeToMilliseconds(timestamp int64, unit string) int64 {
	switch unit {
	case "seconds":
		return timestamp * 1000
	case "milliseconds":
		return timestamp
	case "microseconds":
		return timestamp / 1000
	case "nanoseconds":
		return timestamp / 1000000000
	default:
		return timestamp
	}
}

// Converts any timestamp unit to microseconds
func normalizeToMicroseconds(timestamp int64, unit string) int64 {
	switch unit {
	case "seconds":
		return timestamp * 1000000
	case "milliseconds":
		return timestamp * 1000
	case "microseconds":
		return timestamp
	case "nanoseconds":
		return timestamp / 1000000
	default:
		return timestamp
	}
}

// Converts any timestamp unit to nanoseconds
func normalizeToNanoseconds(timestamp int64, unit string) int64 {
	switch unit {
	case "seconds":
		return timestamp * 1000000000
	case "milliseconds":
		return timestamp * 1000000
	case "microseconds":
		return timestamp * 1000
	case "nanoseconds":
		return timestamp / 1000
	default:
		return timestamp
	}
}

// formatTimezoneTime formats the time in the specified timezone
func formatTimezoneTime(t time.Time, timezone string) (string, error) {
	// Load the timezone location
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return "", err
	}

	// Convert to the specified timezone and format
	return t.In(loc).Format(time.RFC3339), nil
}

// tryParseDateFormats attempts to parse a date string using multiple format layouts
// Returns the parsed time and the format name that matched, or an error if no format matches
func tryParseDateFormats(dateString string) (time.Time, string, error) {
	// Strip "GMT" prefix if present to enable local timezone parsing
	cleanedString := dateString
	if len(dateString) > 3 && dateString[:3] == "GMT" {
		cleanedString = dateString[3:]
	}

	// Define common date formats in order of preference
	formats := []struct {
		layout string
		name   string
	}{
		// RFC formats
		{time.RFC3339, "RFC3339"},
		{time.RFC3339Nano, "RFC3339Nano"},
		{time.RFC1123, "RFC1123"},
		{time.RFC1123Z, "RFC1123Z"},
		{time.RFC822, "RFC822"},
		{time.RFC822Z, "RFC822Z"},
		{time.RFC850, "RFC850"},

		// ISO formats
		{"2006-01-02T15:04:05Z", "ISO8601"},
		{"2006-01-02T15:04:05.000Z", "ISO8601_milliseconds"},
		{"2006-01-02T15:04:05.000000Z", "ISO8601_microseconds"},
		{"2006-01-02T15:04:05.000000000Z", "ISO8601_nanoseconds"},

		// Date only formats
		{"2006-01-02", "Y-M-D"},
		{"2006/01/02", "Y/M/D"},
		{"02-01-2006", "D-M-Y"},
		{"02/01/2006", "D/M/Y"},
		{"01-02-2006", "M-D-Y"},
		{"01/02/2006", "M/D/Y"},
		{"2006-1-2", "Y-M-D_single_digit"},
		{"2006/1/2", "Y/M/D_single_digit"},
		{"2-1-2006", "D-M-Y_single_digit"},
		{"2/1/2006", "D/M/Y_single_digit"},
		{"1-2-2006", "M-D-Y_single_digit"},
		{"1/2/2006", "M/D/Y_single_digit"},

		// Date with time formats
		{"2006-01-02 15:04:05", "Y-M-D H:M:S"},
		{"2006/01/02 15:04:05", "Y/M/D H:M:S"},
		{"02-01-2006 15:04:05", "D-M-Y H:M:S"},
		{"02/01/2006 15:04:05", "D/M/Y H:M:S"},
		{"01-02-2006 15:04:05", "M-D-Y H:M:S"},
		{"01/02/2006 15:04:05", "M/D/Y H:M:S"},

		// Date with time and timezone
		{"2006-01-02 15:04:05 MST", "Y-M-D H:M:S MST"},
		{"2006/01/02 15:04:05 MST", "Y/M/D H:M:S MST"},
		{"02-01-2006 15:04:05 MST", "D-M-Y H:M:S MST"},
		{"02/01/2006 15:04:05 MST", "D/M/Y H:M:S MST"},
		{"01-02-2006 15:04:05 MST", "M-D-Y H:M:S MST"},
		{"01/02/2006 15:04:05 MST", "M/D/Y H:M:S MST"},
	}

	// Try each format
	for _, format := range formats {
		if t, err := time.Parse(format.layout, cleanedString); err == nil {
			return t, format.name, nil
		}
	}

	return time.Time{}, "", errors.New("unable to parse date string with any supported format")
}
