package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	nlogger "git.n-hub.ru/neosy/npulse-shared/logger"
	redissh "git.n-hub.ru/neosy/npulse-shared/redis"
	httpsrv "git.n-hub.ru/neosy/npulse-watcher/adapter/inbound/rest"
	rrepositories "git.n-hub.ru/neosy/npulse-watcher/adapter/outbound/redis"
	"git.n-hub.ru/neosy/npulse-watcher/application/usecases"
	"git.n-hub.ru/neosy/npulse-watcher/application/usecases/watcher"
	iconfig "git.n-hub.ru/neosy/npulse-watcher/infrastructure/config"
)

const (
	// Максимальное количество повторений
	redisMaxRetries = 10
	// Задержка перед повторением (сек)
	redisRetryInterval = 1
)

func main() {
	cfg := iconfig.New()

	ctx, cancel := context.WithCancel(context.Background())

	// Создаем обработчик с уровнем Info, используя HandlerOptions
	handlerOptions := &slog.HandlerOptions{
		// Устанавливаем уровень логирования
		Level: nlogger.LevelToSlogLevel(cfg.AppConfig.LogLevel),
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, handlerOptions))

	// Redis
	conType, err := redissh.ParseConnectionType(cfg.Redis.ConnectionType)
	if err != nil {
		logger.ErrorContext(ctx, fmt.Sprintf("Failed to parse Redis connection type: %v", err))
		return
	}

	redisClient, err := redissh.NewRedisClientBuilder().
		SetConnectionType(conType).
		SetRedisAddresses(cfg.Redis.Addresses).
		SetSentinelMasterName(cfg.Redis.SentinelMasterName).
		SetRetries(redisMaxRetries).
		SetRetryInterval(redisRetryInterval * time.Second).
		SetRouteByLatency(true).
		Build()
	if err != nil {
		logger.ErrorContext(ctx, fmt.Sprintf("Failed to initialize Redis client: %v", err))
		return
	}
	defer redisClient.Close()

	// Redis repositories
	rRepositories := rrepositories.New(redisClient, cfg.Redis.PrefixKey)

	// Usecases
	watcherConfig := &watcher.Config{
		ScanInterval:    time.Duration(cfg.Watcher.ScanInterval) * time.Second,
		ResponseTimeout: time.Duration(cfg.Watcher.ResponseTimeout) * time.Second,
		TelegramToken:   cfg.Watcher.Telegram.Token,
		TelegramChatId:  cfg.Watcher.Telegram.ChatId,
	}
	ucDeps := &usecases.Dependencies{
		Repositories: usecases.DepRepositories{
			PulseState: rRepositories.PulseState,
		},
	}
	uc := usecases.New(logger, watcherConfig, ucDeps)
	// Initialize
	uc.Init(ctx)

	// Захват сигналов завершения (Ctrl+C, SIGTERM)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// FastHTTP server
	go func(ctx context.Context) {
		deps := &httpsrv.Dependencies{
			Usecases: uc,
		}

		httpServer := httpsrv.NewServer(logger, deps)
		err = httpServer.ListenAndServe(ctx, cfg.HTTPServer.Port)
		if err != nil {
			cancel()
		}
	}(ctx)

	// Ждем сигнал завершения или отмены контекста
	select {
	case <-ctx.Done():
		logger.ErrorContext(ctx, "Context complete, shutting down services...")
	case sig := <-sigChan:
		logger.ErrorContext(ctx, fmt.Sprintf("Signal received: %v, shutting down...", sig))
		cancel()
	}

	if err != nil {
		log.Print(err)
	}
}
