package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App  AppConfig
	SMTP SMTPConfig
}

type AppConfig struct {
	Port uint16
}

type SMTPConfig struct {
	Host string
	Port uint16
	User string
	Pass string
}

var config *Config

func Init() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".config/")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("error parsing default config file")
	}

	config = &Config{
		App: AppConfig{
			Port: viper.GetUint16("service.port"),
		},
		SMTP: SMTPConfig{
			Host: viper.GetString("smtp.host"),
			Port: viper.GetUint16("smtp.port"),
			User: viper.GetString("smtp.user"),
			Pass: viper.GetString("smtp.pass"),
		},
	}
}

func GetConfig() *Config {
	return config
}
