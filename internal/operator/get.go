package operator

import (
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/yamabuki/nacos-config-tool/config"
)

func GetConfig(dataId string, group string) (string, error) {
	client := config.GlobalNacosClient
	content, err := client.GetConfig(
		vo.ConfigParam{
			DataId: dataId,
			Group:  group,
		},
	)

	if err != nil {
		return "", err
	}

	return content, nil
}
