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

func handleTagGetAll(c *fiber.Ctx) error {
	var tags []Tag
	err := tagService.FindAll(&TagWhere{}, &tags)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(tags)
}

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
