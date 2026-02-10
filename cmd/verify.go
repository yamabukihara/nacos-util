package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/yamabuki/nacos-config-tool/internal/operator"
	"github.com/yamabuki/nacos-config-tool/internal/utils"
)

var (
	verifyDataId   string
	verifyGroup    string
	verifyFilePath string
)

var verifyCmd = cobra.Command{
	Use:   "verify",
	Short: "verify remote config is same with local",
	Run: func(cmd *cobra.Command, args []string) {

		//调用operator下载配置
		content, err := operator.GetConfig(verifyDataId, verifyGroup)
		if err != nil {
			log.Panicf("download remote config failed: %v", err)
		}
		if content == "" {
			log.Fatalf("config(group=%v, dataid=%v) not exist", verifyGroup, verifyDataId)
		}
		target, err := os.ReadFile(verifyFilePath)
		if err != nil {
			log.Fatalf("config file read failed: %v", err)
		}
		result, err := utils.Compare(content, string(target))
		if err != nil {
			log.Fatalf("compare config content failed: %v", err)
		}
		if result != "" {
			log.Fatalf("remote is different from local config\n %v", result)
		} else {
			log.Println("remote is same as local config")
		}
	},
}

func init() {
	rootCmd.AddCommand(&verifyCmd)

	verifyCmd.Flags().StringVarP(&verifyDataId, "dataid", "d", "", "Nacos DATA ID(Not null)")
	verifyCmd.Flags().StringVarP(&verifyGroup, "group", "g", "", "Nacos GROUP(Not null)")
	verifyCmd.Flags().StringVarP(&verifyFilePath, "file", "f", "", "config file(Not null)")

	verifyCmd.MarkFlagRequired("dataid")
	verifyCmd.MarkFlagRequired("group")
	verifyCmd.MarkFlagRequired("file")
}
