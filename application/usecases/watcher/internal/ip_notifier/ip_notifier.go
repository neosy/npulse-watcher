package ipnotifier

import (
	"log/slog"

	ntelegram "git.n-hub.ru/neosy/npulse-shared/telegram"
	pulsestate "git.n-hub.ru/neosy/npulse-watcher/application/usecases/watcher/internal/pulse_state"
	"git.n-hub.ru/neosy/npulse-watcher/port/persistence"
)

type Config struct {
	TelegramToken  string
	TelegramChatId string
}

type IPNotifier struct {
	logger *slog.Logger

	// Config
	config *Config

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
	config *Config,

	// Repositories
	pulseStateRep persistence.PulseStateRepository,

	// Internal
	pulseState *pulsestate.PulseState,
) *IPNotifier {
	return &IPNotifier{
		logger: logger,

		// Config
		config: config,

		// Repositories
		pulseStateRep: pulseStateRep,

		// Services
		telegram: ntelegram.New(config.TelegramToken),

		// Internal
		pulseState: pulseState,
	}
}
