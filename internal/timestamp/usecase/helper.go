package usecase

import "time"

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
