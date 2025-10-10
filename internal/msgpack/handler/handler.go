package msgpack

import (
	"konverter/internal/models"
	msgpackModel "konverter/internal/msgpack/models"
	"konverter/internal/msgpack/usecase"

	"github.com/gofiber/fiber/v2"
)

func Encode(c *fiber.Ctx) error {
	req := msgpackModel.EncodeRequest{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusOK).JSON(models.Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	res, err := usecase.Encode(req)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(models.Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Success: true,
		Data:    res,
	})
}

func Decode(c *fiber.Ctx) error {
	req := msgpackModel.DecodeRequest{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusOK).JSON(models.Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	res, err := usecase.Decode(req)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(models.Response{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Success: true,
		Data:    res,
	})

}
