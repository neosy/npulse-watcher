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
}
