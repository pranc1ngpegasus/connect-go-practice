package configuration

type Configuration interface {
	Common() *Common
	GCP() *GCP
}

type (
	Config struct {
		Common Common
		GCP    GCP
	}

	Common struct {
		Debug bool `env:"DEBUG"`
	}

	GCP struct {
		ProjectID string `env:"GCP_PROJECT_ID"`
	}
)
