package httpsrv

import (
	"git.n-hub.ru/neosy/npulse-watcher/adapter/inbound/rest/handlers"
	"github.com/fasthttp/router"
)

const (
	// Groups
	groupWatcher = "/watcher"

	// Ping
	pathPing = "/ping"
)

// newRouter returns a new router.
func (s *httpServer) newRouter() *router.Router {
	r := router.New()
	r.RedirectTrailingSlash = false

	handlers := handlers.New(
		s.usecases,
	)

	group := r.Group(groupWatcher)
	{
		group.GET(pathPing, handlers.Watcher.PingHandler)
	}

	return r
}
