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
		Port   string `yaml:"port" env-default:"4444"`
	} `yaml:"listen"`
	MongoDB struct {
		Host       string `json:"host" env-default:"localhost"`
		Port       string `json:"port" env-default:"27019"`
		Database   string `json:"database" env-default:"user-service"`
		Auth_db    string `json:"auth_Db" env-default:"admin"`
		Username   string `json:"username" env-default:"root"`
		Password   string `json:"password" env-default:"qwerty"`
		Collection string `json:"collection" env-default:"users"`
	} `json:"mongodb"`
	Secret struct {
		Jwtkey string `yaml:"jwtkey" env-default:"jwtkey"`
	} `yaml:"secret"`
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
