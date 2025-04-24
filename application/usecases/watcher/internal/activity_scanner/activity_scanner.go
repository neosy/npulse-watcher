package ascanner

import (
	"log/slog"
	"time"

	pulsestate "git.n-hub.ru/neosy/npulse-watcher/application/usecases/watcher/internal/pulse_state"
	"git.n-hub.ru/neosy/npulse-watcher/port/persistence"
)

type ActivityScanner struct {
	logger *slog.Logger

	// Config
	responseTimeout time.Duration

	// Repositories
	pulseStateRep persistence.PulseStateRepository

	// Internal
	pulseState *pulsestate.PulseState
}

func NewActivityScanner(
	logger *slog.Logger,

	// Config
	responseTimeout time.Duration,

	// Repositories
	pulseStateRep persistence.PulseStateRepository,

	// Internal
	pulseSate *pulsestate.PulseState,

) *ActivityScanner {
	return &ActivityScanner{
		logger: logger,

		// Config
		responseTimeout: responseTimeout,

		// Repositories
		pulseStateRep: pulseStateRep,

		// Internal
		pulseState: pulseSate,
	}
}
