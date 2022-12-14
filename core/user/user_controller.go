package user

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type UserController struct{}

func (u *UserController) ApplyController(router fiber.Router) error {
	router.Get("/users", handleUserGetAll)
	fmt.Println("UserController: /users - GET added")

	router.Get("/users/:id", handleUserGet)
	fmt.Println("UserController: /users/:id - GET added")

	router.Post("/users", handleUserPost)
	fmt.Println("UserController: /users - POST added")

	router.Put("/users/:id", handleUserPut)
	fmt.Println("UserController: /users/:id - PUT added")

	router.Delete("/users/:id", handleUserDelete)
	fmt.Println("UserController: /users/:id - DELETE added")
	return nil
}

func handleUserGetAll(c *fiber.Ctx) error {
	return c.SendString("GET /users")
}

func handleUserGet(c *fiber.Ctx) error {
	return c.SendString("GET /users/:id")
}

func handleUserPost(c *fiber.Ctx) error {
	return c.SendString("POST /users")
}

func handleUserPut(c *fiber.Ctx) error {
	return c.SendString("PUT /users/:id")
}

func handleUserDelete(c *fiber.Ctx) error {
	return c.SendString("DELETE /users/:id")
}
