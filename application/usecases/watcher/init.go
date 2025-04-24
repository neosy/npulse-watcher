package watcher

import (
	"context"
)

func (u *Watcher) Init(ctx context.Context) error {
	err := u.activityScanner.Scan(ctx, true)
	u.runScanner(ctx)

	return err
}
