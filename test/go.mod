module github.com/yunify/qingstor-sdk-go/test

go 1.12

require (
	github.com/DATA-DOG/godog v0.7.13
	github.com/yunify/qingstor-sdk-go/v3 v3.0.0
	gopkg.in/yaml.v2 v2.2.2
)

replace github.com/yunify/qingstor-sdk-go/v3 => ../
