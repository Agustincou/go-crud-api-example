package main

import (
	"log"

	"github.com/Agustincou/go-crud-api-example/internal/handlers"
	"github.com/Agustincou/go-crud-api-example/internal/routes"
)

func main() {
	app := routes.MakeApp(handlers.ProductsHandler{})

	log.Fatal(app.Listen(":8080"))
}
