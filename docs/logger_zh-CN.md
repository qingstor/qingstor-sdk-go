# 自定义日志组件

从 v4.4.0 版本开始，我们引入了 `zap.Logger` 组件，这是一个被广泛应用于生产环境的[日志组件](https://github.com/uber-go/zap). 
并且我们废弃了配置文件中的 `log_level` 字段，你可以初始化自定义的 logger，并且通过 `context` 来进行传递。

## log 包

SDK 中，日志组件的帮助方法被定义在了 `log` 包中，这个包主要包含了以下两个方法：
`func ContextWithLogger(ctx context.Context, l *zap.Logger) context.Context` 和
`func FromContext(ctx context.Context) *zap.Logger`.

你可以构建自己的 `*zap.Logger` 实例，通过 `zap.New`, `zap.NewDevelopment`, `zap.NewProducton` 等方法，
然后将该实例设置在 `context` 中，通过 `ContextWithLogger` 方法，之后你可以将得到的 `context` 作为参数传至 API 请求方法中，例如 
`PutObjectWithContext`.

## 默认的 logger

通常情况下，我们会通过 `FromContext` 方法从 `context` 中获取 `*zap.Logger` 实例。但如果 `context` 为 `nil`，或者之前没有实例被设置，
我们会返回一个默认的实例，该实例是通过 `zap.NewProduction` 方法初始化的，并被设置为 `LevelWarn` 等级。
点击 [这里](https://github.com/qingstor/qingstor-sdk-go/blob/master/log/context.go#L39) 可以看到具体的函数逻辑。

## 使用自定义的 logger

你也可以根据实际场景需要，自定义一个 `*zap.Logger` 实例，并通过 `ContextWithLogger` 来构建 `context`. 

示例代码如下:

```go
package main

import (
	"context"

	"github.com/qingstor/qingstor-sdk-go/v4/log"
	"go.uber.org/zap"
)

func main() {
	// 忽略构造 bucketService 的过程

	// 没有 context 参数，将会使用默认的 logger
	bucketService.PutObject(objectKey, input) 

	options := []zap.Option{
		// 自定义你的 logger 选项
	}
	logger, _ := zap.NewDevelopment(options...)
	ctx := log.ContextWithLogger(context.Background(), logger)

	// 上边构造的 logger 将会被使用
	bucketService.PutObjectWithContext(ctx, objectKey, input) 
}
```