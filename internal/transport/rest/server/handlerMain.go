// Пакет для сервера fastHTTP - Основной обработчик
package server

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

// Основной обработчик сервера FastHTTP
func (s *HTTPServer) HandlerMain(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/watcher/reg":
		s.handlerWatcherRegistration(ctx)
	default:
		ctx.Error(fmt.Sprintf("Unsupported path %q", ctx.Path()), fasthttp.StatusNotFound)
	}
}
