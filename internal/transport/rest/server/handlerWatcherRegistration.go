package server

import (
	"encoding/json"
	"fmt"

	"git.n-hub.ru/neosy/npulse-watcher/internal/models"
	"github.com/neosy/gofw/nbasic"
	"github.com/neosy/gofw/nfasthttp"
	"github.com/valyala/fasthttp"
)

// Обработчик /payment/create
func (server *HTTPServer) handlerWatcherRegistration(ctx *fasthttp.RequestCtx) {
	if string(ctx.Method()) == nfasthttp.GET {
		body := ctx.PostBody()
		var watcherRegData WatcherRegRequest
		if err := json.Unmarshal(body, &watcherRegData); err != nil {
			nfasthttp.ResponseFailStandard(ctx, fasthttp.StatusBadRequest, "Error when parsing JSON")
			return
		}
		if err := watcherRegData.validate(); err != nil {
			nfasthttp.ResponseFailStandard(ctx, fasthttp.StatusBadRequest, "Error when validate JSON")
			return
		}

		// Обработка запроса
		err := server.uc.Registration(ctx, &watcherRegData.WatcherRegRequest)
		if err != nil {
			nfasthttp.ResponseFailStandard(ctx, fasthttp.StatusNotAcceptable, fmt.Sprintf("Transaction recording error. %v", err))

		} else {
			response := models.WatcherRegSuccessResponse{
				Success: true,
				Message: "OK",
			}

			responseJSON, _ := nbasic.StructToJSON(response)

			nfasthttp.ResponseSuccess(ctx, fasthttp.StatusOK, responseJSON)
		}
	} else {
		nfasthttp.ResponseFailStandard(ctx, fasthttp.StatusMethodNotAllowed, "Method not allowed")
	}
}
