package corp

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v3"
)

// ListCorpsHandler 查询企业列表
func ListCorpsHandler(c fiber.Ctx) error {
	var req CorpQueryRequest
	body := c.Body()
	if err := json.Unmarshal(body, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// 实际应用中应该调用服务层查询企业列表
	// response, err := corpService.ListCorps(req)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	// 模拟企业列表
	now := time.Now()
	corps := []Corp{
		{
			ID:                1,
			Name:              "测试企业1",
			UnifiedCode:       "91110000MA12345678",
			LegalRepresentative: "张三",
			RegisteredCapital: "1000万",
			EstablishDate:     &now,
			BusinessScope:     "技术开发、技术服务",
			RegisteredAddress: "北京市海淀区",
			Status:            "存续",
			VerifyStatus:      VerifyStatusVerified,
			CreatedAt:         now,
			UpdatedAt:         now,
		},
		{
			ID:                2,
			Name:              "测试企业2",
			UnifiedCode:       "91110000MA87654321",
			LegalRepresentative: "李四",
			RegisteredCapital: "500万",
			EstablishDate:     &now,
			BusinessScope:     "销售、咨询",
			RegisteredAddress: "上海市浦东新区",
			Status:            "存续",
			VerifyStatus:      VerifyStatusUnverified,
			CreatedAt:         now,
			UpdatedAt:         now,
		},
	}

	response := CorpQueryResponse{
		Total: len(corps),
		List:  corps,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// GetCorpByIdHandler 根据ID查询企业详情
func GetCorpByIdHandler(c fiber.Ctx) error {
	// id := c.Params("id")

	// 实际应用中应该调用服务层查询企业详情
	// corp, err := corpService.GetCorpById(id)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	// 模拟企业详情
	now := time.Now()
	corp := Corp{
		ID:                1,
		Name:              "测试企业1",
		UnifiedCode:       "91110000MA12345678",
		LegalRepresentative: "张三",
		RegisteredCapital: "1000万",
		EstablishDate:     &now,
		BusinessScope:     "技术开发、技术服务",
		RegisteredAddress: "北京市海淀区",
		Status:            "存续",
		VerifyStatus:      VerifyStatusVerified,
		CreatedAt:         now,
		UpdatedAt:         now,
	}

	return c.Status(fiber.StatusOK).JSON(corp)
}

// GetCorpByUnifiedCodeHandler 根据统一社会信用代码查询企业
func GetCorpByUnifiedCodeHandler(c fiber.Ctx) error {
	unifiedCode := c.Query("unifiedCode")

	// 实际应用中应该调用服务层查询企业
	// corp, err := corpService.GetCorpByUnifiedCode(unifiedCode)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	// 模拟企业详情
	now := time.Now()
	corp := Corp{
		ID:                1,
		Name:              "测试企业1",
		UnifiedCode:       unifiedCode,
		LegalRepresentative: "张三",
		RegisteredCapital: "1000万",
		EstablishDate:     &now,
		BusinessScope:     "技术开发、技术服务",
		RegisteredAddress: "北京市海淀区",
		Status:            "存续",
		VerifyStatus:      VerifyStatusVerified,
		CreatedAt:         now,
		UpdatedAt:         now,
	}

	return c.Status(fiber.StatusOK).JSON(corp)
}

// VerifyCorpHandler 企业信息核实
func VerifyCorpHandler(c fiber.Ctx) error {
	var req CorpVerifyRequest
	body := c.Body()
	if err := json.Unmarshal(body, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// 实际应用中应该调用服务层进行企业核实
	// response, err := corpService.VerifyCorp(req)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	// 模拟企业核实
	response := CorpVerifyResponse{
		ID:           req.ID,
		VerifyStatus: VerifyStatusVerified,
		Message:      "企业信息核实成功",
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
