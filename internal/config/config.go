package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
	"tg_weather_bot/pkg/logging"
)

type Config struct {
	TelegramAPIKey string `yaml:"tg_api_key"`
	WeatherAPIKey  string `yaml:"weather_api_key"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application configuration")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
