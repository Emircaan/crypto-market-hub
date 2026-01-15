package server

import (
	"time"

	"github.com/Emircaan/crypto-service/internal/handler/http"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Server struct {
	app *fiber.App
}

func NewServer(marketHandler *http.MarketHandler) *Server {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	marketHandler.RegisterRoutes(app)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	return &Server{
		app: app,
	}
}

func (s *Server) Listen(addr string) error {
	return s.app.Listen(addr)
}
