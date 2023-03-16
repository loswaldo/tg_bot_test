package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
	"tg_weather_bot/pkg/logging"
)

type APIConfig struct {
	TelegramAPIKey string `yaml:"tg_api_key"`
	WeatherAPIKey  string `yaml:"weather_api_key"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
	SSLMode  string `yaml:"ssl_mode"`
}

var apiInstance *APIConfig
var dbInstance *DBConfig
var apiOnce sync.Once
var dbOnce sync.Once

func GetAPIConfig() *APIConfig {
	apiOnce.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application configuration")
		apiInstance = &APIConfig{}
		if err := cleanenv.ReadConfig("api-config.yml", apiInstance); err != nil {
			help, _ := cleanenv.GetDescription(apiInstance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return apiInstance
}

func GetDBConfig() *DBConfig {
	dbOnce.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read database configuration")
		dbInstance = &DBConfig{}
		if err := cleanenv.ReadConfig("db-config.yml", dbInstance); err != nil {
			help, _ := cleanenv.GetDescription(dbInstance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return dbInstance
}
