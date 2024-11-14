package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func LoadConfig() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("error reading .env file, %s", err))
	}
	log.Println(viper.AllSettings())
}
