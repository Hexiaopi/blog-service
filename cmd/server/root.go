package main

import (
	"log"

	"github.com/spf13/cobra"
	_ "go.uber.org/automaxprocs"

	"github.com/hexiaopi/blog-service/internal/config"
	"github.com/hexiaopi/blog-service/internal/server"
)

var rootCmd = &cobra.Command{
	Use:   "blog",
	Short: "blog-service is a blog service by go and vue",
	Run: func(cmd *cobra.Command, args []string) {
		server.Run()
	},
}

func init() {
	config.Init(rootCmd)
}

// @title Blog Service API
// @version 1.0
// @description This is a blog server restful api docs.
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8080
// @BasePath  /
// @securityDefinitions.apikey JWT
// @in header
// @name Authorization
func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("server run err:%v", err)
	}
}
