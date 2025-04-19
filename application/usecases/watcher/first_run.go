package watcher

import (
	"context"
)

func (u *Watcher) FirstRun(ctx context.Context) error {
	return u.pulseStateRep.RemoveAllSuccess(ctx)
}
