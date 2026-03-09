package iam

import (
	"time"

	"murphyl.com/lego/fns/entry"
)

/* 基于角色的访问控制模块 */

type Role struct {
	entry.BaseEntry
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	TenantID    uint64    `json:"tenantId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type User struct {
	entry.BaseEntry
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	TenantID uint64 `json:"tenantId"`
}

type Perm struct {
	entry.BaseEntry
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`   // menu, button, api
	Path        string `json:"path"`   // 菜单路径或API路径
	Method      string `json:"method"` // HTTP方法
}

// 角色-权限关联
type RolePerm struct {
	entry.BaseEntry
	RoleID uint64 `json:"roleId"`
	PermID uint64 `json:"permId"`
}

// 用户-角色关联
type UserRole struct {
	entry.BaseEntry
	UserID   uint64 `json:"userId"`
	RoleID   uint64 `json:"roleId"`
	TenantID uint64 `json:"tenantId"`
}
