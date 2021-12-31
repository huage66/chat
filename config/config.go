package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Setting struct {
	Application Application `json:"application"`
	Redis       Redis       `json:"redis"`
}

type Application struct {
	Module string `json:"module"` // 模块名称
	Host   string `json:"host"`   // host
	Port   int `json:"port"`   // port
}

type Redis struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

var (
	Config          Setting
	RedisConf       Redis
	ApplicationConf Application
)

func Setup() {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile("./config/settings.yml")
	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("Setup: read config fail, err = %v\n", err)
		panic(err)
	}
	if err := v.Unmarshal(&Config); err != nil {
		fmt.Printf("Setup: Unmarshal config fail, err = %v\n", err)
		panic(err)
	}

	RedisConf = Config.Redis

	ApplicationConf = Config.Application

	fmt.Println("settings = ", Config)
}
