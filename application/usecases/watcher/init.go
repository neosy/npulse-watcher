package watcher

import (
	"context"
)

func (u *Watcher) Init(ctx context.Context) error {
	err := u.activityScaner.Scan(ctx, true)
	u.runIPNotifier(ctx)

	return err
}
