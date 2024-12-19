package main

import (
	"fmt"
	"log"
	"net/http"
	"news/pkg/api"
	db "news/pkg/postgres"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	portString := os.Getenv("NEWS_PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}
	dbString := os.Getenv("DB_URL")
	if dbString == "" {
		log.Fatal("DB_URL is not found in the environment")
	}
	log.Print("server has started")
	pgdb, err := db.New(dbString)
	if err != nil {
		log.Printf("error starting the database %v", err)
	}

	router := api.StartAPI(pgdb)
	err = http.ListenAndServe(fmt.Sprintf(":%s", portString), router)
	if err != nil {
		log.Printf("error from router %v\n", err)
	}
}
