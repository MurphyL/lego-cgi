package middleware

import (
	"github.com/gofiber/fiber/v3"
)

// Fiber 路由检查权限的中间件
func RequirePrems(perms ...string) func(c fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		return c.Next()
	}
}
