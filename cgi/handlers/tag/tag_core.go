package tag

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
	"murphyl.com/lego/cgi"
	"murphyl.com/lego/fns/entry"
)

// Tag - 打标签、做标记（通用）

// TagTypeEnum 标签类型
type TagTypeEnum string

const (
	TypeSystem    TagTypeEnum = "system"     // 系统标签
	TypeManual    TagTypeEnum = "manual"     // 手动标签
	TypeRuleBased TagTypeEnum = "rule_based" // 规则标签
)

func NewTagHandler(db *gorm.DB) *tagHandler {
	return &tagHandler{db: db}
}

type tagHandler struct {
	db *gorm.DB
}

func (h *tagHandler) RegisterRoutes(router fiber.Router) {
	router.Get("/tags", func(c fiber.Ctx) error {
		return cgi.RetrieveAll[struct{ ID uint64 }, Tag](c, h.db)
	})
	router.Get("/tags/:id", func(c fiber.Ctx) error {
		return cgi.RetrieveOne[struct{ ID uint64 }, Tag](c, h.db)
	})
}

// Tag 标签
type Tag struct {
	entry.BaseEntry
	Code        string      `json:"code"`
	Name        string      `json:"name"`
	Description string      `json:"description,omitempty"`
	TagType     TagTypeEnum `json:"type"`
	Weight      int         `json:"weight"`
}

// PeriodTag 效期标签
type PeriodTag struct {
	entry.BaseEntry
	entry.PeriodEntry
	TagID uint `json:"tag_id"`
}

func (t *Tag) TableName() string {
	return "lego_tag_info"
}

func (t *PeriodTag) TableName() string {
	return "lego_tag_period"
}

// IsValid 效期标签是否有效
func (t *PeriodTag) IsValid() bool {
	return t.Status.IsEnabled() && !t.IsExpired()
}
