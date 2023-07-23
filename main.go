package main

import (
	"fmt"
	"log"
	"net/http"

	routes "github.com/pashamakhilkumarreddy/golang-ci-cd/routes"
	"github.com/spf13/viper"
)

func loadEnv(filePath string) func(string) string {
	envFilePath := filePath
	if len(envFilePath) == 0 {
		envFilePath = ".env.example"
		log.Fatalf(".env file path not provided using the default path %s", envFilePath)
	}
	viper.SetConfigFile(envFilePath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file %s", err)
	}
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
