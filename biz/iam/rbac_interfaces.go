package iam

import (
	"time"

	"murphyl.com/lego/biz/system"
)

/* 基于角色的访问控制模块 */

type Role struct {
	ID          uint64         `json:"id"`
	Code        string         `json:"code"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Status      system.Status  `json:"status"`
	TenantID    uint64         `json:"tenantId"`
	Scope       system.ScopeEntry `json:"scope"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
}

type User struct {
	ID          uint64         `json:"id"`
	Username    string         `json:"username"`
	Name        string         `json:"name"`
	Email       string         `json:"email"`
	Mobile      string         `json:"mobile"`
	Status      system.Status  `json:"status"`
	TenantID    uint64         `json:"tenantId"`
	Scope       system.ScopeEntry `json:"scope"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
}

type Perm struct {
	ID          uint64         `json:"id"`
	Code        string         `json:"code"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Type        string         `json:"type"` // menu, button, api
	Path        string         `json:"path"` // 菜单路径或API路径
	Method      string         `json:"method"` // HTTP方法
	Status      system.Status  `json:"status"`
	Scope       system.ScopeEntry `json:"scope"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
}

// 角色-权限关联
 type RolePerm struct {
	ID        uint64    `json:"id"`
	RoleID    uint64    `json:"roleId"`
	PermID    uint64    `json:"permId"`
	CreatedAt time.Time `json:"createdAt"`
}

// 用户-角色关联
 type UserRole struct {
	ID        uint64    `json:"id"`
	UserID    uint64    `json:"userId"`
	RoleID    uint64    `json:"roleId"`
	TenantID  uint64    `json:"tenantId"`
	CreatedAt time.Time `json:"createdAt"`
}

type GetById[K any, R Role | User | Perm] func(K) *R
