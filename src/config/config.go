package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper
var Appconfig *viper.Viper

func Init() {
	var err error
	config = viper.New()
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath("/Go/Gin/src/config/")
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on configuration file", err.Error())
	}
}

func GetConfig() *viper.Viper {
	return config
}
