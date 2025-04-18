package watcherhandlers

import (
	"git.n-hub.ru/neosy/npulse-watcher/adapter/inbound/rest/handlers/watcher/mappers"
	"git.n-hub.ru/neosy/npulse-watcher/application/usecases"
)

type WatcherHandlers struct {
	mappers  *mappers.Mappers
	usecases *usecases.Usecases
}

func NewWatcherHandlers(
	usecases *usecases.Usecases,
) *WatcherHandlers {
	return &WatcherHandlers{
		mappers:  mappers.NewMappers(),
		usecases: usecases,
	}
}
