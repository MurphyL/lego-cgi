package biz

import (
	"github.com/gofiber/fiber/v3"
	"murphyl.com/lego/biz/cate"
	"murphyl.com/lego/biz/iam"
	"murphyl.com/lego/biz/system"
	"murphyl.com/lego/biz/tag"
)

// UseIdentifyManager 身份管理模块
func UseIdentifyManager(router fiber.Router) {
	router.Post("/login", iam.LoginHandler)
	router.Post("/logout", iam.LogoutHandler)
	router.Get("/profile", iam.GetUserProfileHandler)
	router.Post("/reset-password", iam.ResetPasswordHandler)
	router.Post("/captcha", iam.CaptchaHandler)
}

// UseRBACManager RBAC管理模块
func UseRBACManager(router fiber.Router) {
	// 角色管理
	router.Post("/roles", iam.CreateRoleHandler)
	router.Put("/roles/:id", iam.UpdateRoleHandler)
	router.Delete("/roles/:id", iam.DeleteRoleHandler)
	router.Get("/roles/:id", iam.GetRoleHandler)
	router.Get("/roles", iam.ListRolesHandler)

	// 权限管理
	router.Post("/permissions", iam.CreatePermHandler)
	router.Put("/permissions/:id", iam.UpdatePermHandler)
	router.Delete("/permissions/:id", iam.DeletePermHandler)
	router.Get("/permissions/:id", iam.GetPermHandler)
	router.Get("/permissions", iam.ListPermsHandler)

	// 用户-角色关联
	router.Post("/user-roles", iam.AssignRoleToUserHandler)
	router.Delete("/user-roles", iam.RemoveRoleFromUserHandler)

	// 角色-权限关联
	router.Post("/role-permissions", iam.AssignPermToRoleHandler)
	router.Delete("/role-permissions", iam.RemovePermFromRoleHandler)
}

// UseTenantManager 租户管理模块
func UseTenantManager(router fiber.Router) {
	// 租户管理
	router.Post("/tenants", iam.CreateTenantHandler)
	router.Put("/tenants/:id", iam.UpdateTenantHandler)
	router.Delete("/tenants/:id", iam.DeleteTenantHandler)
	router.Get("/tenants/:id", iam.GetTenantHandler)
	router.Get("/tenants", iam.ListTenantsHandler)

	// 租户成员管理
	router.Post("/tenant-members", iam.AddTenantMemberHandler)
	router.Put("/tenant-members/:id", iam.UpdateTenantMemberHandler)
	router.Delete("/tenant-members", iam.RemoveTenantMemberHandler)
	router.Get("/tenants/:tenantId/members", iam.ListTenantMembersHandler)
}

func UseSystemDictManager(router fiber.Router) {
	router.Get("/dict/items", system.SearchDictTypeHandler)
}

// UseTagManager 标签管理模块
func UseTagManager(router fiber.Router) {
	router.Post("/tags", tag.CreateTagHandler)
	router.Put("/tags/:id", tag.UpdateTagHandler)
	router.Delete("/tags/:id", tag.DeleteTagHandler)
	router.Get("/tags/:id", tag.GetTagHandler)
	router.Get("/tags", tag.ListTagsHandler)
}

// UseCategoryManager 分类管理模块
func UseCategoryManager(router fiber.Router) {
	router.Post("/categories", cate.CreateCategoryHandler)
	router.Put("/categories/:id", cate.UpdateCategoryHandler)
	router.Delete("/categories/:id", cate.DeleteCategoryHandler)
	router.Get("/categories/:id", cate.GetCategoryHandler)
	router.Get("/categories", cate.ListCategoriesHandler)
	router.Get("/categories/tree", cate.GetCategoryTreeHandler)
}
