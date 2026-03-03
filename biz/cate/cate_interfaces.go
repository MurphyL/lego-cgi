package cate

import (
	"time"

	"murphyl.com/lego/dal"
)

// Category 分类定义
type Category struct {
	ID          uint64         `json:"id"`
	ParentID    *uint64        `json:"parentId"` // 父分类ID，根分类为nil
	Code        string         `json:"code"`
	Name        string         `json:"name"`
	Description string         `json:"description,omitempty"`
	Level       int            `json:"level"`  // 分类层级
	Path        string         `json:"path"`   // 分类路径，如：1/2/3
	Weight      int            `json:"weight"` // 排序权重
	Status      dal.StatusEnum `json:"status"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
}

// CategoryWithChildren 带子分类的分类定义
type CategoryWithChildren struct {
	Category
	Children []CategoryWithChildren `json:"children,omitempty"`
}

// IsValid 检查分类是否有效
func (c *Category) IsValid() bool {
	return c.Status.IsEnabled()
}
