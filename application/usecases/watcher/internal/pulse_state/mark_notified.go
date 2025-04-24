package pulsestate

import (
	"context"

	dpulse "git.n-hub.ru/neosy/npulse-watcher/domain"
)

// MarkNotified updates the IsNotified flag and manages the notify IP list.
func (u *PulseState) MarkNotified(ctx context.Context, pulseState dpulse.PulseState, isNotified bool) error {
	// Устанавливаем признак уведомления
	pulseState.IsNotified = isNotified
	err := u.pulseStateRep.Update(ctx, &pulseState)

	// Добавляем ip в список уведомлений
	if pulseState.IsNotified {
		u.pulseStateRep.RemoveNotifyIP(ctx, pulseState.IPAddress)
	} else {
		u.pulseStateRep.AddNotifyIP(ctx, pulseState.IPAddress)
	}

	return err
}
