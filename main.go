package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/rezairfanwijaya/go-autoscale.git/handler"
)

func main() {
	env, err := getEnv()
	if err != nil {
		return
	}
	port := env["PORT"]

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Root)
	mux.HandleFunc("/users", handler.GetUserList)

	log.Printf("will start the app on port %s\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux); err != nil {
		log.Fatalf("failed to start the app on port: %s, err: %s\n", port, err)
	}
}

func getEnv() (map[string]string, error) {
	env, err := godotenv.Read("application.yml")
	if err != nil {
		log.Printf("failed to load env file, err: %s\n", err)
		return env, err
	}

	return env, nil
}
