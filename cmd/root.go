package main

import (
	"log"
	"math/rand"
	"time"

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
	cobra.OnInitialize(config.Init)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
