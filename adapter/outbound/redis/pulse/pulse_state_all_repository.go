package rpulse

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	epulse "git.n-hub.ru/neosy/npulse-watcher/adapter/outbound/redis/pulse/entity"
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
	if pulseState == nil {
		return errors.New("function parameter is a null pointer")
	}

	ipStateKey := fmt.Sprintf("%s:%s:%s", r.getBaseKey(), "all", pulseState.IPAddress)

	ePulseState := r.mappers.MapPulseStateDomainToEntity(pulseState)

	pulseStateBytes, err := json.Marshal(ePulseState)
	if err != nil {
		return fmt.Errorf("failed to marshal PulseState: %v", err)
	}

	err = r.client.Set(ctx, ipStateKey, pulseStateBytes, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to set value in Redis: %v", err)
	}

	return nil
}

// FindByIP retrieves the PulseState by the given IP address.
func (r *PulseStateRepository) FindByIP(ctx context.Context, ip string) (*dpulse.PulseState, error) {
	ipStateKey := fmt.Sprintf("%s:%s:%s", r.getBaseKey(), "all", ip)

	pulseStateBytes, err := r.client.Get(ctx, ipStateKey).Bytes()
	if err != nil {
		if err == redis.Nil {
			// Если данных нет (ошибка Nil)
			return nil, nil
		}
		// Если другая ошибка, то возвращаем ошибку
		return nil, err
	}

	var ePulseState = &epulse.PulseState{}
	err = json.Unmarshal(pulseStateBytes, ePulseState)
	if err != nil {
		return nil, fmt.Errorf("unmarshaling error of paymentMethod from Redis: %v", err)
	}

	pulseState, err := r.mappers.MapPulseStateEntityToDomain(ePulseState)
	if err != nil {
		return nil, fmt.Errorf("failed to map PulseState to domain %v", err)
	}

	return pulseState, nil
}
