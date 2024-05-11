package cmd

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saifuljnu/blog/config"
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

	configs := configSet()

	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":      "pong",
			"app name": viper.Get("App.Name"),
		})
	})

	err := r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port))

	if err != nil {
		panic("Server can not connect")
	}
}

func configSet() config.Config {

	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("config") // path to look for the config file in

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading the configs: %v", err)
	}

	var configs config.Config
	err := viper.Unmarshal(&configs)

	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	return configs
}
