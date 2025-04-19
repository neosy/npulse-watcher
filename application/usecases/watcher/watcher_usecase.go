package watcher

import (
	"log/slog"

	"git.n-hub.ru/neosy/npulse-watcher/application/usecases/watcher/mappers"
	"git.n-hub.ru/neosy/npulse-watcher/port/persistence"
)

type Watcher struct {
	logger *slog.Logger
	mapper *mappers.Mappers

	// repositories
	pulseStateRep persistence.PulseStateRepository
}

func NewWatcher(
	logger *slog.Logger,

	// repositories
	pulseStateRep persistence.PulseStateRepository,
) (u *Watcher) {
	return &Watcher{
		logger: logger,
		mapper: mappers.NewMappers(),

		// repositories
		pulseStateRep: pulseStateRep,
	}
}
