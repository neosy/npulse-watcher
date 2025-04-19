package mappers

import (
	"time"

	appdto "git.n-hub.ru/neosy/npulse-watcher/application/dto"
	dpulse "git.n-hub.ru/neosy/npulse-watcher/domain"
	dtypes "git.n-hub.ru/neosy/npulse-watcher/domain/types"
)

func (m *Mappers) MapRegisterRequestToPulseStateDomain(req *appdto.RegisterRequest) *dpulse.PulseState {
	return &dpulse.PulseState{
		IPAddress:  req.IPAddress,
		HostName:   req.HostName,
		LastTime:   time.Now(),
		Status:     dtypes.PulseStateStatusSuccess,
		IsNotified: false,
	}
}
