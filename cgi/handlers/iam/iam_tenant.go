package iam

import (
	"murphyl.com/lego/fns/entry"
)

/** 租户 **/

type Tenant struct {
	entry.BaseEntry
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

// 租户成员
type TenantMember struct {
	entry.BaseEntry
	TenantID uint64 `json:"tenantId"`
	UserID   uint64 `json:"userId"`
	RoleID   uint64 `json:"roleId"`
}
