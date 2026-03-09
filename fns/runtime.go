package fns

import (
	"net/http"
	"os"
)

// fns 包是共享工具包，提供了通用的工具函数

// 第三方支持：如开发平台，通用接口等
type PerformAgent[R any, P any] func(P) (R, error)

type PerformClient struct {
	PerformAgent[*http.Request, []byte]
	HttpClient *http.Client
}

// LoadProperty 加载配置，参数优先级：CLI参数 > 环境变量 > 默认值
func LoadProperty(p *string, name, defaultValue string) {
	value, exists := os.LookupEnv(name)
	if !exists {
		*p = defaultValue
	} else {
		*p = value
	}
}
