package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/Agustincou/go-crud-api-example/internal/errors"
	"github.com/Agustincou/go-crud-api-example/internal/handlers"
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

	matrixAPI := app.Group("/product")
	{
		matrixAPI.Post("/", handler.CreateProduct())
		matrixAPI.Get("/:id", handler.GetProduct())
		matrixAPI.Delete("/:id", handler.DeleteProduct())
		matrixAPI.Put("/:id", handler.UpdateProduct())
		matrixAPI.Get("/search", handler.SearchProduct())
	}

	return app
}
