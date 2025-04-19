package mappers

import (
	"git.n-hub.ru/neosy/npulse-watcher/adapter/inbound/rest/handlers/watcher/dto"
	appdto "git.n-hub.ru/neosy/npulse-watcher/application/dto"
	dpulse "git.n-hub.ru/neosy/npulse-watcher/domain"
)

func (m *Mappers) MapRegisterRequestToUsecase(req *dto.RegisterRequest) *appdto.RegisterRequest {
	return &appdto.RegisterRequest{
		IPAddress: req.IPAddress,
		HostName:  req.HostName,
	}
}

func (m *Mappers) MapPulseStateDomainToRegisterResponse(pulseState *dpulse.PulseState) *dto.RegisterResponse {
	status := dto.RegisterStatusSuccess

	if pulseState == nil {
		status = dto.RegisterStatusFailed
	}

	return &dto.RegisterResponse{
		Status: status,
	}
}
