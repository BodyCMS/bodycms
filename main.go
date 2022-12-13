package main

import (
	_ "github.com/BodyCMS/bodycms/docs"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

// @title BodyCMS API
// @version 0.0.1
// @description API that BodyCMS provides to its users.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email hungtp.play@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /
func main() {
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:          "http://example.com/doc.json",
		DeepLinking:  false,
		DocExpansion: "none",
	}))

	app.Listen(":8080")
}
