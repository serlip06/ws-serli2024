package main

import (
	"log"

	"github.com/serlip06/ws-serli2024/config"
	

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"


	"github.com/serlip06/ws-serli2024/url"

	"github.com/gofiber/fiber/v2"
)

func main() {
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
