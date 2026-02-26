package misc

import (
	"os"
	"strings"
)

func LoadProperty(p *string, name string, defaultValue, usage string) {
	value, exists := os.LookupEnv(name)
	if !exists {
		p = &defaultValue
	} else {
		p = &value
	}
}

func UniqueId(items ...string) string {
	return strings.Join(items, ":")
}
