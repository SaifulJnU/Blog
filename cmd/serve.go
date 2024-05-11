package cmd

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saifuljnu/blog/pkg/config"
	"github.com/saifuljnu/blog/pkg/routing"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

	config.Set()

	routing.Init()

	router := routing.GetRouter()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":      "pong",
			"app name": viper.Get("App.Name"),
		})
	})
	routing.Serve()
}
