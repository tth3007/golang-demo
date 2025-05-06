package utils

import (
	"demo/structs/common"
	"log"

	"github.com/spf13/viper"
)

func LoadConfigs() common.Config {
	config, err := loadEnvConfigs("resources/")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	} else {
		log.Println("Init env configs: ", config)
	}
	return config
}

func loadEnvConfigs(path string) (config common.Config, err error) {
	const CONFIG_NAME = "configs"
	const CONFIG_TYPE = "env"
	viper.AddConfigPath(path)
	viper.SetConfigName(CONFIG_NAME)
	viper.SetConfigType(CONFIG_TYPE)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
