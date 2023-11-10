package config

import (
	"curd/database"
	"curd/route"
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
	Port string `json:"port"`
}

var Data Configuration

func LoadConfig() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	viper.Unmarshal(&Data)
	database.Start()
	route.Start(Data.Port)
}
