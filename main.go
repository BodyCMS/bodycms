package main

import (
	"os"

	"github.com/BodyCMS/bodycms/core/category"
	"github.com/BodyCMS/bodycms/core/post"
	"github.com/BodyCMS/bodycms/core/tag"
	"github.com/BodyCMS/bodycms/core/user"
	_ "github.com/BodyCMS/bodycms/docs"
	"github.com/BodyCMS/bodycms/lib/controller"
	"github.com/BodyCMS/bodycms/lib/db"
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
	pdb := db.GetDbInstance()
	err := pdb.Connect()
	if err != nil {
		panic(err)
	}

	// Migration
	db.MigrateAll(&tag.TagRepo{
		CmsDb: pdb,
	}, &category.CategoryRepo{
		CmsDb: pdb,
	}, &user.UserRepo{
		CmsDb: pdb,
	}, &post.PostRepo{
		CmsDb: pdb,
	})

	app := fiber.New()

	// Controller
	v1Route := app.Group("/api/v1")
	controller.ApplyController(v1Route, &tag.TagController{})
	controller.ApplyController(v1Route, &category.CategoryController{})
	controller.ApplyController(v1Route, &user.UserController{})
	controller.ApplyController(v1Route, &post.PostController{})

	// Swagger
	controller.ApplySwaggerRoutes(app)

	app.Listen(":" + os.Getenv("PORT"))
}
