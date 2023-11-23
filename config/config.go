package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	App   AppConfig   `json:"app"`
	DB    DBConfig    `json:"db"`
	Redis RedisConfig `json:"redis"`
	Asynq AsynqConfig `json:"asynq"`
}

type DBConfig struct {
	Dialect      string `json:"dialect"`
	Datasource   string `json:"datasource"`
	MaxIdleConns int    `json:"max_idle_conns"`
	MaxOpenConns int    `json:"max_open_conns"`
}

type AppConfig struct {
	Addr string `json:"addr"`
}

type RedisConfig struct {
	Address string `json:"address"`
}

type AsynqConfig struct {
	Concurrency int `json:"concurrency"`
}

var (
	App   AppConfig
	Redis RedisConfig
	Asynq AsynqConfig
)

func SetupConfig() (config *Config) {
	config = &Config{}

	viper.AddConfigPath("config/")
	viper.SetConfigName("config_files")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	_ = viper.Unmarshal(config)

	App = config.App
	Redis = config.Redis
	Asynq = config.Asynq
	return
}
