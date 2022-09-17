package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init(){
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of blog-service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("blog-service version v0.1.0")
	},
}
