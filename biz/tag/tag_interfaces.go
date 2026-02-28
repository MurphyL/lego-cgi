package tag

import (
	"murphyl.com/lego/dal"
	"murphyl.com/lego/udf"
)

// Type 标签类型
type Type uint8

const (
	TypeSystem    Type = 1 // 系统标签
	TypeManual    Type = 2 // 手动标签
	TypeRuleBased Type = 3 // 规则标签
)

// Tag 标签定义
type Tag struct {
	dal.Model
	udf.PeriodValid
	ID          uint64         `json:"id"`
	Code        string         `json:"code"`
	Name        string         `json:"name"`
	Description string         `json:"description,omitempty"`
	Type        Type           `json:"type"`
	Weight      int            `json:"weight"`
	Status      dal.StatusEnum `json:"status"`
}


// IsValid 检查标签是否有效
func (t *Tag) IsValid() bool {
	return t.Status.IsEnabled() && !t.IsExpired()
}
