package handlers

import (
	watcherh "git.n-hub.ru/neosy/npulse-watcher/adapter/inbound/rest/handlers/watcher"
	"git.n-hub.ru/neosy/npulse-watcher/application/usecases"
)

type handlers struct {
	Watcher *watcherh.WatcherHandlers
}

func New(
	usecases *usecases.Usecases,
) *handlers {
	return &handlers{
		Watcher: watcherh.NewWatcherHandlers(usecases),
	}
}
