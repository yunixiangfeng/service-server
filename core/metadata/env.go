package metadata

const (
	EnvDev  = "dev"
	EnvProd = "prod"
)

const (
	ServiceName = "ServiceACS"
)

var RequiredEnvs = []string{
	"ATTAINS_APOLLO_APP_ID",
	"ATTAINS_APOLLO_CLUSTER",
	"ATTAINS_APOLLO_URL",
	"ATTAINS_APOLLO_SECRET",
}
