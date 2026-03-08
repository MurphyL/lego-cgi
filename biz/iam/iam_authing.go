package iam

import "github.com/gofiber/fiber/v3"

// 检查权限
func RequirePrem(perms ...string) func(c fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		return nil
	}
}
