package watcher

import (
	"log/slog"
	"time"

	ascaner "git.n-hub.ru/neosy/npulse-watcher/application/usecases/watcher/internal/activity_scaner"
	ipnotifier "git.n-hub.ru/neosy/npulse-watcher/application/usecases/watcher/internal/ip_notifier"
	"git.n-hub.ru/neosy/npulse-watcher/application/usecases/watcher/mappers"
	"git.n-hub.ru/neosy/npulse-watcher/port/persistence"
)

type Config struct {
	// Interval between checks and sending notifications
	NotifyInterval time.Duration
	// Maximum allowed response time from the host before it's considered unreachable
	ResponseTimeout time.Duration
	//
	TelegramToken string
	//
	TelegramChannelId string
}

type Watcher struct {
	logger *slog.Logger
	mapper *mappers.Mappers

	// Config
	// Interval between checks and sending notifications
	notifyInterval time.Duration

	// Repositories
	pulseStateRep persistence.PulseStateRepository

	// Internal
	activityScaner *ascaner.ActivityScaner
	ipNotifier     *ipnotifier.IPNotifier
}

func NewWatcher(
	logger *slog.Logger,

	// Config
	config *Config,

	// Repositories
	pulseStateRep persistence.PulseStateRepository,
) (u *Watcher) {
	return &Watcher{
		logger: logger,
		mapper: mappers.NewMappers(),

		// Config
		notifyInterval: config.NotifyInterval,

		// Repositories
		pulseStateRep: pulseStateRep,

		// Internal
		activityScaner: ascaner.NewActivityScaner(logger, config.ResponseTimeout, pulseStateRep),
		ipNotifier:     ipnotifier.NewIPNotifier(logger, config.TelegramToken, config.TelegramChannelId, pulseStateRep),
	}
}
