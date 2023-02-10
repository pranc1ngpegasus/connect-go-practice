package configuration

import (
	"fmt"

	env "github.com/Netflix/go-env"
	domain "github.com/Pranc1ngPegasus/connect-go-practice/domain/configuration"
	"github.com/google/wire"
)

var _ domain.Configuration = (*Configuration)(nil)

var NewConfigurationSet = wire.NewSet(
	wire.Bind(new(domain.Configuration), new(*Configuration)),
	NewConfiguration,
)

type Configuration struct {
	common *domain.Common
	gcp    *domain.GCP
}

func NewConfiguration() (*Configuration, error) {
	var config domain.Config

	if _, err := env.UnmarshalFromEnviron(&config); err != nil {
		return nil, fmt.Errorf("failed to load environment variables: %w", err)
	}

	return &Configuration{
		common: &config.Common,
		gcp:    &config.GCP,
	}, nil
}

func (c *Configuration) Common() *domain.Common {
	return c.common
}

func (c *Configuration) GCP() *domain.GCP {
	return c.gcp
}
