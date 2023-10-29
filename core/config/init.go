package config

import (
	"github.com/spf13/viper"
	"log"
	"sync"
)

var configuration *Configuration
var once sync.Once

func All(path string) *Configuration {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(path)

		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file, %s", err)
		}

		if err := viper.Unmarshal(&configuration); err != nil {
			log.Fatalf("Unable to decode into struct, %v", err)
		}
	})
	return configuration
}
