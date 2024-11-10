package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rezairfanwijaya/go-autoscale.git/handler"
	"github.com/spf13/viper"
)

func main() {
	env := getEnv()
	port := env["PORT"]
	appName := env["APP_NAME"]

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Root)
	mux.HandleFunc("/users", handler.GetUserList)

	log.Printf("app %s will start the app on port %s\n", appName, port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux); err != nil {
		log.Fatalf("failed to start the app on port: %s, err: %s\n", port, err)
	}
}

func getEnv() map[string]string {
	v := viper.New()
	v.SetConfigName("application")
	v.SetConfigType("yaml")

	v.AddConfigPath(".")

	err := v.ReadInConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("fatal error config file: %w", err))
	}

	return map[string]string{
		"PORT":     v.GetString("PORT"),
		"APP_NAME": v.GetString("APP_NAME"),
	}
}
