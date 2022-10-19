package main

import (
	"fmt"

	"github.com/anudeep652/golang-bookstore-library-backend/router"
)

func main() {

	app := book.Router()
	fmt.Println(`app is listening on port $PORT`)

	app.Listen(`:$PORT`)
}
