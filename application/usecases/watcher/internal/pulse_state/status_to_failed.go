package pulsestate

import (
	"context"

	dpulse "git.n-hub.ru/neosy/npulse-watcher/domain"
	dtypes "git.n-hub.ru/neosy/npulse-watcher/domain/types"
)

// ToFailed.
func (u *PulseState) ToFailed(ctx context.Context, pulseState dpulse.PulseState) error {
	if pulseState.Status != dtypes.PulseStateStatusFailed {
		pulseState.Status = dtypes.PulseStateStatusFailed
		pulseState.IsNotified = false

		err := u.pulseStateRep.Update(ctx, &pulseState)
		if err != nil {
			return err
		}
	}

	if !pulseState.IsNotified {
		u.pulseStateRep.AddNotifyIP(ctx, pulseState.IPAddress)
	}

	u.pulseStateRep.RemoveActiveIP(ctx, pulseState.IPAddress)

	return nil
}
