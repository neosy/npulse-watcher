package iconfig

import (
	"fmt"
	"log"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

// Основные настройки
type Config struct {
	Name       string `env:"APP_NAME" envDefault:"nPulseWatcher"`
	HTTPServer HTTPServerConfig
	Watcher    WatcherConfig
	Redis      RedisConfig
}

// Настройки HTTP сервера
type HTTPServerConfig struct {
	Address  string `env:"HTTP_SERVER_ADDRESS" envDefault:""`
	Port     string `env:"HTTP_SERVER_PORT" envDefault:"8080"`
	Compress bool   `env:"HTTP_SERVER_COMPRESS" envDefault:"true"`
}

type RedisConfig struct {
	// SIMPLE(Обычный редис) or SENTINEL
	ConnectionType     string   `env:"REDIS_CONNECTION_TYPE" envDefault:"SIMPLE"`
	Addresses          []string `env:"REDIS_ADDRESSES" envSeparator:"," envDefault:"localhost:6379"`
	SentinelMasterName string   `env:"REDIS_SENTINEL_MASTER_NAME" envDefault:"mymaster"`
	PrefixKey          string   `env:"REDIS_PREFIX_KEY" envDefault:"microservice"`
}

type WatcherConfig struct {
	// Частота проверки и рассылки уведомлений (в секундах)
	FreqSec int `env:"WATCHER_FREQ" envDefault:"60"`
	// Допустимое время получения ответа от серверов (в секундах)
	ResponseDeadlineSec int    `env:"WATCHER_RESPONSE_DEADLINE" envDefault:"180"`
	LogFolderPath       string `env:"WATCHER_LOG_FOLDERPATH" envDefault:"/app_n/log"`
	LogFileName         string `env:"WATCHER_LOG_FILENAME" envDefault:"nPulse_watcher.log"`
	Telegram            TelegramConfig
}

type TelegramConfig struct {
	Token     string `env:"WATCHER_TELEGRAM_TOKEN" envDefault:""`
	TokenFile string `env:"WATCHER_TELEGRAM_TOKEN_FILE" envDefault:"/run/secrets/npulse_telegram_token"`
	ChatId    string `env:"WATCHER_TELEGRAM_CHATID" envDefault:""`
}

// Создание объекта Config
func New() *Config {
	c := &Config{}

	c.load()

	tg := &c.Watcher.Telegram

	if tg.Token == "" {
		tg.Token = SecretFileRead(tg.TokenFile)
	}

	return c
}

// Load config from environment variables
func (config *Config) load() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file, proceeding with environment variables only")
	}
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
