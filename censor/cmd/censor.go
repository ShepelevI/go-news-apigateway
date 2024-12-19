package main

import (
	"censor/pkg/api"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	portString := os.Getenv("CENSOR_PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	log.Print("server has started")
	router := api.StartAPI()
	err := http.ListenAndServe(fmt.Sprintf(":%s", portString), router)
	if err != nil {
		log.Printf("error from router %v\n", err)
	}
}
