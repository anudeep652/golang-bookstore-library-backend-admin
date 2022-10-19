package main

import (
	"github.com/anudeep652/golang-bookstore-library-backend/router"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	_ = godotenv.Load()

	app := book.Router()

	app.Listen(":" + os.Getenv("PORT"))
}
