package cmd

import (
	"fmt"
	"releasr/modules/deploy"
	"releasr/modules/release"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application
// The values are read by viper from a config file or environment variable
type Config struct {
	ReleaseService release.Config `mapstructure:"release,omitempty"`
	DeployService  deploy.Config  `mapstructure:"deploy,omitempty"`
}

// LoadConfig reads configuration from file or environment variables
func LoadConfig(cfgFile string) (config Config, err error) {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigFile("./config.yaml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	err = viper.Unmarshal(&config)
	return
}
