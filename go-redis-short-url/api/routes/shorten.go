package routes

import (
	"go-redis-short-url/api/database"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Request struct {
	URL         string        `json:"url"`
	Customshort string        `json:"short"`
	Expiry      time.Duration `json:"expires"`
}

type Response struct {
	URL                 string        `json:"url"`
	Customshort         string        `json:"short"`
	Expiry              time.Duration `json:"expires"`
	XRateLimitReset     int           `json:"x_rate_limit_reset"`
	XRateLimitRemaining time.Duration `json:"x_rate_limit_rem"`
}

func ShortenUrl(c *fiber.Ctx) error {

	body := new(Request)
	// get the json and mapped to golang struct

	if body.URL == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "url is required",
		})
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	//implememting rate limiter
	r2 := database.CreateClient(1)
	defer r2.Close()
	_, err := r2.Get(database.Ctx, c.IP()).Result()
	if err != nil {

		_ = r2.Set(database.Ctx, c.IP(), os.Getenv("API_QUOTA"), 30*60*time.Second).Err()
	} else {
		val, _ := r2.Get(database.Ctx, c.IP()).Result()
		valInt, _ := strconv.Atoi(val)
		if valInt <= 0 {

			limit, _ := r2.TTL(database.Ctx, c.IP()).Result()
			return c.Status(429).JSON(fiber.Map{
				"error": "rate limit exceeded",
				"limit": limit,
			})

		}
	}

	//validate for the actual URL

	/* if govalidator.IsUrl(c.Request.URL) {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid url",
		})
	} */

	//check if user has provided a custom short url  if yes then proceed or else
	//create a new using first 6 digits of uuid

	/* var id string
	if body.Customshort != "" {
		id = body.Customshort
	} else {
		id = uuid.New().String()[:6]
	} */

	return nil

}
