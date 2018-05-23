package main

import (
	"github.com/kmilinho/twcli/cmd/twcli/cli"
)

//type TwitterCred struct {
//	ConsumerKey    string
//	ConsumerSecret string
//	OauthToken     string
//	OauthSecret    string
//}

func main() {
	//var credentials TwitterCred
	//
	//viper.SetConfigName("credentials")
	//viper.AddConfigPath(".")
	//
	//if err := viper.ReadInConfig(); err != nil {
	//	log.Fatalf("Error reading credentials file, %s", err)
	//}
	//err := viper.Unmarshal(&credentials)
	//if err != nil {
	//	log.Fatalf("unable to decode into struct, %v", err)
	//}

	cli.Run()
}
