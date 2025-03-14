package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/torrez/cmd/api"
	"github.com/torrez/config"
	"github.com/torrez/src"
)

func main() {

	config.Load()

	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, World!")
	})

	c := src.SetupContainer()
	api.SetupApi(app, c)

	log.Println("Server started on port " + config.Port)
	app.Listen(":" + config.Port)

}
