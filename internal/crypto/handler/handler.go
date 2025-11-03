package crypto

import (
	cryptomodels "konverter/internal/crypto/models"
	"konverter/internal/crypto/usecase"
	"konverter/internal/models"

	"github.com/gofiber/fiber/v2"
)

func Encrypt(c *fiber.Ctx) error {
	req := cryptomodels.EncryptRequest{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusOK).JSON(models.Response{Success: false, Error: err.Error()})
	}

	res, err := usecase.Encrypt(req)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(models.Response{Success: false, Error: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{Success: true, Data: res})
}

func Decrypt(c *fiber.Ctx) error {
	req := cryptomodels.DecryptRequest{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusOK).JSON(models.Response{Success: false, Error: err.Error()})
	}

	res, err := usecase.Decrypt(req)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(models.Response{Success: false, Error: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{Success: true, Data: res})
}

