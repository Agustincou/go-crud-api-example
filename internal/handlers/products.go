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

	CreateProductHandler() func(*fiber.Ctx) error
	GetProductHandler() func(*fiber.Ctx) error
	SearchProductHandler() func(*fiber.Ctx) error
	DeleteProductHandler() func(*fiber.Ctx) error
	UpdateProductHandler() func(*fiber.Ctx) error
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
