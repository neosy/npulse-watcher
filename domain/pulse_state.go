package dpulse

import (
	"time"

	dtypes "git.n-hub.ru/neosy/npulse-watcher/domain/types"
)

// PulseState represents the status of a monitored host.
type PulseState struct {
	// IPAddress is the IP address of the monitored host.
	IPAddress string
	// HostName is the name of the monitored host.
	HostName string
	// LastTime records the last time the availability of the host was checked.
	LastTime time.Time
	// Status indicates the result of the last availability check (e.g., success, failed).
	Status dtypes.PulseStateStatus
	// IsNotified indicates whether a notification has been sent regarding the status change.
	IsNotified bool
}
