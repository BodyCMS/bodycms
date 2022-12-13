package tag

import "github.com/gofiber/fiber/v2"

type TagController struct{}

func (t *TagController) ApplyController(router fiber.Router) error {
	router.Get("/tags", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	return nil
}
