package timestamp

import (
	"konverter/internal/models"
	timestampmodels "konverter/internal/timestamp/models"
	"konverter/internal/timestamp/usecase"

	"github.com/gofiber/fiber/v2"
)

// ConvertHumanize handles timestamp conversion requests
func ConvertHumanize(c *fiber.Ctx) error {
	req := timestampmodels.ConvertRequest{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusOK).JSON(models.Response{Success: false, Error: err.Error()})
	}

	res, err := usecase.ConvertHumanize(req)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(models.Response{Success: false, Error: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{Success: true, Data: res})
}
