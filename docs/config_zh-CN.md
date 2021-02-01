# 配置指南

## 总览

SDK 封装一个 `Config` struct 用于存储和管理配置信息，您可以通过 ["config/config.go"](https://github.com/qingstor/qingstor-sdk-go/blob/master/config/config.go) 来获取它的详细信息（如内部成员和导出方法等）。

除 `ACCESS_KEY_ID` 和 `SECRET_ACCESS_KEY` 之外，您还可以配置 `API服务器` 以进行私有云使用场景。 所有可用的可配置项目都列在默认配置文件中。

___默认的 Config 对应的 yaml 文件:___

``` yaml
# QingStor services configuration

access_key_id: 'ACCESS_KEY_ID'
secret_access_key: 'SECRET_ACCESS_KEY'

host: 'qingstor.com'
port: 443
protocol: 'https'
connection_retries: 3

endpoint: 'https://qingstor.com:443'

enable_virtual_host_style: false # default false.
enable_dual_stack: false # default false.

# Valid log levels are "debug", "info", "warn", "error", and "fatal".
log_level: 'warn'
```

我们也支持设置如下环境变量：

- QINGSTOR_ACCESS_KEY_ID
- QINGSTOR_SECRET_ACCESS_KEY
- QINGSTOR_CONFIG_PATH
- QINGSTOR_ENABLE_VIRTUAL_HOST_STYLE
- QINGSTOR_ENABLE_DUAL_STACK


## 使用

只需使用 API Access Key Pair 来创建 `Config` 结构实例，并使用 service 包下的Init() 函数来初始化 Config 对应的服务。

## 代码片段

创建默认的 Config 结构。

```go
defaultConfig, _ := config.NewDefault()
```

通过密钥来创建 Config。

```go
configuration, _ := config.New("ACCESS_KEY_ID", "SECRET_ACCESS_KEY")

anotherConfiguration := config.NewDefault()
anotherConfiguration.AccessKeyID = "ACCESS_KEY_ID"
anotherConfiguration.SecretAccessKey = "SECRET_ACCESS_KEY"
```

下面的代码从默认路径 `~/.qingstor/config.yaml` 读取配置信息来创建 Config。

```go
userConfig, _ := config.NewDefault().LoadUserConfig()
```

您也可以指定文件路径来初始化 Config。

```go
configFromFile, _ := config.NewDefault().LoadConfigFromFilepath("PATH/TO/FILE")
```

选择更换 API 服务器：

```go
moreConfiguration, _ := config.NewDefault()

moreConfiguration.Protocol = "http"
moreConfiguration.Host = "api.private.com"
moreConfiguration.Port = 80
```

动态修改 http 超时时间：

```go
customConfiguration, _ := config.NewDefault().LoadUserConfig()
// For the default value refers to DefaultHTTPClientSettings in config package
// ReadTimeout affect each call to HTTPResponse.Body.Read()
customConfiguration.HTTPSettings.ReadTimeout = 2 * time.Minute
// WriteTimeout affect each write in io.Copy while sending HTTPRequest
customConfiguration.HTTPSettings.WriteTimeout = 2 * time.Minute
// Re-initialize the client to take effect
customConfiguration.InitHTTPClient()
```
