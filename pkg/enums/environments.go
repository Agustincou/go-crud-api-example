package enums

type EnvironmentKey string

const (
	GoEnvironmentKey EnvironmentKey = "GO_ENVIRONMENT"
)

type Environment string

const (
	LocalEnv  Environment = "localhost"
	RemoteEnv Environment = "remote"
)
