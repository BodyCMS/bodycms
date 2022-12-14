package post

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type PostController struct{}

func (p *PostController) ApplyController(router fiber.Router) error {
	router.Get("/posts", handlePostGetAll)
	fmt.Println("PostController: /posts - GET added")

	router.Get("/posts/:id", handlePostGet)
	fmt.Println("PostController: /posts/:id - GET added")

	router.Post("/posts", handlePostPost)
	fmt.Println("PostController: /posts - POST added")

	router.Put("/posts/:id", handlePostPut)
	fmt.Println("PostController: /posts/:id - PUT added")

	router.Delete("/posts/:id", handlePostDelete)
	fmt.Println("PostController: /posts/:id - DELETE added")
	return nil
}

func handlePostGetAll(c *fiber.Ctx) error {
	return c.SendString("GET /posts")
}

func handlePostGet(c *fiber.Ctx) error {
	return c.SendString("GET /posts/:id")
}

func handlePostPost(c *fiber.Ctx) error {
	return c.SendString("POST /posts")
}

func handlePostPut(c *fiber.Ctx) error {
	return c.SendString("PUT /posts/:id")
}

func handlePostDelete(c *fiber.Ctx) error {
	return c.SendString("DELETE /posts/:id")
}
