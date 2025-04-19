package mappers

import (
	"time"

	appdto "git.n-hub.ru/neosy/npulse-watcher/application/dto"
	dpulse "git.n-hub.ru/neosy/npulse-watcher/domain"
	dtypes "git.n-hub.ru/neosy/npulse-watcher/domain/types"
)

func (m *Mappers) MapRegisterRequestToPulseStateDomain(
	req *appdto.RegisterRequest,
	pulseStateOld *dpulse.PulseState,
) *dpulse.PulseState {
	isNotified := false

	if pulseStateOld != nil {
		isNotified = pulseStateOld.IsNotified

		isChanged := pulseStateOld.Status != dtypes.PulseStateStatusSuccess
		if isChanged {
			isNotified = false
		}
	}

	return &dpulse.PulseState{
		IPAddress:  req.IPAddress,
		HostName:   req.HostName,
		LastTime:   time.Now(),
		Status:     dtypes.PulseStateStatusSuccess,
		IsNotified: isNotified,
	}
}
