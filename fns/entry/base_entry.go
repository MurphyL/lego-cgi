package entry

import "time"

// Status 状态枚举
type StatusEnum uint8

type BaseEntry struct {
	ID         uint64     `json:"id" gorm:"primaryKey"`
	Status     StatusEnum `json:"status" gorm:"index,default:1"`
	CreateTime time.Time  `json:"createTime" gorm:"autoCreateTime"`
	UpdateTime time.Time  `json:"updateTime" gorm:"autoUpdateTime"`
}

// Status 枚举
const (
	StatusDisabled StatusEnum = 0 // 禁用
	StatusEnabled  StatusEnum = 1 // 启用
	StatusDeleted  StatusEnum = 2 // 逻辑删除
)

// IsEnabled 是否启用
func (s StatusEnum) IsEnabled() bool {
	return s == StatusEnabled
}

// IsDisabled 是否禁用
func (s StatusEnum) IsDisabled() bool {
	return s == StatusDisabled
}

// IsDeleted 是否逻辑删除
func (s StatusEnum) IsDeleted() bool {
	return s == StatusDeleted
}
