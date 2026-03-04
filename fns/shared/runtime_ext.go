package shared

import (
	"os"
)

// shared 包是共享工具包，提供了通用的工具函数
// 主要功能包括：配置加载、对象键生成等

const (
	DeafultKey = "default"
)

// LoadProperty 加载配置，参数优先级：CLI参数 > 环境变量 > 默认值
func LoadProperty(p *string, name string, defaultValue, usage string) {
	value, exists := os.LookupEnv(name)
	if !exists {
		*p = defaultValue
	} else {
		*p = value
	}
}
