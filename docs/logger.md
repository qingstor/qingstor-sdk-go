# Custom Logger

From v4.4.0, we introduced `zap.Logger` which is a widely used, production
ready [logger component](https://github.com/uber-go/zap). The `log_level` field in config is deprecated, you can
initialize your own logger and pass it by `context`.

## Package log

The logger helper methods are defined in `log` package, which contains
`func ContextWithLogger(ctx context.Context, l *zap.Logger) context.Context` and
`func FromContext(ctx context.Context) *zap.Logger` two methods.

You can conduct your own `*zap.Logger` by methods such as `zap.New`, `zap.NewDevelopment`, `zap.NewProducton` and so on,
and set it into `context` by `ContextWithLogger`, then you can pass this `context` into API function
like `PutObjectWithContext`.

## Default logger

Usually, we get logger from `context` by method `FromContext`. If `context` is `nil` or no `logger` is set before, 
a default logger is returned, which is initialized by `zap.NewProduction` and `LevelWarn` is set.
[Here](https://github.com/qingstor/qingstor-sdk-go/blob/master/log/context.go#L39) is the detail.

## Use custom logger

You can also customize you own logger depend on your scenario. Then conduct the `context` by `ContextWithLogger`.

Here are some examples:

```go
package main

import (
	"context"

	"github.com/qingstor/qingstor-sdk-go/v4/log"
	"go.uber.org/zap"
)

func main() {
	// ignore the process of conducting bucketService

	// no context passed, will use default logger
	bucketService.PutObject(objectKey, input)

	options := []zap.Option{
		// customize your own options
	}
	logger, _ := zap.NewDevelopment(options...)
	ctx := log.ContextWithLogger(context.Background(), logger)

	// logger set above will be used
	bucketService.PutObjectWithContext(ctx, objectKey, input)
}
```