package config

import (
	"fmt"
	"github.com/nutteen/png-core/core/db"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type Config struct {
	Server 			Server
	Database		db.Config
}

type Server struct {
	Port			string
}

var AppConfig Config

func LoadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("/app/config")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		if _, isFileNotFound := err.(viper.ConfigFileNotFoundError); isFileNotFound {
			log.Fatalln("CONFIG", "config file not found")
		} else {
			log.Fatalln(fmt.Errorf("CONFIG:fatal error config file: %s ", err))
		}
	}

	viper.AutomaticEnv()
	err := viper.Unmarshal(&AppConfig)

	return err
}