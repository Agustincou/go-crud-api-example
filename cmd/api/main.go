package main

import (
	"log"
	"os"

	"github.com/Agustincou/go-crud-api-example/internal/clients"
	"github.com/Agustincou/go-crud-api-example/internal/handlers"
	"github.com/Agustincou/go-crud-api-example/internal/routes"
	"github.com/Agustincou/go-crud-api-example/pkg/enums"
	"go.uber.org/zap"
)

func main() {
	handler, handlerErr := makeHandler()
	if handlerErr != nil {
		log.Fatal(handlerErr)
	}

	app := routes.MakeApp(handler)

	log.Fatal(app.Listen(":8080"))
}

func makeHandler() (handlers.Handler, error) {
	var kvsClient clients.KVS
	var telemetryClient clients.Telemetry
	var logger *zap.Logger

	if os.Getenv(string(enums.GoEnvironmentKey)) == string(enums.RemoteEnv) {
		// ... Loading production clients
	} else {
		kvsClient = clients.NewInMemoryKVS()
		telemetryClient = clients.NewLogTelemetry()
		logger, _ = zap.NewDevelopment()
	}

	timeHandler := clients.NewSystemTime()
	idGen := clients.NewUuidGenerator()

	logger.Sync()

	return handlers.NewProductsHandler().
		WithIDGen(idGen).
		WithKvsClient(kvsClient).
		WithLogger(logger).
		WithTelemetry(telemetryClient).
		WithTimeHandler(timeHandler).
		Build()
}
