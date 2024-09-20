package server

import (
	"encoding/json"
	"fmt"

	"git.n-hub.ru/neosy/npulse-watcher/internal/models"
	"github.com/neosy/gofw/nbasic"
	"github.com/neosy/gofw/nfasthttp"
	"github.com/valyala/fasthttp"
)

// Обработчик /watcher/ping
func (server *HTTPServer) handlerWatcherPing(ctx *fasthttp.RequestCtx) {
	if string(ctx.Method()) == nfasthttp.GET {
		body := ctx.PostBody()

		var watcherPingData WatcherPingRequest
		if err := json.Unmarshal(body, &watcherPingData); err != nil {
			nfasthttp.ResponseFailDefault(ctx, fasthttp.StatusBadRequest, "Error when parsing JSON")
			return
		}
		if err := watcherPingData.validate(); err != nil {
			nfasthttp.ResponseFailDefault(ctx, fasthttp.StatusBadRequest, "Error when validate JSON")
			return
		}

		// Ответ
		if watcherPingData.Text != "Ping" {
			nfasthttp.ResponseFailDefault(ctx, fasthttp.StatusNotAcceptable, fmt.Sprintf("Ping error. %v", nil))

		} else {
			response := models.WatcherPingResponse{
				Text: "Pong",
			}

			responseJSON, _ := nbasic.StructToJSON(response)

			nfasthttp.ResponseSuccessOK(ctx, responseJSON)
		}
	} else {
		nfasthttp.ResponseFailDefault(ctx, fasthttp.StatusMethodNotAllowed, "Method not allowed")
	}
}
