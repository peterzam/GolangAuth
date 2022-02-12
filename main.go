package main

import (
	"log"
	"os"

	"Golang-User-Auth/routes"

	"github.com/joho/godotenv"
)

func main() {
	e := godotenv.Load()
	r := routes.Handlers()

	if e != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	r.Run(":" + port)

}
