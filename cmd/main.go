package main

import (
	"context"
	"log"
	"time"

	"git.n-hub.ru/neosy/npulse-watcher/internal/config"
	"git.n-hub.ru/neosy/npulse-watcher/internal/transport/rest/server"
	usecase "git.n-hub.ru/neosy/npulse-watcher/internal/usecase"
)

var (
	cfg *config.Config
)

func init() {
	cfg = config.New()

	// TODO: Закомментировать код
	//cfg.Repo.Address = "192.168.23.12"
	//cfg.Repo.Port = 2402
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Инициализация UseCase
	uc := usecase.New(&cfg.Watcher)

	// Инициализация fastHTTTP сервера
	server, err := server.New(uc)
	server.Compress = cfg.Server.Compress

	if err != nil {
		log.Println("can't init server")
		return
	}

	// Запуск fastHTTTP сервера
	log.Printf("fastHTTTP server is running on %s:%d", cfg.Server.Address, cfg.Server.Port)
	server.RunServer(cfg.Server.Address, cfg.Server.Port)

	<-ctx.Done()

	log.Println("application shutdown")
}
