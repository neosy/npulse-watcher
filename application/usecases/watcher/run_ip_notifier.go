package watcher

import (
	"context"
	"time"
)

func (u *Watcher) runScanner(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(u.scanInterval)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				u.activityScanner.Scan(ctx, false)
				u.ipNotifier.Notify(ctx)
			}
		}
	}()
}
