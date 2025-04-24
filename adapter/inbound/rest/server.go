package httpsrv

import (
	"context"
	"fmt"
	"log/slog"

	"git.n-hub.ru/neosy/npulse-watcher/application/usecases"
	"github.com/valyala/fasthttp"
)

type Dependencies struct {
	Usecases *usecases.Usecases
}

type httpServer struct {
	logger   *slog.Logger
	usecases *usecases.Usecases
}

func NewServer(
	logger *slog.Logger,
	deps *Dependencies,
) (server *httpServer) {
	server = &httpServer{
		logger:   logger,
		usecases: deps.Usecases,
	}

	return
}

func (s *httpServer) ListenAndServe(ctx context.Context, port string) error {
	addr := fmt.Sprintf(":%s", port)

	router := s.newRouter()
	handler := newHandler(s.logger, router.Handler)

	s.logger.InfoContext(ctx, fmt.Sprintf("HTTP server listening on %s", addr))

	err := fasthttp.ListenAndServe(addr, handler)
	if err != nil {
		s.logger.ErrorContext(ctx, fmt.Sprintf("error run server: %v", err))
		return err
	}

	return nil
}

func newHandler(logger *slog.Logger, h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		logger.InfoContext(
			ctx,
			fmt.Sprintf("Request %s %s %s", ctx.Method(), ctx.Host(), ctx.RequestURI()),
		)
		h(ctx)
	}
}
