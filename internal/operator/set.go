package operator

import (
	"errors"
	"log"
	"os"

	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/yamabuki/nacos-config-tool/config"
)

func SetConfig(dataId, group, filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("config file read failed: %v", err)
	}
	client := config.GlobalNacosClient
	result, err := client.PublishConfig(vo.ConfigParam{
		DataId:  dataId,
		Group:   group,
		Content: string(content),
	})
	if err != nil {
		return err
	}
	if !result {
		return errors.New("config upload failed, detail info in nacos sdk log")
	}
	return nil
}
