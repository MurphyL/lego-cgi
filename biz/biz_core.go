package biz

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"murphyl.com/lego/biz/cate"
	"murphyl.com/lego/biz/corp"
	"murphyl.com/lego/biz/iam"
	"murphyl.com/lego/biz/tag"
)

// biz 模块是通用业务模块，包含了各种业务逻辑的实现
// 主要包括：身份管理、RBAC权限控制、租户管理、数据字典管理、标签管理、分类管理和企业管理

// UseRBACManager RBAC管理模块
func UseRBACManager(dao *gorm.DB) func(router fiber.Router) {
	return func(router fiber.Router) {
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
}

// UseTenantManager 租户管理模块
func UseTenantManager(dao *gorm.DB) func(router fiber.Router) {
	return func(router fiber.Router) {
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
}

// UseTagManager 标签管理模块
func UseTagManager(dao *gorm.DB) func(router fiber.Router) {
	return func(router fiber.Router) {
		router.Post("/tags", tag.CreateTagHandler)
		router.Put("/tags/:id", tag.UpdateTagHandler)
		router.Delete("/tags/:id", tag.DeleteTagHandler)
		router.Get("/tags/:id", tag.GetTagHandler)
		router.Get("/tags", tag.ListTagsHandler)
	}
}

// UseCategoryManager 分类管理模块
func UseCategoryManager(dao *gorm.DB) func(router fiber.Router) {
	return func(router fiber.Router) {
		router.Post("/categories", cate.CreateCategoryHandler)
		router.Put("/categories/:id", cate.UpdateCategoryHandler)
		router.Delete("/categories/:id", cate.DeleteCategoryHandler)
		router.Get("/categories/:id", cate.GetCategoryHandler)
		router.Get("/categories", cate.ListCategoriesHandler)
		router.Get("/categories/tree", cate.GetCategoryTreeHandler)
	}
}

// UseCorpManager 企业管理模块
func UseCorpManager(dao *gorm.DB) func(router fiber.Router) {
	return func(router fiber.Router) {
		router.Post("/corps", corp.ListCorpsHandler)
		router.Get("/corps/:id", corp.GetCorpByIdHandler)
		router.Get("/corps/by-code", corp.GetCorpByUnifiedCodeHandler)
		router.Post("/corps/verify", corp.VerifyCorpHandler)
	}
}
