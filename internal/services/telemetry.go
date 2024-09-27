package services

import (
	"github.com/Agustincou/go-crud-api-example/internal/clients"
	"github.com/Agustincou/go-crud-api-example/internal/enums"
)

type Telemetry interface {
	Send(name enums.MetricName, value float64, tags map[enums.MetricTagKey]string) //nolint:unused
}

type telemetryImpl struct {
	client clients.Telemetry
}

func NewTelemetry(client clients.Telemetry) Telemetry {
	return &telemetryImpl{
		client: client,
	}
}

func (t *telemetryImpl) Send(name enums.MetricName, value float64, tags map[enums.MetricTagKey]string) {
	t.client.Post(name, value, tags)
}
