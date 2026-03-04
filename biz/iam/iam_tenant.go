package iam

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v3"
	"murphyl.com/lego/fns/shared"
)

/** 租户 **/

type Tenant struct {
	ID          uint64        `json:"id"`
	Name        string        `json:"name"`
	Code        string        `json:"code"`
	Description string        `json:"description"`
	Status      shared.Status `json:"status"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
	ExpiredAt   *time.Time    `json:"expiredAt"`
}

// 租户成员
type TenantMember struct {
	ID        uint64        `json:"id"`
	TenantID  uint64        `json:"tenantId"`
	UserID    uint64        `json:"userId"`
	RoleID    uint64        `json:"roleId"`
	Status    shared.Status `json:"status"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
}

// TenantRequest 租户请求
type TenantRequest struct {
	Name        string     `json:"name"`
	Code        string     `json:"code"`
	Description string     `json:"description"`
	Status      int        `json:"status"`
	ExpiredAt   *time.Time `json:"expiredAt"`
}

// TenantMemberRequest 租户成员请求
type TenantMemberRequest struct {
	TenantID uint64 `json:"tenantId"`
	UserID   uint64 `json:"userId"`
	RoleID   uint64 `json:"roleId"`
	Status   int    `json:"status"`
}

// CreateTenantHandler 创建租户
func CreateTenantHandler(c fiber.Ctx) error {
	var req TenantRequest
	body := c.Body()
	if err := json.Unmarshal(body, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// 实际应用中应该调用服务层创建租户
	// tenant, err := tenantService.CreateTenant(req)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	// 模拟租户创建
	tenant := Tenant{
		ID:          1,
		Name:        req.Name,
		Code:        req.Code,
		Description: req.Description,
		Status:      1,
		ExpiredAt:   req.ExpiredAt,
	}

	return c.Status(fiber.StatusOK).JSON(tenant)
}

// UpdateTenantHandler 更新租户
func UpdateTenantHandler(c fiber.Ctx) error {
	var req TenantRequest
	body := c.Body()
	if err := json.Unmarshal(body, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// 实际应用中应该调用服务层更新租户
	// tenant, err := tenantService.UpdateTenant(req)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	// 模拟租户更新
	tenant := Tenant{
		ID:          1,
		Name:        req.Name,
		Code:        req.Code,
		Description: req.Description,
		Status:      1,
		ExpiredAt:   req.ExpiredAt,
	}

	return c.Status(fiber.StatusOK).JSON(tenant)
}

// DeleteTenantHandler 删除租户
func DeleteTenantHandler(c fiber.Ctx) error {
	// id := c.Params("id")

	// 实际应用中应该调用服务层删除租户
	// err := tenantService.DeleteTenant(id)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
}

// GetTenantHandler 获取租户
func GetTenantHandler(c fiber.Ctx) error {
	// id := c.Params("id")

	// 实际应用中应该调用服务层获取租户
	// tenant, err := tenantService.GetTenant(id)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	// 模拟租户获取
	tenant := Tenant{
		ID:          1,
		Name:        "测试租户",
		Code:        "test",
		Description: "测试租户",
		Status:      1,
	}

	return c.Status(fiber.StatusOK).JSON(tenant)
}

// ListTenantsHandler 列出租户
func ListTenantsHandler(c fiber.Ctx) error {
	// 实际应用中应该调用服务层列出租户
	// tenants, err := tenantService.ListTenants()
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	// 模拟租户列表
	tenants := []Tenant{
		{
			ID:          1,
			Name:        "测试租户1",
			Code:        "test1",
			Description: "测试租户1",
			Status:      1,
		},
		{
			ID:          2,
			Name:        "测试租户2",
			Code:        "test2",
			Description: "测试租户2",
			Status:      1,
		},
	}

	return c.Status(fiber.StatusOK).JSON(tenants)
}

// AddTenantMemberHandler 添加租户成员
func AddTenantMemberHandler(c fiber.Ctx) error {
	var req TenantMemberRequest
	body := c.Body()
	if err := json.Unmarshal(body, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// 实际应用中应该调用服务层添加租户成员
	// member, err := tenantService.AddTenantMember(req)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	// 模拟租户成员添加
	member := TenantMember{
		ID:       1,
		TenantID: req.TenantID,
		UserID:   req.UserID,
		RoleID:   req.RoleID,
		Status:   1,
	}

	return c.Status(fiber.StatusOK).JSON(member)
}

// UpdateTenantMemberHandler 更新租户成员
func UpdateTenantMemberHandler(c fiber.Ctx) error {
	var req TenantMemberRequest
	body := c.Body()
	if err := json.Unmarshal(body, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// 实际应用中应该调用服务层更新租户成员
	// member, err := tenantService.UpdateTenantMember(req)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	// 模拟租户成员更新
	member := TenantMember{
		ID:       1,
		TenantID: req.TenantID,
		UserID:   req.UserID,
		RoleID:   req.RoleID,
		Status:   1,
	}

	return c.Status(fiber.StatusOK).JSON(member)
}

// RemoveTenantMemberHandler 移除租户成员
func RemoveTenantMemberHandler(c fiber.Ctx) error {
	var req TenantMemberRequest
	body := c.Body()
	if err := json.Unmarshal(body, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// 实际应用中应该调用服务层移除租户成员
	// err := tenantService.RemoveTenantMember(req)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
}

// ListTenantMembersHandler 列出租户成员
func ListTenantMembersHandler(c fiber.Ctx) error {
	// tenantID := c.Params("tenantId")

	// 实际应用中应该调用服务层列出租户成员
	// members, err := tenantService.ListTenantMembers(tenantID)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	// 模拟租户成员列表
	members := []TenantMember{
		{
			ID:       1,
			TenantID: 1,
			UserID:   1,
			RoleID:   1,
			Status:   1,
		},
		{
			ID:       2,
			TenantID: 1,
			UserID:   2,
			RoleID:   2,
			Status:   1,
		},
	}

	return c.Status(fiber.StatusOK).JSON(members)
}
