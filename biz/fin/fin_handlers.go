package fin

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

// FinanceHandler 财务管理处理器
type FinanceHandler struct {
	db *gorm.DB
}

// NewFinanceHandler 创建财务管理处理器
func NewFinanceHandler(dao *gorm.DB) func(router fiber.Router) {
	return func(router fiber.Router) {
		h := &FinanceHandler{db: dao}
		h.RegisterRoutes(router)
	}
}

// RegisterRoutes 注册路由
func (h *FinanceHandler) RegisterRoutes(router fiber.Router) {
	// 账单管理
	// 财务报表
}
