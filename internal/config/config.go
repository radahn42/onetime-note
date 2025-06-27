package config

import (
	"github.com/spf13/viper"
	"log"
	"strings"
)

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
}

type AppConfig struct {
	Addr string `mapstructure:"addr"`
}

type Config struct {
	Redis RedisConfig `mapstructure:"redis"`
	App   AppConfig   `mapstructure:"app"`
}

func MustLoad() *Config {
	viper.SetDefault("app.addr", ":8080")
	viper.SetDefault("redis.addr", "localhost:6379")
	viper.SetDefault("redis.password", "")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("No config file found, relying on environment variables: %v", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		panic("failed to unmarshal config: " + err.Error())
	}

	return &cfg
}
