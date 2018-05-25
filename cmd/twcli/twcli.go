package main

import (
	"github.com/kmilinho/twcli/pkg/keys"
	"fmt"
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

	exitApp := make(chan bool)

	keyEventManager := keys.New()

	keyEventManager.Register("q", func(key string) {
		keyEventManager.Stop()
		exitApp<-true
	})

	keyEventManager.Register("w", func(key string) {
		fmt.Println("some tweets")
	})


	keyEventManager.Register("s", func(key string) {
		fmt.Println("some other tweets")
	})

	keyEventManager.Start()

	<-exitApp
}
