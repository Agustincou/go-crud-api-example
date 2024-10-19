package clients

import (
	"go.uber.org/zap"

	"github.com/Agustincou/go-crud-api-example/internal/enums"
)

type Telemetry interface {
	Post(name enums.MetricName, value float64, tags map[enums.MetricTagKey]string) //nolint:unused
}

type logTelemetry struct {
	logger *zap.Logger
}

func NewLogTelemetry() Telemetry {
	logger, _ := zap.NewDevelopment()

	return &logTelemetry{
		logger: logger,
	}
}

func (l *logTelemetry) Post(name enums.MetricName, _ float64, tags map[enums.MetricTagKey]string) { //nolint: unused
	defer l.logger.Sync()

	fields := make([]zap.Field, 0, len(tags))

	for tagKey, tagValue := range tags {
		fields = append(fields, zap.String(string(tagKey), tagValue))
	}

	l.logger.Info("Telemetry "+string(name), fields...)
}
