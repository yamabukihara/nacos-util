package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/yamabuki/nacos-config-tool/internal/operator"
)

var (
	delDataId string
	delGroup  string
)

var delCmd = cobra.Command{
	Use:   "delete",
	Short: "delete config from nacos",
	Run: func(cmd *cobra.Command, args []string) {

		//调用operator下载配置
		_, err := operator.DelConfig(delDataId, delGroup)
		if err != nil {
			log.Panicf("download config failed: %v", err)
		}

		log.Printf("config(group=%v, dataid=%v) delete successfully", delGroup, delDataId)

	},
}

func init() {
	rootCmd.AddCommand(&delCmd)

	delCmd.Flags().StringVarP(&delDataId, "dataid", "d", "", "Nacos DATA ID(Not null)")
	delCmd.Flags().StringVarP(&delGroup, "group", "g", "", "Nacos GROUP(Not null)")

	delCmd.MarkFlagRequired("dataid")
	delCmd.MarkFlagRequired("group")
}
