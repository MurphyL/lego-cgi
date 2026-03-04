package dal

import (
	"time"

	"murphyl.com/lego/fns/shared"
)

// StatusEnum 状态码
type StatusEnum uint8

const (
	StatusDisabled StatusEnum = 0 // 禁用
	StatusEnabled  StatusEnum = 1 // 启用
	StatusDeleted  StatusEnum = 2 // 逻辑删除
)

type BaseEntry struct {
	ID         uint64        `json:"id" gorm:"primaryKey"`
	Status     shared.Status `json:"status" gorm:"index,default:1"`
	CreateTime time.Time     `json:"createTime" gorm:"autoCreateTime"`
	UpdateTime time.Time     `json:"updateTime" gorm:"autoUpdateTime"`
}

func (s StatusEnum) IsEnabled() bool {
	return s == StatusEnabled
}

func (s StatusEnum) IsDisabled() bool {
	return s == StatusDisabled
}

func (s StatusEnum) IsDeleted() bool {
	return s == StatusDeleted
}
