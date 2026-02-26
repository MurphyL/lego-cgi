package idm

import (
	"github.com/gofiber/fiber/v3"

	"murphyl.com/app/idm/handlers"
)

/* 身份管理模块 */

// UseIdentifyManager 身份管理模块
func UseIdentifyManager(router fiber.Router) {
	router.Get("/login", handlers.GetUserProfileHandler)
	router.Get("/logout", handlers.GetUserProfileHandler)
	router.Get("/profile", handlers.GetUserProfileHandler)
	router.Get("/reset-password", handlers.GetUserProfileHandler)
	router.Get("/captcha", handlers.GetUserProfileHandler)
}
