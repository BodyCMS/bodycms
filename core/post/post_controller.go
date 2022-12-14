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

// Show all posts
// @Summary Show all posts
// @Description Show all posts
// @Tags posts
// @Accept  json
// @Produce  json
// @Success 200 {object} Post
// @Failure 500 {string} string "Internal Server Error"
// @Router /posts [get]
func handlePostGetAll(c *fiber.Ctx) error {
	return c.SendString("GET /posts")
}

// Show post by id
// @Summary Show post by id
// @Description Show post by id
// @Tags posts
// @Accept  json
// @Produce  json
// @Param id path int true "Post ID"
// @Success 200 {object} Post
// @Failure 500 {string} string "Internal Server Error"
// @Router /posts/:id [get]
func handlePostGet(c *fiber.Ctx) error {
	return c.SendString("GET /posts/:id")
}

// Create post
// @Summary Create post
// @Description Create post
// @Tags posts
// @Accept  json
// @Produce  json
// @Success 200 {object} Post
// @Failure 500 {string} string "Internal Server Error"
// @Router /posts [post]
func handlePostPost(c *fiber.Ctx) error {
	return c.SendString("POST /posts")
}

// Update post
// @Summary Update post
// @Description Update post
// @Tags posts
// @Accept  json
// @Produce  json
// @Param id path int true "Post ID"
// @Success 200 {object} Post
// @Failure 500 {string} string "Internal Server Error"
// @Router /posts/:id [put]
func handlePostPut(c *fiber.Ctx) error {
	return c.SendString("PUT /posts/:id")
}

// Delete post
// @Summary Delete post
// @Description Delete post
// @Tags posts
// @Accept  json
// @Produce  json
// @Param id path int true "Post ID"
// @Success 200 {object} Post
// @Failure 500 {string} string "Internal Server Error"
// @Router /posts/:id [delete]
func handlePostDelete(c *fiber.Ctx) error {
	return c.SendString("DELETE /posts/:id")
}
