// Пакет User cases
package uc

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"git.n-hub.ru/neosy/npulse-watcher/internal/config"
	"git.n-hub.ru/neosy/npulse-watcher/internal/models"
)

type UseCase struct {
	config *config.WatcherConfig
}

func New(config *config.WatcherConfig) *UseCase {
	return &UseCase{
		config: config,
	}
}

func (uc *UseCase) Registration(ctx context.Context, data *models.WatcherRegRequest) error {
	var fileName = fileNameAddYear(uc.config.LogFileName)
	var logFolderPath = uc.config.LogFolderPath
	var filePath = logFolderPath + "/" + fileName
	var status = 1

	timeNow := time.Now()

	//compIPs, err := fileFindCompIPs(filePath)

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Error opening the file:", err)
		return err
	}
	defer file.Close()

	// Форматируем строку
	line := fmt.Sprintf("%s %s %s %s %d\n", timeNow.Format("2006-01-02"), timeNow.Format("15:04:05"), data.ServerIP, data.ServerName, status)

	// Записываем строку в файл
	if _, err := file.WriteString(line); err != nil {
		log.Println("Error writing to the file:", err)
	}

	return err
}
