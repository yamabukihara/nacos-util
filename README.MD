nacos-config-tool/
├── cmd/
│   ├── backup.go
│   ├── get.go
│   ├── root.go
│   └── set.go
├── config/
│   └── nacos.go
├── internal/
│   ├── operator/
│   │   ├── backup.go
│   │   └── get.go
│   └── utils/
│       ├── backup.go
│       └── format.go
├── config.yaml
├── go.mod
├── go.sum
└── main.go

一款适用于生产运维的 Nacos 配置管理工具

Usage:
  nacos-tool [command]

Available Commands:
  backup      备份 Nacos 配置（后续实现）
  completion  Generate the autocompletion script for the specified shell
  get         查询 Nacos 中的单个配置
  help        Help about any command
  set         新增/更新 Nacos 配置（后续实现）

Flags:
  -h, --help              help for nacos-tool
  -n, --namespace string  Nacos 命名空间 ID（覆盖配置文件）
  -p, --password string   Nacos 登录密码（覆盖配置文件）
  -s, --server strings    Nacos 服务地址（覆盖配置文件）
  -u, --username string   Nacos 登录用户名（覆盖配置文件）

Use "nacos-tool [command] --help" for more information about a command.



依赖
go get -u github.com/nacos-group/nacos-sdk-go@v1.1.5
go get github.com/spf13/cobra
go get github.com/spf13/viper
go get gopkg.in/yaml.v3


windows调试
$env:configNacosAddress = "127.0.0.1:8848"
$env:configUsername = "nacos"
$env:configPassword = "nacos"
$env:configNamespace = "dev"