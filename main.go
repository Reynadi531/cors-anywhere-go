package main

import (
	_ "fmt"
	"log"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
  client := resty.New()
  app := fiber.New() 
  app.Use(logger.New())
  app.Use(cors.New())

  app.Get("/", func(c *fiber.Ctx) error {

    if len(c.Query("url")) == 0 {
      return c.SendString("Welcome to CORS Proxy")
    }

    res, err := client.R().EnableTrace().Get(c.Query("url"))
    if err != nil {
      log.Fatal(err)
    }

    c.Response().Header.Add("X-Powered-By", "CORS Anywhere Proxy")
    return c.Send(res.Body())
  })

  log.Fatal(app.Listen(":3000"))
}
