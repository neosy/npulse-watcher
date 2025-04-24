package ipnotifier

import (
	"context"
	"fmt"
	"os"

	dpulse "git.n-hub.ru/neosy/npulse-watcher/domain"
	dtypes "git.n-hub.ru/neosy/npulse-watcher/domain/types"
)

// Notify checks IPs pending notification, marks them, and sends a message via Telegram.
func (u *IPNotifier) Notify(ctx context.Context) error {
	// Получаем список IP-адресов, по которым нужно отправить уведомления
	ips, err := u.pulseStateRep.GetAllNotifyIPs(ctx)
	if err != nil {
		return err
	}

	// Если список пуст, выходим
	if len(ips) == 0 {
		return nil
	}

	var notifyPulseStates = make([]*dpulse.PulseState, 0, len(ips))

	for _, ip := range ips {
		// Получаем состояние по IP
		pulseState, err := u.pulseStateRep.FindByIP(ctx, ip)
		if err != nil {

			continue
		}

		// Проверяем, что ещё не было уведомления
		if !pulseState.IsNotified {
			switch pulseState.Status {
			case dtypes.PulseStateStatusSuccess, dtypes.PulseStateStatusFailed:
				notifyPulseStates = append(notifyPulseStates, pulseState)
			}

			// Устанавливаем признак уведомления
			u.pulseState.MarkNotified(ctx, *pulseState, true)
		}
	}

	const ch_br = "\n"
	var msg string

	for _, pulseState := range notifyPulseStates {
		if msg == "" {
			hostName, _ := os.Hostname()
			msg = fmt.Sprintf("Server name '%s'", hostName)
			msg = msg + ch_br + "Check servers:"
		}

		// Добавляем информацию о каждом IP
		msg = msg + ch_br + fmt.Sprintf("    '%s' %s", pulseState.HostName, pulseState.IPAddress)
		switch pulseState.Status {
		case dtypes.PulseStateStatusSuccess:
			msg += " - ✓"
		case dtypes.PulseStateStatusFailed:
			msg += " - ✗"
		}
	}

	// Отправляем сообщение в Telegram
	err = u.telegram.Send(u.config.TelegramChatId, msg)
	if err != nil {
		u.logger.Warn(
			"Ошибка отправки сообщения в Telegram",
			"error", err,
		)
	}

	return err
}
