package book

import (
	"os"

	"github.com/anudeep652/golang-bookstore-library-backend/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Router() *fiber.App {

	router := fiber.New()
	router.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     os.Getenv("ADMIN_FRONTEND_URL"),
		AllowCredentials: true,
		AllowMethods:     "POST",
	}))

	router.Post("/admin/login", bookController.AdminLogin)
	router.Post("/admin/createbook", bookController.RegisterBook)
	return router
}
