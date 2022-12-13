package main

import (
	"os"

	"github.com/BodyCMS/bodycms/core/tag"
	_ "github.com/BodyCMS/bodycms/docs"
	"github.com/BodyCMS/bodycms/lib/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
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

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	app := fiber.New()

	// Controller
	v1Route := app.Group("/api/v1")
	controller.ApplyController(v1Route, &tag.TagController{})

	// Swagger
	controller.ApplySwaggerRoutes(app)

	app.Listen(":" + os.Getenv("PORT"))
}
