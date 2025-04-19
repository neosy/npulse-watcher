package persistence

import (
	"context"

	dpulse "git.n-hub.ru/neosy/npulse-watcher/domain"
)

// PulseStateRepository defines methods for managing PulseState records.
type PulseStateRepository interface {
	// Add saves a PulseState record.
	// Returns an error if the operation fails.
	Add(ctx context.Context, pulseState *dpulse.PulseState) error
	// AddSuccess saves ip as success.
	AddSuccess(ctx context.Context, ip string) error
	// RemoveSuccess unmarks the given IP as successful.
	RemoveSuccess(ctx context.Context, ip string) error
	// RemoveAllSuccess clears all successful IPs.
	RemoveAllSuccess(ctx context.Context) error

	// FindByIP retrieves the PulseState by the given IP address.
	FindByIP(ctx context.Context, ip string) (*dpulse.PulseState, error)
	// GetAllSuccessful returns all successful ips
	GetAllSuccessful(ctx context.Context) ([]string, error)
}
