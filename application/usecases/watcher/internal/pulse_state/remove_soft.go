package pulsestate

import (
	"context"

	dpulse "git.n-hub.ru/neosy/npulse-watcher/domain"
	dtypes "git.n-hub.ru/neosy/npulse-watcher/domain/types"
)

// RemoveSoft marks the pulse as failed, sets notified, and cleans up related IPs.
func (u *PulseState) RemoveSoft(ctx context.Context, pulseState dpulse.PulseState) {
	pulseState.Status = dtypes.PulseStateStatusFailed
	pulseState.IsNotified = true
	u.pulseStateRep.Update(ctx, &pulseState)

	u.pulseStateRep.RemoveActiveIP(ctx, pulseState.IPAddress)
	u.pulseStateRep.RemoveNotifyIP(ctx, pulseState.IPAddress)
}
