package http

import (
	"github.com/Emircaan/crypto-service/internal/service"
	"github.com/gofiber/fiber/v2"
)

type MarketHandler struct {
	service *service.MarketService
}

func NewMarketHandler(service *service.MarketService) *MarketHandler {
	return &MarketHandler{
		service: service,
	}
}

func (h *MarketHandler) RegisterRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Get("/tickers/:exchange", h.GetTickers)
	api.Get("/exchanges", h.GetSupportedExchanges)
}

func (h *MarketHandler) GetTickers(c *fiber.Ctx) error {
	exchange := c.Params("exchange")

	tickers, err := h.service.GetTickers(c.Context(), exchange)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    tickers,
	})
}

func (h *MarketHandler) GetSupportedExchanges(c *fiber.Ctx) error {
	exchanges, err := h.service.GetSupportedExchanges(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    exchanges,
	})
}
