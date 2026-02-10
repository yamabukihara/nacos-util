package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/yamabuki/nacos-config-tool/internal/operator"
)

var (
	getDataId  string
	getGroup   string
	backupFile string
)

var getCmd = cobra.Command{
	Use:   "get",
	Short: "download config from nacos",
	Run: func(cmd *cobra.Command, args []string) {

		//调用operator下载配置
		content, err := operator.GetConfig(getDataId, getGroup)
		if err != nil {
			log.Panicf("download config failed: %v", err)
		}
		if content == "" {
			log.Fatalf("config(group=%v, dataid=%v) is empty", getGroup, getDataId)
		}
		log.Printf("config(group=%v, dataid=%v) download successfully", getGroup, getDataId)
		// utils.PrettyYAML(content)
		err = writeToFile(content, backupFile)
		if err != nil {
			log.Fatalf("write to file failed: %v", err)
		}
		log.Printf("backup to %v successfully", backupFile)
	},
}

func init() {
	rootCmd.AddCommand(&getCmd)

	getCmd.Flags().StringVarP(&getDataId, "dataid", "d", "", "Nacos DATA ID(Not null)")
	getCmd.Flags().StringVarP(&getGroup, "group", "g", "", "Nacos GROUP(Not null)")
	getCmd.Flags().StringVarP(&backupFile, "file", "f", "", "download file(Not null)")

	getCmd.MarkFlagRequired("dataid")
	getCmd.MarkFlagRequired("group")
	getCmd.MarkFlagRequired("file")
}

func writeToFile(content, filename string) error {
	_, err := os.Stat(filename)
	if err == nil {
		log.Fatalf("file %v already exist", filename)
	}
	if !os.IsNotExist(err) {
		return fmt.Errorf("check %v status abnormal: %v", filename, err)
	}
	filedir := filepath.Dir(filename)

	err = os.MkdirAll(filedir, 0755)
	if err != nil {
		return fmt.Errorf("create directory %v failed: %v", filedir, err)
	}
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("create file %v failed: %v", filename, err)
	}
	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("write conten to file failed: %w", err)
	}

	return nil
}
