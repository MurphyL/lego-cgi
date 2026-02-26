package misc

import (
	"os"
	"strings"
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

func UniqueId(items ...string) string {
	return strings.Join(items, ":")
}
