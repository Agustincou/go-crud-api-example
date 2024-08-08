package enums

type MetricName string

const (
	MyMetricName MetricName = "custom.metric.name"
)

type MetricTagKey string

const (
	ApiActionKey MetricTagKey = "action"
)
