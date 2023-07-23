package main

import (
	"fmt"
	"log"
	"net/http"

	routes "github.com/pashamakhilkumarreddy/golang-ci-cd/routes"
	"github.com/spf13/viper"
)

func loadEnv(filePath string) func(string) string {
	envFile := filePath
	envFilePath := "/go/bin/"
	if len(envFile) == 0 {
		envFile = "example.env"
		log.Printf(".env file path not provided using the default file %s\n", envFile)
	}
	envFilePath = envFilePath + envFile
	viper.SetDefault("PORT", "8080")
	viper.AddConfigPath(".")
	// viper.SetConfigName("example")
	viper.SetConfigType("env")
	viper.SetConfigFile(envFilePath)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("Config file not found")
		} else {
			if err != nil {
				log.Fatalf("Error reading config file %s", err)
			}
		}
	}

	log.Printf("Using file %s\n", viper.ConfigFileUsed())

	return func(key string) string {
		value, ok := viper.Get(key).(string)
		if !ok {
			log.Fatalf("Invalid type assertion")
		}
		return value
	}
}

func main() {
	getEnv := loadEnv("")
	port := getEnv("PORT")
	fmt.Printf("The server is up and running on localhost:%v\n", port)
	routes.SetUpRoutes()
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
