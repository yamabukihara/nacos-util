package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/yamabuki/nacos-config-tool/config"
)

var rootCmd = &cobra.Command{
	Use:   "nacos-util",
	Short: "Nacos config tool",
	Long:  "support to download/upload config to nacos",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		_, err := config.InitNacosClient()
		if err != nil {
			log.Fatalf("nacos client initialize failed: %v", err)
		}
		return err
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {

}
