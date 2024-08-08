package routes

import (
	"log"

	"github.com/Agustincou/go-crud-api-example/internal/errors"
	"github.com/Agustincou/go-crud-api-example/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func MakeApp(handler handlers.Handler) *fiber.App {
	app := fiber.New()

	// Definición del middleware para manejo de errores global
	app.Use(func(c *fiber.Ctx) error {
		if err := c.Next(); err != nil {
			log.Printf("Error: %v\n", err)

			if apiError, isAPIErr := err.(*errors.API); isAPIErr {
				return c.Status(apiError.HttpStatus).JSON(fiber.Map{
					"code":    apiError.Code,
					"message": apiError.Error(),
				})
			}

			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return nil
	})

	matrixApi := app.Group("/products")
	{
		matrixApi.Post("/", handler.GetCreateProductHandler())
		matrixApi.Get("/{id}", handler.GetGetProductHandler())
		matrixApi.Delete("/{id}", handler.GetDeleteProductHandler())
		matrixApi.Put("/{id}", handler.GetUpdateProductHandler())
		matrixApi.Get("/search", handler.GetSearchProductHandler())
	}

	return app
}
