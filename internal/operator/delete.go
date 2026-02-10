package operator

import (
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/yamabuki/nacos-config-tool/config"
)

func DelConfig(dataId string, group string) (bool, error) {
	client := config.GlobalNacosClient
	result, err := client.DeleteConfig(
		vo.ConfigParam{
			DataId: dataId,
			Group:  group,
		},
	)

	if err != nil {
		return false, err
	}

	return result, nil
}
