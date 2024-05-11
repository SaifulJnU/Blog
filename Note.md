# Power of viper for .env reading
```package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL string
	APIKey      string
	Port        int
}

func configSet() Config {
	viper.SetDefault("DatabaseURL", "localhost:5432")
	viper.SetDefault("APIKey", "default_api_key")
	viper.SetDefault("Port", 8080)

	viper.AutomaticEnv()

	// Optionally, you can set up Viper to read from a config file
	// viper.SetConfigName("config") // Name of config file (without extension)
	// viper.SetConfigType("yaml")   // Required if the config file does not have the extension in the name
	// viper.AddConfigPath("config") // Path to look for the config file in

	// Read the config file
	// if err := viper.ReadInConfig(); err != nil {
	// 	fmt.Printf("Error reading the configs: %v", err)
	// }

	// Unmarshal the configuration into a struct
	var configs Config
	if err := viper.Unmarshal(&configs); err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	return configs
}

func main() {
	config := configSet()
	fmt.Printf("DatabaseURL: %s\n", config.DatabaseURL)
	fmt.Printf("APIKey: %s\n", config.APIKey)
	fmt.Printf("Port: %d\n", config.Port)
}
```