package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"time"
)

type environment string

const (
	LocalEnv environment = "local"
	DevEnv   environment = "dev"
	ProdEnv  environment = "prod"
)

type (
	server struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
	}

	database struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"name"`
	}

	jwt struct {
		// Время жизни токена
		AccessTokenLifeTime time.Duration `mapstructure:"access_token_lifetime"`
		// Время жизни токена для обновление основного токена
		RefreshTokenLifeTime time.Duration `mapstructure:"refresh_token_lifetime"`
		TokenKey             string        `mapstructure:"-" env:"TOKEN_KEY,required"`
		RefreshTokenKey      string        `mapstructure:"-" env:"REFRESH_TOKEN_KEY,required"`
	}

	Config struct {
		Env      environment `mapstructure:"env"`
		Server   server      `mapstructure:"server"`
		Database database    `mapstructure:"database"`
		JWT      jwt         `mapstructure:"jwt"`
	}
)

func MustGetConfig(path string) Config {
	var config Config
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed to read config file: %s", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("failed to unmarshal config file: %s", err)
	}

	err := godotenv.Load(path + "/.env")
	if err != nil {
		log.Fatalf("failed to read enviroment file: %s", err)
	}

	err = env.Parse(&config)
	if err != nil {
		log.Fatalf("failed to parse enviroment file: %s", err)
	}

	return config
}
