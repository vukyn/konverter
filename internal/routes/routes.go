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

func SetupFaviconRoute(app *fiber.App) {
	app.Get("/favicon.ico", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "image/x-icon")
		return c.SendFile("assets/favicon.ico", true)
	})
}

func SetupHealthCheckRoute(app *fiber.App) {
	healthCheck := func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message":   "OK",
			"timestamp": time.Now().UTC(),
			"service":   "konverter",
		})
	}
	app.Get("/", healthCheck)
	app.Head("/", healthCheck)
}

func msgpackRoutes(router fiber.Router) {
	rMsgPack := router.Group("/msgpack")
	rMsgPack.Post("/encode", msgpackHandler.Encode)
	rMsgPack.Post("/decode", msgpackHandler.Decode)
}
