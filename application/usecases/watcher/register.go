package watcher

import (
	"context"
	"errors"

	appdto "git.n-hub.ru/neosy/npulse-watcher/application/dto"
	dpulse "git.n-hub.ru/neosy/npulse-watcher/domain"
)

func (u *Watcher) Register(ctx context.Context, req *appdto.RegisterRequest) (*dpulse.PulseState, error) {
	err := u.validateRegister(ctx, req)
	if err != nil {
		return nil, err
	}

	pulseState, err := u.pulseState.Add(ctx, req.IPAddress, req.HostName)

	return pulseState, err
}

func (u *Watcher) validateRegister(ctx context.Context, req *appdto.RegisterRequest) error {
	if req == nil {
		u.logger.Error(
			"Пустой указатель в функции Register",
			"registerRequest", req,
		)

		return errors.New("function parameter is a null pointer")
	}

	return nil
}
