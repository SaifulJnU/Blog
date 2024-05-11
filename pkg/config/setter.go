package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Set() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("config") // path to look for the config file in

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading the configs: %v", err)
	}

	err := viper.Unmarshal(&configurations)

	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
}
