package main

import (
	"go-redis-short-url/api/routes"

	"github.com/gofiber/fiber/v2"
)

func setUpRoutres(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenUrl)

}

func main() {

	app := fiber.New()
	setUpRoutres(app)

	app.Listen(":8080")

}
