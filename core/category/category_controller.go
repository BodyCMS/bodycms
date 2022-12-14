package category

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var categoryService *CategoryService

func init() {
	categoryService = &CategoryService{
		CategoryRepo: &CategoryRepo{},
	}

	categoryService = categoryService.New()
}

type CategoryController struct{}

func (c *CategoryController) ApplyController(router fiber.Router) error {
	router.Get("/categories", handleCategoryGetAll)
	fmt.Println("CategoryController: /categories - GET added")

	router.Get("/categories/:id", handleCategoryGet)
	fmt.Println("CategoryController: /categories/:id - GET added")

	router.Post("/categories", handleCategoryPost)
	fmt.Println("CategoryController: /categories - POST added")

	router.Put("/categories/:id", handleCategoryPut)
	fmt.Println("CategoryController: /categories/:id - PUT added")

	router.Delete("/categories/:id", handleCategoryDelete)
	fmt.Println("CategoryController: /categories/:id - DELETE added")
	return nil
}

func handleCategoryGetAll(c *fiber.Ctx) error {
	var categories []Category
	err := categoryService.FindAll(&CategoryWhere{}, &categories)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(categories)
}

func handleCategoryGet(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(500)
	}

	var category Category
	err = categoryService.FindOneId(id, &category)

	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(category)
}

func handleCategoryPost(c *fiber.Ctx) error {
	var category Category
	err := c.BodyParser(&category)
	if err != nil {
		return c.SendStatus(500)
	}

	err = categoryService.Create(&category)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(category)
}

func handleCategoryPut(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(500)
	}

	var category Category
	err = c.BodyParser(&category)
	if err != nil {
		return c.SendStatus(500)
	}

	err = categoryService.Update(&CategoryWhere{Id: id}, &CategoryUpdate{Name: category.Name}, &category)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(category)
}

func handleCategoryDelete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(500)
	}
	err = categoryService.Delete(&CategoryWhere{
		Id: id,
	})
	if err != nil {
		return c.SendStatus(500)
	}

	return c.SendStatus(200)
}
