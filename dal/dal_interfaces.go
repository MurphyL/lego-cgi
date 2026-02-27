package dal

import (
	"time"
)

// StatusEnum 状态码
type StatusEnum uint8

const (
	StatusDisabled StatusEnum = 0 // 禁用
	StatusEnabled  StatusEnum = 1 // 启用
	StatusDeleted  StatusEnum = 2 // 逻辑删除
)

type Model struct {
	Id        uint64     `json:"id"`
	Status    StatusEnum `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
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
