package dtypes

import (
	"errors"
	"strings"
)

type PulseStateStatus string

const (
	PulseStateStatusSuccess PulseStateStatus = "success"
	PulseStateStatusFailed  PulseStateStatus = "failed"
)

var (
	// mapPulseStateStatus implementation of a set for PulseStateStatus
	mapPulseStateStatus = map[PulseStateStatus]struct{}{
		PulseStateStatusSuccess: {},
		PulseStateStatusFailed:  {},
	}
)

// ParseAppealStatus converting string to status
func ParsePulseStateStatus(str string) (PulseStateStatus, error) {
	status := PulseStateStatus(strings.ToLower(str))

	if _, exists := mapPulseStateStatus[status]; !exists {
		return "", errors.New("invalid value for PulseStateStatus")
	}

	return status, nil
}
