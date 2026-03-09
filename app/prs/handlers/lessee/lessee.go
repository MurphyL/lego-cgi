package tenant

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

/**
 * Tenant 租户
 */

// TenantHandler 租户处理器
type TenantHandler struct {
	db *gorm.DB
}

// NewTenantHandler 创建租户处理器
func NewTenantHandler(dao *gorm.DB) func(router fiber.Router) {
	return func(router fiber.Router) {
		h := &TenantHandler{db: dao}
		h.RegisterRoutes(router)
	}
}

// RegisterRoutes 注册路由
func (h *TenantHandler) RegisterRoutes(router fiber.Router) {
}

// Tenant 租户基础信息
type Tenant struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	TenantCode string `gorm:"size:30;uniqueIndex" json:"tenant_code"`
	Name       string `gorm:"size:50" json:"name"`
}

func (*Tenant) TableName() string {
	return "hrs_tenant"
}
