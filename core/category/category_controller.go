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

// Show all categories
// @Summary Show all categories
// @Description Show all categories
// @Tags categories
// @Accept  json
// @Produce  json
// @Success 200 {object} Category
// @Failure 500 {string} string "Internal Server Error"
// @Router /categories [get]
func handleCategoryGetAll(c *fiber.Ctx) error {
	var categories []Category
	err := categoryService.FindAll(&CategoryWhere{}, &categories)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(categories)
}

// Show category by id
// @Summary Show category by id
// @Description Show category by id
// @Tags categories
// @Accept  json
// @Produce  json
// @Param id path int true "Category ID"
// @Success 200 {object} Category
// @Failure 500 {string} string "Internal Server Error"
// @Router /categories/{id} [get]
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

// Create category
// @Summary Create category
// @Description Create category
// @Tags categories
// @Accept  json
// @Produce  json
// @Param category body Category true "Category"
// @Success 200 {object} Category
// @Failure 500 {string} string "Internal Server Error"
// @Router /categories [post]
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

// Update category
// @Summary Update category
// @Description Update category
// @Tags categories
// @Accept  json
// @Produce  json
// @Param id path int true "Category ID"
// @Param category body Category true "Category"
// @Success 200 {object} Category
// @Failure 500 {string} string "Internal Server Error"
// @Router /categories/{id} [put]
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

// Delete category
// @Summary Delete category
// @Description Delete category
// @Tags categories
// @Accept  json
// @Produce  json
// @Param id path int true "Category ID"
// @Success 200 {string} string "OK"
// @Failure 500 {string} string "Internal Server Error"
// @Router /categories/{id} [delete]
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
