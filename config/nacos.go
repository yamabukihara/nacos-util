package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type NacosConfig struct {
	ServerAddrs []string
	Username    string
	Password    string
	Namespace   string
}

const (
	EnvNacosAddr      = "configNacosAddress"
	EnvNacosUsername  = "configUsername"
	EnvNacosPassword  = "configPassword"
	EnvNacosNamespace = "configNamespace"
)

var GlobalNacosClient config_client.IConfigClient

func InitConfig() (*NacosConfig, error) {

	nacosCfg := &NacosConfig{
		ServerAddrs: parseServerAddr(os.Getenv(EnvNacosAddr)),
		Username:    os.Getenv(EnvNacosUsername),
		Password:    os.Getenv(EnvNacosPassword),
		Namespace:   os.Getenv(EnvNacosNamespace),
	}

	if err := nacosCfg.Validate(); err != nil {
		return nil, err
	}

	return nacosCfg, nil
}

func (c *NacosConfig) Validate() error {
	if len(c.ServerAddrs) == 0 {
		return errors.New("ServerAddrs must be not null (env: " + EnvNacosAddr + "）")
	}
	if c.Username == "" {
		return errors.New("Username must be not null (env:  " + EnvNacosUsername + "）")
	}
	if c.Password == "" {
		return errors.New("Password must be not null (env:  " + EnvNacosPassword + "）")
	}
	if c.Namespace == "" {
		return errors.New("Namespace must be not null (env:  " + EnvNacosNamespace + "）")
	}
	return nil
}

func parseServerAddr(serverAddr string) []string {
	if serverAddr == "" {
		return nil
	}
	serverAddrArr := strings.Split(serverAddr, ",")
	return serverAddrArr
}

func InitNacosClient() (config_client.IConfigClient, error) {

	nacosCfg, err := InitConfig()
	if err != nil {
		return nil, fmt.Errorf("nacos client initialize failed: %w", err)
	}

	clientCfg := constant.NewClientConfig(
		constant.WithNamespaceId(nacosCfg.Namespace),
		constant.WithUsername(nacosCfg.Username),
		constant.WithPassword(nacosCfg.Password),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogLevel("Debug"),
		constant.WithLogDir("./log/"),
	)
	var serverCfgs []constant.ServerConfig

	for _, addr := range nacosCfg.ServerAddrs {
		parts := strings.Split(addr, ":")
		port := uint64(8848)
		if len(parts) > 1 {
			fmt.Sscanf(parts[1], "%d", &port)
		}
		serverCfgs = append(serverCfgs, *constant.NewServerConfig(
			parts[0],
			port,
		))
	}

	client, err := clients.NewConfigClient(vo.NacosClientParam{
		ClientConfig:  clientCfg,
		ServerConfigs: serverCfgs,
	})
	if err != nil {
		return nil, fmt.Errorf("create nacos client failed: %w", err)
	}

	GlobalNacosClient = client
	return client, nil
}
