package main

import (
	"log"

	"./config"
	"github.com/kekik/viper"
)

func main() {

	//We can set the environment variables prefix using SetEnvPrefix
	//viper.SetEnvPrefix("GOFS")

	// Load the config from the environment
	viper.AutomaticEnv()

	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	var configuration config.Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	//configuration can be accessed from the `configuration` instance
	//configuration.BasePath
}
