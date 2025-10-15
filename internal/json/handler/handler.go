package json

import (
	jsonmodels "konverter/internal/json/models"
	"konverter/internal/json/usecase"
	"konverter/internal/models"

	"github.com/gofiber/fiber/v2"
)

func Escape(c *fiber.Ctx) error {
	req := jsonmodels.EscapeRequest{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusOK).JSON(models.Response{Success: false, Error: err.Error()})
	}

	res, err := usecase.Escape(req)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(models.Response{Success: false, Error: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{Success: true, Data: res})
}

func Unescape(c *fiber.Ctx) error {
	req := jsonmodels.UnescapeRequest{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusOK).JSON(models.Response{Success: false, Error: err.Error()})
	}

	res, err := usecase.Unescape(req)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(models.Response{Success: false, Error: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{Success: true, Data: res})
}

func Format(c *fiber.Ctx) error {
	req := jsonmodels.FormatRequest{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusOK).JSON(models.Response{Success: false, Error: err.Error()})
	}

	res, err := usecase.Format(req)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(models.Response{Success: false, Error: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{Success: true, Data: res})
}
