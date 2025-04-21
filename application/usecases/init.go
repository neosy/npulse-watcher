package usecases

import "context"

func (u *Usecases) Init(ctx context.Context) {
	u.Watcher.Init(ctx)
}
