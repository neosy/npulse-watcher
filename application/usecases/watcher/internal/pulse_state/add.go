package pulsestate

import (
	"context"
	"errors"
	"time"

	dpulse "git.n-hub.ru/neosy/npulse-watcher/domain"
	dtypes "git.n-hub.ru/neosy/npulse-watcher/domain/types"
)

// Add.
func (u *PulseState) Add(
	ctx context.Context,
	ipAddress string,
	hostName string,
) (*dpulse.PulseState, error) {
	pulseState := &dpulse.PulseState{
		IPAddress:  ipAddress,
		HostName:   hostName,
		LastTime:   time.Now(),
		Status:     dtypes.PulseStateStatusSuccess,
		IsNotified: false,
	}

	err := u.validateAdd(pulseState)
	if err != nil {
		return nil, err
	}

	pulseStateOld, err := u.pulseStateRep.FindByIP(ctx, pulseState.IPAddress)
	if err != nil {
		return nil, err
	}

	if pulseStateOld != nil {
		pulseState.IsNotified = pulseStateOld.IsNotified

		isChanged := pulseStateOld.Status != dtypes.PulseStateStatusSuccess
		if isChanged {
			pulseState.IsNotified = false
		}
	}

	err = u.pulseStateRep.Add(ctx, pulseState)
	if err != nil {
		return nil, err
	}

	pulseState, err = u.pulseStateRep.FindByIP(ctx, pulseState.IPAddress)
	if err != nil {
		return nil, err
	}

	err = u.pulseStateRep.AddActiveIP(ctx, pulseState.IPAddress)
	if err != nil {
		return nil, err
	}

	if !pulseState.IsNotified {
		u.pulseStateRep.AddNotifyIP(ctx, pulseState.IPAddress)
	}

	return pulseState, nil
}

func (u *PulseState) validateAdd(pulseState *dpulse.PulseState) error {
	if pulseState.IPAddress == "" {
		return errors.New("ip address is required")
	}

	if pulseState.HostName == "" {
		return errors.New("host name is required")
	}

	return nil
}
