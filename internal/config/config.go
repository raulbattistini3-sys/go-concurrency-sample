package config

import (

	"github.com/spf13/viper"
)

var cfg *Config

type Config struct {
	Workers  WorkersConfig
	Database DatabaseConfig
}

type WorkersConfig struct {
	Count int
	Queue int
}

type DatabaseConfig struct {
	Host          string
	Port          string
	User          string
	Password      string
	Name          string
	MaxOpenConns  int
	MaxIdleConns  int
}

func init() {
	viper.SetDefault("workers.count", 4)
	viper.SetDefault("workers.queue", 1000)

	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "3306")
	viper.SetDefault("database.user", "root")
	viper.SetDefault("database.password", "")
	viper.SetDefault("database.name", "app_dev")

	viper.SetDefault("database.max_open_conns", 10)
	viper.SetDefault("database.max_idle_conns", 10)
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = &Config{
		Workers: WorkersConfig{
			Count: viper.GetInt("workers.count"),
			Queue: viper.GetInt("workers.queue"),
		},

		Database: DatabaseConfig{
			Host:         viper.GetString("database.host"),
			Port:         viper.GetString("database.port"),
			User:         viper.GetString("database.user"),
			Password:     viper.GetString("database.password"),
			Name:         viper.GetString("database.name"),
			MaxOpenConns: viper.GetInt("database.max_open_conns"),
			MaxIdleConns: viper.GetInt("database.max_idle_conns"),
		},
	}

	return nil
}

func Get() *Config {
	return cfg
}

func GetWorkers() WorkersConfig {
	return cfg.Workers
}

func GetDatabase() DatabaseConfig {
	return cfg.Database
}
