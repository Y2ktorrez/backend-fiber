package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/torrez/src"
)

func SetupApi(app *fiber.App, c *src.Container) {
	api := app.Group("/api")

	handlers := []func(fiber.Router){
		//UserHandler
		c.UserHandler.RegisterRoutes,
	}

	for _, register := range handlers {
		register(api)
	}

}
