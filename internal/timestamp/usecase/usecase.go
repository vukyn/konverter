package usecase

import (
	"errors"
	"time"

	timestampmodels "konverter/internal/timestamp/models"

	humanize "github.com/dustin/go-humanize"
)

// Convert processes a timestamp conversion request and returns formatted results
func ConvertHumanize(req timestampmodels.ConvertRequest) (timestampmodels.ConvertHumanizeResponse, error) {
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
