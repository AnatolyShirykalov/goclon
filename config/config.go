package config

import (
	"github.com/jinzhu/configor"
)

var Config = struct {
	Port uint `default:"5292" env:"PORT"`
}{}

func init() {
	err := configor.Load(&Config, "config/database.yml")
	if err != nil {
		panic(err)
	}
}
