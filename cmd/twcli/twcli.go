package main

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	ConsumerKey    string
	ConsumerSecret string
	OauthToken     string
	OauthSecret    string
}

func main() {
	var config Config

	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	log.Printf("Your twitter consumer key is %s", config.ConsumerKey)
	log.Printf("Your twitter consumer secret is %s", config.ConsumerSecret)
	log.Printf("Your twitter oauth token is %s", config.OauthToken)
	log.Printf("Your twitter oauth secret is %s", config.OauthSecret)
}
