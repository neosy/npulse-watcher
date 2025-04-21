package ascaner

import (
	"log/slog"
	"time"

	pulsestate "git.n-hub.ru/neosy/npulse-watcher/application/usecases/watcher/internal/pulse_state"
	"git.n-hub.ru/neosy/npulse-watcher/port/persistence"
)

type ActivityScaner struct {
	logger *slog.Logger

	// Config
	responseTimeout time.Duration

	// Repositories
	pulseStateRep persistence.PulseStateRepository

	// Internal
	pulseState *pulsestate.PulseState
}

func NewActivityScaner(
	logger *slog.Logger,

	// Config
	responseTimeout time.Duration,

	// Repositories
	pulseStateRep persistence.PulseStateRepository,

) *ActivityScaner {
	return &ActivityScaner{
		logger: logger,

		// Config
		responseTimeout: responseTimeout,

		// Repositories
		pulseStateRep: pulseStateRep,

		// Internal
		pulseState: pulsestate.NewPulseState(logger, pulseStateRep),
	}
}
