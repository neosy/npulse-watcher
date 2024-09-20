package server

import (
	"encoding/json"
	"fmt"

	"github.com/neosy/gofw/nfasthttp"
	"github.com/valyala/fasthttp"
)

// Обработчик /payment/create
func (server *HTTPServer) handlerWatcherRegistration(ctx *fasthttp.RequestCtx) {
	if string(ctx.Method()) == nfasthttp.GET {
		body := ctx.PostBody()
		var watcherRegData WatcherRegRequest
		if err := json.Unmarshal(body, &watcherRegData); err != nil {
			nfasthttp.ResponseFailDefault(ctx, fasthttp.StatusBadRequest, "Error when parsing JSON")
			return
		}
		if err := watcherRegData.validate(); err != nil {
			nfasthttp.ResponseFailDefault(ctx, fasthttp.StatusBadRequest, "Error when validate JSON")
			return
		}

		// Обработка запроса
		err := server.uc.Registration(ctx, &watcherRegData.WatcherRegRequest)
		if err != nil {
			nfasthttp.ResponseFailDefault(ctx, fasthttp.StatusNotAcceptable, fmt.Sprintf("Transaction recording error. %v", err))

		} else {
			nfasthttp.ResponseSuccessOKDefault(ctx, "OK")
		}
	} else {
		nfasthttp.ResponseFailDefault(ctx, fasthttp.StatusMethodNotAllowed, "Method not allowed")
	}
}
