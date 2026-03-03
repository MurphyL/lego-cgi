package biz

import (
	"github.com/gofiber/fiber/v3"
	"murphyl.com/lego/biz/cate"
	"murphyl.com/lego/biz/corp"
	"murphyl.com/lego/biz/excel"
	"murphyl.com/lego/biz/iam"
	"murphyl.com/lego/biz/system"
	"murphyl.com/lego/biz/tag"
)

// biz 模块是通用业务模块，包含了各种业务逻辑的实现
// 主要包括：身份管理、RBAC权限控制、租户管理、数据字典管理、标签管理、分类管理和企业管理

// UseIdentifyManager 身份管理模块
func UseIdentifyManager(router fiber.Router) {
	router.Post("/login", iam.LoginHandler)
	router.Post("/logout", iam.LogoutHandler)
	router.Get("/profile", iam.GetUserProfileHandler)
	router.Put("/profile", iam.UpdateProfileHandler)
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

// UseSystemDictManager 数据字典管理模块
func UseSystemDictManager(router fiber.Router) {
	// 字典类型管理
	router.Post("/dict/types", system.CreateDictTypeHandler)
	router.Put("/dict/types", system.UpdateDictTypeHandler)
	router.Delete("/dict/types/:dictCode", system.DeleteDictTypeHandler)
	router.Get("/dict/types/:dictCode", system.GetDictTypeHandler)
	router.Get("/dict/types", system.ListDictTypesHandler)

	// 字典项管理
	router.Post("/dict/items", system.CreateDictItemHandler)
	router.Put("/dict/items", system.UpdateDictItemHandler)
	router.Delete("/dict/items/:id", system.DeleteDictItemHandler)
	router.Get("/dict/items/:id", system.GetDictItemHandler)
	router.Get("/dict/items", system.ListDictItemsHandler)

	// 字典组管理
	router.Get("/dict/groups/:dictCode", system.GetDictGroupHandler)
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

// UseCorpManager 企业管理模块
func UseCorpManager(router fiber.Router) {
	router.Post("/corps", corp.ListCorpsHandler)
	router.Get("/corps/:id", corp.GetCorpByIdHandler)
	router.Get("/corps/by-code", corp.GetCorpByUnifiedCodeHandler)
	router.Post("/corps/verify", corp.VerifyCorpHandler)
}

// UseExcelManager Excel导出模块
func UseExcelManager(router fiber.Router) {
	router.Post("/excel/export", excel.ExportExcelHandler)
	router.Get("/excel/demo", excel.ExportDemoHandler)
}
