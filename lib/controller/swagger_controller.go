package controller

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func ApplySwaggerRoutes(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)     // default
	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		DeepLinking:  false,
		DocExpansion: "none",
	}))
}
