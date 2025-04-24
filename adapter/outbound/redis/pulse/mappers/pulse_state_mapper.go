package mappers

import (
	"git.n-hub.ru/neosy/npulse-shared/utils/convertor"
	epulse "git.n-hub.ru/neosy/npulse-watcher/adapter/outbound/redis/pulse/entity"
	dpulse "git.n-hub.ru/neosy/npulse-watcher/domain"
)

func (m *Mappers) MapPulseStateEntityToDomain(ePulseState *epulse.PulseState) (*dpulse.PulseState, error) {
	lastTime, err := convertor.ParseTime(ePulseState.LastTime)
	if err != nil {
		return nil, err
	}

	return &dpulse.PulseState{
		IPAddress:  ePulseState.IPAddress,
		HostName:   ePulseState.HostName,
		LastTime:   lastTime,
		Status:     ePulseState.Status,
		IsNotified: ePulseState.IsNotified,
	}, nil
}

func (m *Mappers) MapPulseStateDomainToEntity(pulseState *dpulse.PulseState) *epulse.PulseState {
	return &epulse.PulseState{
		IPAddress:  pulseState.IPAddress,
		HostName:   pulseState.HostName,
		LastTime:   convertor.ConvertTimeToString(pulseState.LastTime, nil),
		Status:     pulseState.Status,
		IsNotified: pulseState.IsNotified,
	}
}
