# Installation Guide

## Requirement

This SDK requires Go 1.12 and higher go module feature.

## Installation

Using Qingstor SDK is easy. First, use `go get` to install the latest version of the library from GitHub:
Please always replace the version in the command with the latest [release](https://github.com/qingstor/qingstor-sdk-go/releases).

``` bash
$ GO111MODULE=on go get -u github.com/qingstor/qingstor-sdk-go/v4@v3.0.2
```

Next, include Qingstor sdk in your application:

```
import "github.com/qingstor/qingstor-sdk-go/v4/config"
import "github.com/qingstor/qingstor-sdk-go/v4/service"
```


