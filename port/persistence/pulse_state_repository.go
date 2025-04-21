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
	// Update saves a PulseState record.
	Update(ctx context.Context, pulseState *dpulse.PulseState) error
	// RemoveIP remove IP address.
	RemoveIP(ctx context.Context, ip string) error
	// FindByIP retrieves the PulseState by the given IP address.
	FindByIP(ctx context.Context, ip string) (*dpulse.PulseState, error)

	// AddActiveIP
	AddActiveIP(ctx context.Context, ip string) error
	// RemoveActiveIP
	RemoveActiveIP(ctx context.Context, ip string) error
	// RemoveAllActiveIPs
	RemoveAllActiveIPs(ctx context.Context) error
	// GetAllActiveIPs
	GetAllActiveIPs(ctx context.Context) ([]string, error)

	// AddNotifyIP
	AddNotifyIP(ctx context.Context, ip string) error
	// RemoveNotifyIP
	RemoveNotifyIP(ctx context.Context, ip string) error
	// GetAllNotifyIPs
	GetAllNotifyIPs(ctx context.Context) ([]string, error)
}
