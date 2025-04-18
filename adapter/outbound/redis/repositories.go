package rrepositories

import (
	rpulse "git.n-hub.ru/neosy/npulse-watcher/adapter/outbound/redis/pulse"
	"github.com/redis/go-redis/v9"
)

// Repositories groups all database repositories.
type Repositories struct {
	PulseState *rpulse.PulseStateRepository
}

// New returns a new Repositories struct with database connections.
func New(client redis.UniversalClient, keyPrefix string) *Repositories {
	return &Repositories{
		// TODO Заменить appName на значение из shared, когда будет возможно
		PulseState: rpulse.NewPulseStateRepository(client, keyPrefix),
	}
}
