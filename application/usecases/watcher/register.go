package watcher

import (
	"context"
	"errors"

	appdto "git.n-hub.ru/neosy/npulse-watcher/application/dto"
	dpulse "git.n-hub.ru/neosy/npulse-watcher/domain"
	dtypes "git.n-hub.ru/neosy/npulse-watcher/domain/types"
)

func (u *Watcher) Register(ctx context.Context, req *appdto.RegisterRequest) (*dpulse.PulseState, error) {
	err := u.validateRegister(ctx, req)
	if err != nil {
		return nil, err
	}

	pulseStateOld, err := u.pulseStateRep.FindByIP(ctx, req.IPAddress)
	if err != nil {
		return nil, err
	}

	pulseState := u.mapper.MapRegisterRequestToPulseStateDomain(req, pulseStateOld)

	err = u.pulseStateRep.Add(ctx, pulseState)
	if err != nil {
		return nil, err
	}

	if pulseState.Status == dtypes.PulseStateStatusSuccess {
		err = u.pulseStateRep.AddSuccess(ctx, pulseState.IPAddress)
		if err != nil {
			return nil, err
		}
	}

	return pulseState, nil
}

func (u *Watcher) validateRegister(ctx context.Context, req *appdto.RegisterRequest) error {
	if req == nil {
		u.logger.Error(
			"Пустой указатель в функции Register",
			"registerRequest", req,
		)

		return errors.New("function parameter is a null pointer")
	}

	if req.IPAddress == "" {
		return errors.New("ip address is required")
	}

	if req.HostName == "" {
		return errors.New("host name is required")
	}

	return nil
}
