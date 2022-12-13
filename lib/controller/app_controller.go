package controller

import "github.com/gofiber/fiber/v2"

type IController interface {
	ApplyController(router fiber.Router) error
}

func ApplyController(router fiber.Router, controller IController) error {
	return controller.ApplyController(router)
}
