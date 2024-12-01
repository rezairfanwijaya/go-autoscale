package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rezairfanwijaya/go-autoscale.git/handler"
	"github.com/spf13/viper"
)

func StartApp() {
	env := getEnv()
	port := env["APP_PORT"]
	appName := env["APP_NAME"]

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Root)
	mux.HandleFunc("/users", handler.GetUserList)
	mux.HandleFunc("/users/panic", handler.GetUserWithPanic)

	log.Printf("app %s will start the app on port %s\n", appName, port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux); err != nil {
		log.Fatalf("failed to start the app on port: %s, err: %s\n", port, err)
	}
}

func StartWorker() {
	env := getEnv()
	port := env["WORKER_PORT"]
	appName := env["APP_NAME"]

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Root)
	mux.HandleFunc("/workers", handler.GetWorker)

	log.Printf("app %s will start the app on port %s\n", appName, port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux); err != nil {
		log.Fatalf("failed to start the app on port: %s, err: %s\n", port, err)
	}
}

func getEnv() map[string]string {
	v := viper.New()
	v.AutomaticEnv()
	v.SetConfigName("application")
	v.SetConfigType("yaml")

	v.AddConfigPath("../../../")
	v.AddConfigPath("../../")
	v.AddConfigPath("../")
	v.AddConfigPath(".")

	err := v.ReadInConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("fatal error config file: %w", err))
	}

	return map[string]string{
		"APP_PORT":    v.GetString("APP_PORT"),
		"WORKER_PORT": v.GetString("WORKER_PORT"),
		"APP_NAME":    v.GetString("APP_NAME"),
	}
}
