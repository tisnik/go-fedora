package main

import (
	"github.com/spf13/viper"
	"log"
)

func main() {
	log.Println("Reading configuration")

	viper.SetConfigName("config4")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error in config file: %s \n", err)
	}
	log.Println("Done")

	serviceConfig := viper.GetStringMap("service")
	url := serviceConfig["url"]
	port := serviceConfig["port"]

	usersConfig := viper.Sub("users")
	accepted := usersConfig.GetStringSlice("accepted")
	blacklisted := usersConfig.GetStringSlice("blacklisted")
	ids := usersConfig.GetIntSlice("ids")

	log.Printf("Starting the service at address %s:%d\n", url, port)
	log.Printf("Accepted users: %v\n", accepted)
	log.Printf("Blacklisted users: %v\n", blacklisted)
	log.Printf("UIDs: %v\n", ids)
}
