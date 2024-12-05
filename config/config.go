package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

var Cfg AppConfig

func InitConfig() {
	config := DefaultConfig()
	ReadEnv()
	overrideConfig(&config)
	Cfg = config
}
func overrideConfig(config *AppConfig) {
	if err := viper.Unmarshal(config); err != nil {
		log.Fatalf("Failed to unmarshal viper config: %v", err)
	}

	if err := envconfig.Process("", config); err != nil {
		log.Fatalf("Failed to process env config: %v", err)
	}
}
