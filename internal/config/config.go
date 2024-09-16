package config

import (
	"fmt"
	"log"
	"os"

	"github.com/caarlos0/env/v11"
)

// Основные настройки
type Config struct {
	Name    string `env:"APP_NAME" envDefault:"nPulseWatcher"`
	Server  ServerConfig
	Watcher WatcherConfig
}

// Настройки HTTP сервера
type ServerConfig struct {
	Address  string `env:"SERVER_ADDRESS" envDefault:""`
	Port     int    `env:"SERVER_PORT" envDefault:"8080"`
	Compress bool   `env:"SERVER_COMPRESS" envDefault:"true"`
}

type WatcherConfig struct {
	Freq          int    `env:"WATCHER_FREQ" envDefault:"5"` // Частота проверок (в минутах)
	LogFolderPath string `env:"WATCHER_LOG_FOLDERPATH" envDefault:"~/log/nPulseWatcher"`
	LogFileName   string `env:"WATCHER_LOG_FILENAME" envDefault:"nPulse_watcher.log"`
	Telegram      TelegramConfig
}

type TelegramConfig struct {
	Token      string `env:"WATCHER_TELEGRAM_TOKEN" envDefault:""`
	TokenFile  string `env:"WATCHER_TELEGRAM_TOKEN_FILE" envDefault:"/run/secrets/telegram_token"`
	ChatId     string `env:"WATCHER_TELEGRAM_CHATID" envDefault:""`
	ChatIdFile string `env:"WATCHER_TELEGRAM_CHATID_FILE" envDefault:"/run/secrets/telegram_chatid"`
}

// Создание объекта Config
func New() *Config {
	c := &Config{}

	c.load()

	tg := &c.Watcher.Telegram

	if tg.Token == "" {
		tg.Token = SecretFileRead(tg.TokenFile)
	}
	if tg.ChatId == "" {
		tg.ChatId = SecretFileRead(tg.ChatIdFile)
	}

	return c
}

// Load config from environment variables
func (config *Config) load() {
	if err := env.Parse(config); err != nil {
		log.Fatalf("Config load(). Read configuration error: %s\n", err)
	}
}

// Загрузка данных из файла secret
func SecretFileRead(name string) string {
	data, err := os.ReadFile(name)
	if err != nil {
		log.Panic(
			fmt.Sprintf("Can't read secret file %v", name),
			err,
		)
		return ""
	}

	return string(data)
}
