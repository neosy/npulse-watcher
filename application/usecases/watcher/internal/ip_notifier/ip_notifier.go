package ipnotifier

import (
	"log/slog"

	pulsestate "git.n-hub.ru/neosy/npulse-watcher/application/usecases/watcher/internal/pulse_state"
	"git.n-hub.ru/neosy/npulse-watcher/pkg/ntelegram"
	"git.n-hub.ru/neosy/npulse-watcher/port/persistence"
)

type IPNotifier struct {
	logger *slog.Logger

	// Config
	telegramChannelId string

	// Repositories
	pulseStateRep persistence.PulseStateRepository

	// Services
	telegram *ntelegram.Telegram

	// Internal
	pulseState *pulsestate.PulseState
}

func NewIPNotifier(
	logger *slog.Logger,

	// Config
	telegramToken string,
	telegramChannelId string,

	// Repositories
	pulseStateRep persistence.PulseStateRepository,
) *IPNotifier {
	return &IPNotifier{
		logger: logger,

		// Config
		telegramChannelId: telegramChannelId,

		// Repositories
		pulseStateRep: pulseStateRep,

		// Services
		telegram: ntelegram.New(telegramToken),

		// Internal
		pulseState: pulsestate.NewPulseState(logger, pulseStateRep),
	}
}
