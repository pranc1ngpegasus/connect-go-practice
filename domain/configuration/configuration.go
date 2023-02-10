package configuration

type Configuration interface {
	Common() *Common
	GCP() *GCP
	Server() *Server
}

type (
	Config struct {
		Common Common
		GCP    GCP
		Server Server
	}

	Common struct {
		Debug bool `env:"DEBUG"`
	}

	GCP struct {
		ProjectID string `env:"GCP_PROJECT_ID"`
	}

	Server struct {
		Port string `env:"SERVER_PORT"`
	}
)
