package biz

import (
	"github.com/gofiber/fiber/v3"
	"murphyl.com/lego/biz/iam"
	"murphyl.com/lego/biz/system"
)

// UseIdentifyManager 身份管理模块
func UseIdentifyManager(router fiber.Router) {
	router.Get("/login", iam.GetUserProfileHandler)
	router.Get("/logout", iam.GetUserProfileHandler)
	router.Get("/profile", iam.GetUserProfileHandler)
	router.Get("/reset-password", iam.GetUserProfileHandler)
	router.Get("/captcha", iam.GetUserProfileHandler)
}

func UseSystemDictManager(router fiber.Router) {
	router.Get("/dict/items", system.SearchDictTypeHandler)
}
