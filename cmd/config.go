package cmd

import (
	"releasr/pkg/deploy"
	"releasr/pkg/release"
)

// Config stores all configuration of the application
// The values are read by viper from a config file or environment variable
type Config struct {
	ReleaseService release.Config `mapstructure:"release,omitempty"`
	DeployService  deploy.Config  `mapstructure:"deploy,omitempty"`
}

// LoadConfig reads configuration from file or environment variables
