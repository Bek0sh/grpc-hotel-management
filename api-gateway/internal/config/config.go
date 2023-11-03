package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug"`
	Run     struct {
		Port string `yaml:"port"`
	} `yaml:"run"`
	AuthServicePort    string `yaml:"auth_service_port"`
	BookingServicePort string `yaml:"booking_service_port"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(
		func() {
			instance = &Config{}
			if err := cleanenv.ReadConfig("C:/Users/bekar/VScodeProjects/hotel-management/api-gateway/config.yaml", instance); err != nil {
				panic(err)
			}
		},
	)
	return instance
}
