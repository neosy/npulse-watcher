package mappers

import "git.n-hub.ru/neosy/npulse-watcher/adapter/inbound/rest/handlers/watcher/dto"

func (m *Mappers) MapPingRequestToText(req *dto.PingRequest) string {
	return req.Text
}

func (m *Mappers) MapTextToPingResponse(text string) *dto.PingResponse {
	if text == "Ping" {
		text = "Pong"
	}

	return &dto.PingResponse{
		Text: text,
	}
}
