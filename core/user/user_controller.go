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

// Show all users
// @Summary Show all users
// @Description Show all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} User
// @Failure 500 {string} string "Internal Server Error"
// @Router /users [get]
func handleUserGetAll(c *fiber.Ctx) error {
	return c.SendString("GET /users")
}

// Show user by id
// @Summary Show user by id
// @Description Show user by id
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Failure 500 {string} string "Internal Server Error"
// @Router /users/:id [get]
func handleUserGet(c *fiber.Ctx) error {
	return c.SendString("GET /users/:id")
}

// Create user
// @Summary Create user
// @Description Create user
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} User
// @Failure 500 {string} string "Internal Server Error"
// @Router /users [post]
func handleUserPost(c *fiber.Ctx) error {
	return c.SendString("POST /users")
}

// Update user by id
// @Summary Update user by id
// @Description Update user by id
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Failure 500 {string} string "Internal Server Error"
// @Router /users/:id [put]
func handleUserPut(c *fiber.Ctx) error {
	return c.SendString("PUT /users/:id")
}

// Delete user by id
// @Summary Delete user by id
// @Description Delete user by id
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Failure 500 {string} string "Internal Server Error"
// @Router /users/:id [delete]
func handleUserDelete(c *fiber.Ctx) error {
	return c.SendString("DELETE /users/:id")
}
