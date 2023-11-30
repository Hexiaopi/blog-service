package main

import (
	"github.com/hexiaopi/blog-service/internal/config"
	dao "github.com/hexiaopi/blog-service/internal/store/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "migrate",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		storeIns := dao.NewDao(config.DBEngine)
		if err := storeIns.Migration(); err != nil {
			log.Fatalf("migration err:%v", err)
		}
	},
}

func init() {
	config.Init(rootCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("server run err:%v", err)
	}
}
