# 安装指引

## 环境要求

This SDK requires Go 1.12 and higher go module feature.

## 安装

Qingstor SDK 使用起来非常简单。 首先我们使用 `go get` 从 github 上下载最新的版本：
请随时将命令中的版本替换为最新的 [release](https://github.com/yunify/qingstor-sdk-go/releases)。

``` bash
$ GO111MODULE=on go get -u github.com/yunify/qingstor-sdk-go/v3@v3.0.2
```

下一步，您只需要将 Qingstor sdk import 到您的项目中即可:

```
import "github.com/yunify/qingstor-sdk-go/v3/config"
import "github.com/yunify/qingstor-sdk-go/v3/service"
```


