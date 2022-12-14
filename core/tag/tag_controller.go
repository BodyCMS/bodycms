package tag

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type TagController struct{}

func (t *TagController) ApplyController(router fiber.Router) error {
	router.Get("/tags", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	fmt.Println("TagController: /tags - GET added")

	router.Get("/tags/:id", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	fmt.Println("TagController: /tags/:id - GET added")

	router.Post("/tags", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	fmt.Println("TagController: /tags - POST added")

	router.Put("/tags/:id", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	fmt.Println("TagController: /tags/:id - PUT added")

	router.Delete("/tags/:id", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	fmt.Println("TagController: /tags/:id - DELETE added")
	return nil
}
