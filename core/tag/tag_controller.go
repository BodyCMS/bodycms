package tag

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var tagService *TagService

func init() {
	tagService = &TagService{
		TagRepo: &TagRepo{},
	}

	tagService = tagService.New()
}

type TagController struct{}

func (t *TagController) ApplyController(router fiber.Router) error {
	router.Get("/tags", handleTagGetAll)
	fmt.Println("TagController: /tags - GET added")

	router.Get("/tags/:id", handleTagGet)
	fmt.Println("TagController: /tags/:id - GET added")

	router.Post("/tags", handleTagPost)
	fmt.Println("TagController: /tags - POST added")

	router.Put("/tags/:id", handleTagPut)
	fmt.Println("TagController: /tags/:id - PUT added")

	router.Delete("/tags/:id", handleTagDelete)
	fmt.Println("TagController: /tags/:id - DELETE added")
	return nil
}

// Show all tags
// @Summary Show all tags
// @Description Show all tags
// @Tags tags
// @Accept  json
// @Produce  json
// @Success 200 {object} Tag
// @Failure 500 {string} string "Internal Server Error"
// @Router /tags [get]
func handleTagGetAll(c *fiber.Ctx) error {
	var tags []Tag
	err := tagService.FindAll(&TagWhere{}, &tags)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(tags)
}

// Show tag by id
// @Summary Show tag by id
// @Description Show tag by id
// @Tags tags
// @Accept  json
// @Produce  json
// @Param id path int true "Tag ID"
// @Success 200 {object} Tag
// @Failure 500 {string} string "Internal Server Error"
// @Router /tags/:id [get]
func handleTagGet(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(500)
	}

	var tag Tag
	err = tagService.FindOneId(id, &tag)

	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(tag)
}

// Create tag
// @Summary Create tag
// @Description Create tag
// @Tags tags
// @Accept  json
// @Produce  json
// @Param tag body Tag true "Tag"
// @Success 200 {object} Tag
// @Failure 500 {string} string "Internal Server Error"
// @Router /tags [post]
func handleTagPost(c *fiber.Ctx) error {
	var tag Tag
	err := c.BodyParser(&tag)
	if err != nil {
		return c.SendStatus(500)
	}

	err = tagService.Create(&tag)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(tag)
}

// Update tag
// @Summary Update tag
// @Description Update tag
// @Tags tags
// @Accept  json
// @Produce  json
// @Param id path int true "Tag ID"
// @Param tag body Tag true "Tag"
// @Success 200 {object} Tag
// @Failure 500 {string} string "Internal Server Error"
// @Router /tags/:id [put]
func handleTagPut(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(500)
	}

	var tag Tag
	err = c.BodyParser(&tag)
	if err != nil {
		return c.SendStatus(500)
	}

	err = tagService.Update(&TagWhere{
		ID: id,
	}, &TagUpdate{
		Name: tag.Name,
	}, &tag)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(tag)
}

// Delete tag
// @Summary Delete tag
// @Description Delete tag
// @Tags tags
// @Accept  json
// @Produce  json
// @Param id path int true "Tag ID"
// @Success 200 {string} string "OK"
// @Failure 500 {string} string "Internal Server Error"
// @Router /tags/:id [delete]
func handleTagDelete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(500)
	}

	err = tagService.Delete(&TagWhere{
		ID: id,
	})
	if err != nil {
		return c.SendStatus(500)
	}

	return c.SendStatus(200)
}
