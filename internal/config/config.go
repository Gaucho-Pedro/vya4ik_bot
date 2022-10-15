package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"sync"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	LogLevel string `env:"LOG_LEVEL" env-default:"INFO"`
	BotDebug bool   `env:"BOT_DEBUG" env-default:"false"`
	BotToken string `env:"BOT_TOKEN" env-required:"true"`
	Cron     string `env:"CRON" env-default:"*/1 11-12 * * *"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		log.Debug("Reading configuration...")
		if err := cleanenv.ReadEnv(instance); err != nil {
			log.Fatal(err)
		}
	})
	return instance
}
