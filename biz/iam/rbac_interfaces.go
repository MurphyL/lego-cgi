package iam

import "murphyl.com/lego/biz/system"

/* 基于角色的访问控制模块 */

type Role struct {
	system.ScopeEntry
	Code string
	Name string
}

type User struct {
	system.ScopeEntry
	Name string
}

type Perm struct {
	system.ScopeEntry
	Name string
}

type GetById[K any, R Role | User | Perm] func(K) *R
