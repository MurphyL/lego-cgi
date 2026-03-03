package iam

import (
	"time"

	"murphyl.com/lego/biz/system"
)

/** 租户支持 **/

type Tenant struct {
	ID          uint64         `json:"id"`
	Name        string         `json:"name"`
	Code        string         `json:"code"`
	Description string         `json:"description"`
	Status      system.Status  `json:"status"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	ExpiredAt   *time.Time     `json:"expiredAt"`
}

// 租户成员
 type TenantMember struct {
	ID        uint64    `json:"id"`
	TenantID  uint64    `json:"tenantId"`
	UserID    uint64    `json:"userId"`
	RoleID    uint64    `json:"roleId"`
	Status    system.Status `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
