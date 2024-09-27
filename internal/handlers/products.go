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

	GetCreateProductHandler() func(*fiber.Ctx) error
	GetGetProductHandler() func(*fiber.Ctx) error
	GetSearchProductHandler() func(*fiber.Ctx) error
	GetDeleteProductHandler() func(*fiber.Ctx) error
	GetUpdateProductHandler() func(*fiber.Ctx) error
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
