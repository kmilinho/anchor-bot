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

	keys := keys.NewTermBoxKeyListener()

	keys.Register("q", func(key string) {
		keys.Stop()
	})

	keys.Register("w", func(key string) {
		fmt.Println("some tweets")
	})

	keys.Register("s", func(key string) {
		fmt.Println("some other tweets")
	})

	keys.Start()
	keys.Wait()
}
