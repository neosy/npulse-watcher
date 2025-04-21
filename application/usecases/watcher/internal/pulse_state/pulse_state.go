package pulsestate

import (
	"log/slog"

	"git.n-hub.ru/neosy/npulse-watcher/port/persistence"
)

type PulseState struct {
	logger *slog.Logger

	// Repositories
	pulseStateRep persistence.PulseStateRepository
}

func NewPulseState(
	logger *slog.Logger,

	// Repositories
	pulseStateRep persistence.PulseStateRepository,
) *PulseState {
	return &PulseState{
		logger: logger,

		// Repositories
		pulseStateRep: pulseStateRep,
	}
}
