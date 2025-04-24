package watcherh

import (
	"encoding/json"
	"fmt"

	nfasthttp "git.n-hub.ru/neosy/npulse-shared/fasthttp"
	"git.n-hub.ru/neosy/npulse-watcher/adapter/inbound/rest/handlers/watcher/dto"
	"github.com/valyala/fasthttp"
)

// @Router /watcher/reg [get]
func (h *WatcherHandlers) RegisterHandler(ctx *fasthttp.RequestCtx) {
	/*startTime := time.Now()
	defer func() {
		elapsedTime := time.Since(startTime).Seconds()
		log.Printf("время выполнения REST метода RegisterHandler (сек): %v", elapsedTime)
	}()*/

	var reqDto = &dto.RegisterRequest{}

	err := json.Unmarshal(ctx.PostBody(), reqDto)
	if err != nil {
		nfasthttp.WriteError(ctx, fmt.Errorf("unmarshal error: %v", err), fasthttp.StatusBadRequest)
		return
	}

	req := h.mappers.MapRegisterRequestToUsecase(reqDto)

	pulseState, err := h.usecases.Watcher.Register(ctx, req)
	if err != nil {
		nfasthttp.WriteError(ctx, err, fasthttp.StatusInternalServerError)
		return
	}

	resp := h.mappers.MapPulseStateDomainToRegisterResponse(pulseState)

	nfasthttp.WriteResponse(ctx, resp)
}
