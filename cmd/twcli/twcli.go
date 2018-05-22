package main

import (
	"github.com/spf13/viper"
	"log"
	"github.com/kmilinho/twcli/cmd/twcli/core"
)

type TwitterCred struct {
	ConsumerKey    string
	ConsumerSecret string
	OauthToken     string
	OauthSecret    string
}

func main() {
	var credentials TwitterCred

	viper.SetConfigName("credentials")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading credentials file, %s", err)
	}
	err := viper.Unmarshal(&credentials)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	core.Run()
}
