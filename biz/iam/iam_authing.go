package iam

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

// Fiber 路由检查权限的中间件
func RequirePrems(perms ...string) func(c fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		return nil
	}
}

type AuthingHandler struct {
	db *gorm.DB
}

func NewAuthingHandler(dao *gorm.DB) *AuthingHandler {
	return &AuthingHandler{db: dao}
}

func (ah *AuthingHandler) RegisterRoutes(router fiber.Router) {
}
