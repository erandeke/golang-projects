package routes

import (
	"go-redis-short-url/api/database"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func ResolveURL(c *fiber.Ctx) error {

	//get the url passed in the broser by the user

	url := c.Params("url")

	r := database.CreateClient(0)
	defer r.Close()

	//check in database if for provided url does the original url already exists
	value, err := r.Get(database.Ctx, url).Result()

	if err != redis.Nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "url not found",
		})
		// if not able to connect to database
	} else if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "internal server error",
		})

	}

	// If all goes well then redirect the url to the original url
	// increment the counter
	rInr := database.CreateClient(1)
	defer rInr.Close()
	_ = rInr.Incr(database.Ctx, "counter")
	// redirect to original URL
	return c.Redirect(value, 301)

}
