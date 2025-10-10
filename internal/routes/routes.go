package routes

import (
	msgpackHandler "konverter/internal/msgpack/handler"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SetupRouteV1(app *fiber.App) {
	apiV1 := app.Group("/api/v1")
	msgpackRoutes(apiV1)
}

func SetupHealthCheckRoute(app *fiber.App) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message":   "OK",
			"timestamp": time.Now().UTC(),
			"service":   "konverter",
		})
	})
}

func msgpackRoutes(router fiber.Router) {
	rMsgPack := router.Group("/msgpack")
	rMsgPack.Post("/encode", msgpackHandler.Encode)
	rMsgPack.Post("/decode", msgpackHandler.Decode)
}
