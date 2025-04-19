package watcherh

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	nfasthttp "git.n-hub.ru/neosy/npulse-shared/fasthttp"
	"git.n-hub.ru/neosy/npulse-watcher/adapter/inbound/rest/handlers/watcher/dto"
	"github.com/valyala/fasthttp"
)

// @Router /watcher/ping [get]
func (h *WatcherHandlers) PingHandler(ctx *fasthttp.RequestCtx) {
	startTime := time.Now()
	defer func() {
		elapsedTime := time.Since(startTime).Seconds()
		log.Printf("время выполнения REST метода PingHandler (сек): %v", elapsedTime)
	}()
	var req = &dto.PingRequest{}

	err := json.Unmarshal(ctx.PostBody(), req)
	if err != nil {
		nfasthttp.WriteError(ctx, fmt.Errorf("unmarshal error: %v", err), fasthttp.StatusBadRequest)

		return
	}

	txt := h.mappers.MapPingRequestToText(req)
	resp := h.mappers.MapTextToPingResponse(txt)

	nfasthttp.WriteResponse(ctx, resp)
}
