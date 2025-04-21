package watcher

import (
	"context"
	"time"
)

func (u *Watcher) runIPNotifier(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(u.notifyInterval)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				u.activityScaner.Scan(ctx, false)
				u.ipNotifier.Notify(ctx)
			}
		}
	}()
}
