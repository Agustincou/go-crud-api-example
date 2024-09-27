package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/Agustincou/go-crud-api-example/internal/clients"
	"github.com/Agustincou/go-crud-api-example/internal/services"
)

type Handler interface {
	WithKvsClient(clients.KVS) Handler
	WithTelemetry(clients.Telemetry) Handler
	WithLogger(*zap.Logger) Handler
	WithTimeHandler(clients.Time) Handler
	WithIDGen(clients.IDGenerator) Handler

	Build() (Handler, error)

	CreateProduct() func(*fiber.Ctx) error
	GetProduct() func(*fiber.Ctx) error
	SearchProduct() func(*fiber.Ctx) error
	DeleteProduct() func(*fiber.Ctx) error
	UpdateProduct() func(*fiber.Ctx) error
}

type productsHandler struct {
	Handler

	kvsService services.KVS //nolint:unused

	telemetryService services.Telemetry //nolint:unused

	timeClient clients.Time //nolint:unused

	idGenClient clients.IDGenerator //nolint:unused

	logger *zap.Logger //nolint:unused
}

func NewProductsHandler() Handler {
	return &productsHandler{}
}
