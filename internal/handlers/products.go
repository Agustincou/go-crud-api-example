package handlers

import (
	"fmt"

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

	kvsService       services.KVS        //nolint:unused
	telemetryService services.Telemetry  //nolint:unused
	timeClient       clients.Time        //nolint:unused
	idGenClient      clients.IDGenerator //nolint:unused
	logger           *zap.Logger         //nolint:unused
}

func NewProductsHandler() Handler {
	return &productsHandler{}
}

// WithKvsClient establece el cliente KVS.
func (h *productsHandler) WithKvsClient(kvs clients.KVS) Handler {
	h.kvsService = services.NewKVS(kvs)

	return h
}

// WithTelemetry establece el cliente de Telemetría.
func (h *productsHandler) WithTelemetry(telemetry clients.Telemetry) Handler {
	h.telemetryService = services.NewTelemetry(telemetry)

	return h
}

// WithLogger establece el logger.
func (h *productsHandler) WithLogger(logger *zap.Logger) Handler {
	h.logger = logger

	return h
}

// WithTimeHandler establece el cliente de Time.
func (h *productsHandler) WithTimeHandler(time clients.Time) Handler {
	h.timeClient = time

	return h
}

// WithIDGen establece el cliente de generador de IDs.
func (h *productsHandler) WithIDGen(idGen clients.IDGenerator) Handler {
	h.idGenClient = idGen

	return h
}

// Build construye el handler y valida que todos los clientes estén configurados.
func (h *productsHandler) Build() (Handler, error) {
	// Aquí podrías validar que todos los clientes estén configurados correctamente.
	// Por ejemplo:
	if h.kvsService == nil || h.telemetryService == nil || h.timeClient == nil || h.idGenClient == nil || h.logger == nil {
		return nil, fmt.Errorf("todos los clientes deben estar configurados")
	}

	return h, nil
}

// CreateProduct maneja la creación de un producto.
func (h *productsHandler) CreateProduct() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// Implementación para crear un producto.
		return c.Status(fiber.StatusNotImplemented).SendString("CreateProduct no implementado")
	}
}

// GetProduct maneja la obtención de un producto por ID.
func (h *productsHandler) GetProduct() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// Implementación para obtener un producto.
		return c.Status(fiber.StatusNotImplemented).SendString("GetProduct no implementado")
	}
}

// SearchProduct maneja la búsqueda de productos.
func (h *productsHandler) SearchProduct() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// Implementación para buscar productos.
		return c.Status(fiber.StatusNotImplemented).SendString("SearchProduct no implementado")
	}
}

// DeleteProduct maneja la eliminación de un producto.
func (h *productsHandler) DeleteProduct() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// Implementación para eliminar un producto.
		return c.Status(fiber.StatusNotImplemented).SendString("DeleteProduct no implementado")
	}
}

// UpdateProduct maneja la actualización de un producto.
func (h *productsHandler) UpdateProduct() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// Implementación para actualizar un producto.
		return c.Status(fiber.StatusNotImplemented).SendString("UpdateProduct no implementado")
	}
}
