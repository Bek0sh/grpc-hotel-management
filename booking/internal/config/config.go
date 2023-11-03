package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug"`

	Run struct {
		Port string `yaml:"port"`
	} `yaml:"run"`

	Mongo struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Port     string `yaml:"port"`
		Host     string `yaml:"host"`
		AuthDb   string `yaml:"authDB"`
		DbName   string `yaml:"db_name"`
	} `yaml:"mongo"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(
		func() {
			instance = &Config{}
			if err := cleanenv.ReadConfig("C:/Users/bekar/VScodeProjects/hotel-management/booking/config.yaml", instance); err != nil {
				panic(err)
			}
		},
	)
	return instance
}
