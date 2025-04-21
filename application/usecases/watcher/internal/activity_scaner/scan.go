package ascaner

import (
	"context"
	"time"
)

// Лимит для последнего отклика, когда отправляется уведомление
const limitResponseTimeout = 5 * time.Minute

// Scan checks activity of IPs and removes or resets them if inactive.
func (u *ActivityScaner) Scan(ctx context.Context, first bool) error {
	// Получаем список активных IP-адресов
	ips, err := u.pulseStateRep.GetAllActiveIPs(ctx)
	if err != nil {
		return err
	}

	for _, ip := range ips {
		// Получаем состояние по IP
		pulseState, err := u.pulseStateRep.FindByIP(ctx, ip)

		if err != nil || pulseState == nil {
			// Если состояние не найдено — удаляем IP из активных и уведомляемых
			u.pulseStateRep.RemoveActiveIP(ctx, ip)
			u.pulseStateRep.RemoveNotifyIP(ctx, ip)
			continue
		}

		// При первом запуске удаляем IP, если время последнего отклика слишком старое
		if first && time.Since(pulseState.LastTime) > limitResponseTimeout {
			u.pulseState.RemoveSoft(ctx, *pulseState)
			continue
		}

		// Если IP давно не активен — удаляем из активных и сбрасываем уведомление
		if pulseState.LastTime.Before(time.Now().Add(-u.responseTimeout)) {
			u.pulseStateRep.RemoveActiveIP(ctx, ip)
			if pulseState.IsNotified {
				// Сбрасываем признак уведомления
				u.pulseState.MarkNotified(ctx, *pulseState, false)
			}
		}
	}

	return nil
}
