package config

import (
	"flag"
	"log"

	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type Config struct {
	App  AppConfig
	SMTP SMTPConfig
}

type AppConfig struct {
	Mode string
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
	var mode string

	flag.StringVar(&mode, "mode", "dev", "`dev` for development mode. `prod` for production mode")
	flag.Parse()

	// normalize flags
	if mode != "dev" && mode != "prod" {
		mode = "dev"
	}

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
			Mode: mode,
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

func ProvideConfig() fx.Option {
	Init()
	return fx.Provide(GetConfig)
}
