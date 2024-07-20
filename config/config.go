package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress       string
	ConsulAddress       string
	UserServiceAddress  string
	OrderServiceAddress string
}

func LoadConfig() *Config {
	viper.SetConfigFile("config.yaml")
	viper.ReadInConfig()

	return &Config{
		ServerAddress:       viper.GetString("server.address"),
		ConsulAddress:       viper.GetString("consul.address"),
		UserServiceAddress:  viper.GetString("services.user.address"),
		OrderServiceAddress: viper.GetString("services.order.address"),
	}
}
