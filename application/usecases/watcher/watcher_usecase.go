package watcher

import (
	"log/slog"
	"time"

	ascanner "git.n-hub.ru/neosy/npulse-watcher/application/usecases/watcher/internal/activity_scanner"
	ipnotifier "git.n-hub.ru/neosy/npulse-watcher/application/usecases/watcher/internal/ip_notifier"
	pulsestate "git.n-hub.ru/neosy/npulse-watcher/application/usecases/watcher/internal/pulse_state"
	"git.n-hub.ru/neosy/npulse-watcher/application/usecases/watcher/mappers"
	"git.n-hub.ru/neosy/npulse-watcher/port/persistence"
)

type Config struct {
	// Interval between checks and run scanner
	ScanInterval time.Duration
	// Maximum allowed response time from the host before it's considered unreachable
	ResponseTimeout time.Duration
	//
	TelegramToken string
	//
	TelegramChatId string
}

type Watcher struct {
	logger *slog.Logger
	mapper *mappers.Mappers

	// Config
	// Interval between checks and run scanner
	scanInterval time.Duration

	// Repositories
	pulseStateRep persistence.PulseStateRepository

	// Internal
	pulseState      *pulsestate.PulseState
	activityScanner *ascanner.ActivityScanner
	ipNotifier      *ipnotifier.IPNotifier
}

func NewWatcher(
	logger *slog.Logger,

	// Config
	config *Config,

	// Repositories
	pulseStateRep persistence.PulseStateRepository,
) (u *Watcher) {
	pulseState := pulsestate.NewPulseState(logger, pulseStateRep)

	return &Watcher{
		logger: logger,
		mapper: mappers.NewMappers(),

		// Config
		scanInterval: config.ScanInterval,

		// Repositories
		pulseStateRep: pulseStateRep,

		// Internal
		pulseState:      pulseState,
		activityScanner: ascanner.NewActivityScanner(logger, config.ResponseTimeout, pulseStateRep, pulseState),
		ipNotifier: ipnotifier.NewIPNotifier(
			logger,
			&ipnotifier.Config{
				TelegramToken:     config.TelegramToken,
				TelegramChatId: config.TelegramChatId,
			},
			pulseStateRep,
			pulseState,
		),
	}
}
