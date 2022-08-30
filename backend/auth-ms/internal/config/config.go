package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
	"sync"
)

type Config struct {
	Listen struct {
		Type   string `yaml:"type" env-default:"port"`
		BindIP string `yaml:"bind_ip" env-default:"localhost"`
		Port   string `yaml:"port" env-default:"8080"`
	} `yaml:"listen"`
	MongoDB struct {
		Host       string `json:"host"`
		Port       string `json:"port"`
		Database   string `json:"database"`
		Auth_db    string `json:"auth_Db"`
		Username   string `json:"username"`
		Password   string `json:"password"`
		Collection string `json:"collection"`
	} `json:"mongodb"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		if err := cleanenv.ReadConfig("./backend/auth-ms/config.yml", instance); err == nil {
			_, err := cleanenv.GetDescription(instance, nil)
			if err != nil {
				logrus.Fatalf("error read config file: %s", err)
			}
		}
	})
	return instance
}
