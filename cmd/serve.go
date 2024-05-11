package cmd

import (
	"github.com/saifuljnu/blog/pkg/bootstrap"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "App on the server",
	Long:  "Application will be served on host and port in config.yml file",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {

	bootstrap.Serve()
}
