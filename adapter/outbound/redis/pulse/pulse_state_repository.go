package rpulse

import (
	"context"
	"fmt"

	"git.n-hub.ru/neosy/npulse-watcher/adapter/outbound/redis/pulse/mappers"
	dpulse "git.n-hub.ru/neosy/npulse-watcher/domain"
	"github.com/redis/go-redis/v9"
)

// PulseStateRepository
type PulseStateRepository struct {
	client    redis.UniversalClient
	keyPrefix string
	mappers   *mappers.Mappers
}

// getBaseKey generates the base Redis key.
func (r *PulseStateRepository) getBaseKey() string {
	return fmt.Sprintf("%s:pulse_state", r.keyPrefix)
}

// NewPulseStateRepository returns a new PulseStateRepository instance.
func NewPulseStateRepository(
	client redis.UniversalClient,
	keyPrefix string,
) *PulseStateRepository {
	return &PulseStateRepository{
		client:    client,
		keyPrefix: keyPrefix,
		mappers:   mappers.NewMappers(),
	}
}

// Add saves a PulseState record.
// Returns an error if the operation fails.
func (r *PulseStateRepository) Add(ctx context.Context, pulseState *dpulse.PulseState) error {
	return nil
}
