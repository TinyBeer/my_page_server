package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

var v *viper.Viper
var once sync.Once

func Get() *viper.Viper {
	once.Do(func() {
		v = viper.New()
		v.SetConfigFile("./config/config.yaml")
		err := v.ReadInConfig()
		if err != nil {
			log.Fatalln("read config file failed")
		}
	})
	return v
}
