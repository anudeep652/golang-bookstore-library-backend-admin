package main

import (
	"fmt"
	"log"
	"os"

	"github.com/anudeep652/golang-bookstore-library-backend/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	app := book.Router()
	fmt.Println(`app is listening on port $PORT`)

	app.Listen(":" + os.Getenv("PORT"))
}
