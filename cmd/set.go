package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/yamabuki/nacos-config-tool/internal/operator"
	"github.com/yamabuki/nacos-config-tool/internal/utils"
)

var (
	setGroup    string
	setDataId   string
	setFilePath string
	setSilent   bool
)

var setCmd = cobra.Command{
	Use:   "set",
	Short: "upload nacos config",
	Run: func(cmd *cobra.Command, args []string) {

		content, err := operator.GetConfig(setDataId, setGroup)
		if err != nil {
			log.Fatalf("remote config download failed: %v", err)
		}
		var prompt string
		if content == "" {
			prompt = "remote config is empty, confirm to continue?"
		} else {
			target, err := os.ReadFile(setFilePath)
			if err != nil {
				log.Fatalf("config file read failed: %v", err)
			}
			result, err := utils.Compare(content, string(target))
			prompt = "config modification: \n" + result + "confirm to continue?"
		}
		if setSilent {
			err := operator.SetConfig(setDataId, setGroup, setFilePath)
			if err != nil {
				log.Panicf("upload config failed: %v", err)
			}
		} else if utils.Confirm(prompt) {
			err := operator.SetConfig(setDataId, setGroup, setFilePath)
			if err != nil {
				log.Panicf("upload config failed: %v", err)
			}
			log.Printf("config(group=%v, dataid=%v) upload successfully", setGroup, setDataId)
		}
	},
}

func init() {
	rootCmd.AddCommand(&setCmd)

	setCmd.Flags().StringVarP(&setDataId, "dataid", "d", "", "Nacos DATA ID(Not null)")
	setCmd.Flags().StringVarP(&setGroup, "group", "g", "", "Nacos GROUP(Not null)")
	setCmd.Flags().StringVarP(&setFilePath, "file", "f", "", "config file(Not null)")
	setCmd.Flags().BoolVarP(&setSilent, "silent", "s", false, "silent mode(Optional)")

	setCmd.MarkFlagRequired("dataid")
	setCmd.MarkFlagRequired("group")
	setCmd.MarkFlagRequired("file")
}
