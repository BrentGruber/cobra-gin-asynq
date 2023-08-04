/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"releasr/modules/deploy"
	"releasr/modules/release"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func addRoutes(server *gin.Engine, router *gin.Engine) {
	for _, a := range router.Routes() {
		server.Handle(a.Method, a.Path, a.HandlerFunc)
	}
}

// startCmd represents the start command
var (
	cfgFile string

	startCmd = &cobra.Command{
		Use:   "start",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			config, err := LoadConfig(cfgFile)
			if err != nil {
				log.Fatalf("failed to load config file %s", cfgFile)
			}

			server := gin.Default()

			if config.ReleaseService.Enabled {
				releaseService := release.NewReleaseService(config.ReleaseService)
				addRoutes(server, releaseService.Router)
			}

			if config.DeployService.Enabled {
				deployService := deploy.NewDeployService(config.DeployService)
				addRoutes(server, deployService.Router)
			}

			server.Run(":8000")
		},
	}
)

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	startCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Path to config file")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
