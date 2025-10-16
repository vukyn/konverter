package usecase

import (
	"errors"
	"time"

	timestampmodels "konverter/internal/timestamp/models"

	humanize "github.com/dustin/go-humanize"
)

// ConvertHumanize processes a timestamp conversion request and returns formatted results
func ConvertHumanize(req timestampmodels.ConvertHumanizeRequest) (timestampmodels.ConvertHumanizeResponse, error) {
	if err := req.Validate(); err != nil {
		return timestampmodels.ConvertHumanizeResponse{}, err
	}

	// Detect the unit of the input timestamp
	detectedUnit := detectUnit(req.Timestamp)

	// Normalize to seconds for time calculations
	seconds := normalizeToSeconds(req.Timestamp, detectedUnit)

	// Create time object from seconds
	t := time.Unix(seconds, 0)

	// Build response
	responseHumanize := timestampmodels.ConvertHumanizeResponse{
		InputTimestamp: req.Timestamp,
		DetectedUnit:   detectedUnit,
		Seconds:        seconds,
		Milliseconds:   normalizeToMilliseconds(req.Timestamp, detectedUnit),
		Microseconds:   normalizeToMicroseconds(req.Timestamp, detectedUnit),
		Nanoseconds:    normalizeToNanoseconds(req.Timestamp, detectedUnit),
		GMT:            t.UTC().Format(time.RFC3339),
		Relative:       humanize.Time(t),
	}

	// Handle timezone-specific time if provided
	if req.Timezone != "" {
		timezoneTime, err := formatTimezoneTime(t, req.Timezone)
		if err != nil {
			return timestampmodels.ConvertHumanizeResponse{}, errors.New("invalid timezone: " + err.Error())
		}
		responseHumanize.TimezoneTime = timezoneTime
	}

	return responseHumanize, nil
}

// ConvertDateToUnix processes a date string conversion request and returns Unix timestamp results
func ConvertDateToUnix(req timestampmodels.DateToUnixRequest) (timestampmodels.DateToUnixResponse, error) {
	if err := req.Validate(); err != nil {
		return timestampmodels.DateToUnixResponse{}, err
	}

	// Parse the date string using multiple format attempts
	parsedTime, detectedFormat, err := tryParseDateFormats(req.DateString)
	if err != nil {
		return timestampmodels.DateToUnixResponse{}, err
	}

	// Convert to Unix timestamp in seconds
	unixSeconds := parsedTime.Unix()

	// Build response
	response := timestampmodels.DateToUnixResponse{
		InputDateString: req.DateString,
		DetectedFormat:  detectedFormat,
		Seconds:         unixSeconds,
		Milliseconds:    unixSeconds * 1000,
		Microseconds:    unixSeconds * 1000000,
		Nanoseconds:     unixSeconds * 1000000000,
		GMT:             parsedTime.UTC().Format(time.RFC3339),
	}

	// Handle timezone-specific time if provided
	if req.Timezone != "" {
		timezoneTime, err := formatTimezoneTime(parsedTime, req.Timezone)
		if err != nil {
			return timestampmodels.DateToUnixResponse{}, errors.New("invalid timezone: " + err.Error())
		}
		response.TimezoneTime = timezoneTime
	}

	return response, nil
}
