package iconfig

import (
	"fmt"
	"log"

	nlogger "git.n-hub.ru/neosy/npulse-shared/logger"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

// Основные настройки
type Config struct {
	AppName string `env:"APP_NAME" envDefault:"nPulseWatcher"`
	// values: debug, info, warn, error
	LogLevel nlogger.LogLevel `env:"LOG_LEVEL" envDefault:"warn"`

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
	// Interval between checks and run scanner (in seconds)
	ScanInterval uint16 `env:"WATCHER_SCAN_INTERVAL" envDefault:"60"`
	// Maximum allowed response time from the host before it's considered unreachable (in seconds)
	ResponseTimeout uint16 `env:"WATCHER_RESPONSE_TIMEOUT" envDefault:"180"`

	LogFolderPath string `env:"WATCHER_LOG_FOLDERPATH" envDefault:"/app_n/log"`
	LogFileName   string `env:"WATCHER_LOG_FILENAME" envDefault:"nPulse_watcher.log"`

	Telegram TelegramConfig
}

type TelegramConfig struct {
	Token     string `env:"WATCHER_TELEGRAM_TOKEN" envDefault:""`
	TokenFile string `env:"WATCHER_TELEGRAM_TOKEN_FILE" envDefault:"/run/secrets/npulse_telegram_token"`
	ChatId    string `env:"WATCHER_TELEGRAM_CHAT_ID" envDefault:""`
}

// Создание объекта Config
func New() *Config {
	c := &Config{}

	c.load()

	c.LoadTelegramToken()

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
