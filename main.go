package main

import (
	"log"

	"github.com/anudeep652/golang-bookstore-library-backend/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	app := book.Router()

	app.Listen(`0.0.0.0:$PORT`)
}
