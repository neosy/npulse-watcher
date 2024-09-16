// Пакет User cases
package uc

import (
	"context"
	"time"

	"git.n-hub.ru/neosy/npulse-watcher/internal/config"
	"git.n-hub.ru/neosy/npulse-watcher/internal/models"
	wc "git.n-hub.ru/neosy/npulse-watcher/internal/pkg/watchercomp"
)

type UseCase struct {
	config      *config.WatcherConfig
	watcherComp *wc.WatcherComp
}

func New(config *config.WatcherConfig) *UseCase {
	uc := &UseCase{
		config: config,
	}

	uc.watcherComp = wc.New(uc.config.LogFileName, uc.config.LogFolderPath, true)
	uc.watcherComp.TelegramConfigSet(uc.config.Telegram.Token, uc.config.Telegram.ChatId)

	return uc
}

func (uc *UseCase) Registration(ctx context.Context, data *models.WatcherRegRequest) error {

	_, err := uc.watcherComp.Add(data.ServerIP, data.ServerName, time.Now(), wc.WatchStatusOK)

	return err
}

func (uc *UseCase) Daemon() {
	for {
		time.Sleep(5 * time.Second)
		//time.Sleep(time.Duration(uc.config.Freq) * time.Minute)
		uc.watcherComp.Check(10)
	}
}
