package handlers

import (
	watcherhandlers "git.n-hub.ru/neosy/npulse-watcher/adapter/inbound/rest/handlers/watcher"
	"git.n-hub.ru/neosy/npulse-watcher/application/usecases"
)

type handlers struct {
	Watcher *watcherhandlers.WatcherHandlers
}

func New(
	usecases *usecases.Usecases,
) *handlers {
	return &handlers{
		Watcher: watcherhandlers.NewWatcherHandlers(usecases),
	}
}
