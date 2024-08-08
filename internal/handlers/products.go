package handlers

import (
	"github.com/Agustincou/go-crud-api-example/internal/clients"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Handler interface {
	WithKvsClient(clients.KVS) Handler
	WithTelemetry(clients.Telemetry) Handler
	WithLogger(*zap.Logger) Handler
	WithTimeHandler(clients.Time) Handler
	WithIDGen(clients.IDGenerator) Handler

	Build() (Handler, error)

	GetCreateProductHandler() func(c *fiber.Ctx) error
	GetGetProductHandler() func(c *fiber.Ctx) error
	GetSearchProductHandler() func(c *fiber.Ctx) error
	GetDeleteProductHandler() func(c *fiber.Ctx) error
	GetUpdateProductHandler() func(c *fiber.Ctx) error
}

type productsHandler struct {
	Handler
}

func NewProductsHandler() Handler {
	return &productsHandler{}
}
