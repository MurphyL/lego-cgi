package corp

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v3"
)

// corp 模块是企业管理模块，用于管理企业的基础信息和核实状态
// 主要功能包括：企业基础信息查询、企业信息核实等

/**
 * CORP 是 "corporation" 的缩写，通常用于指代股份有限公司或法人实体。与 "company" 相比，"corporation" 更加特指由股东组成的公司，通常规模较大。
 */

// Corp 企业基础信息
type Corp struct {
	ID                  uint64       `json:"id"`
	Name                string       `json:"name"`                // 企业名称
	UnifiedCode         string       `json:"unifiedCode"`         // 统一社会信用代码
	LegalRepresentative string       `json:"legalRepresentative"` // 法定代表人
	RegisteredCapital   string       `json:"registeredCapital"`   // 注册资本
	EstablishDate       *time.Time   `json:"establishDate"`       // 成立日期
	BusinessScope       string       `json:"businessScope"`       // 经营范围
	RegisteredAddress   string       `json:"registeredAddress"`   // 注册地址
	Status              string       `json:"status"`              // 企业状态
	VerifyStatus        VerifyStatus `json:"verifyStatus"`        // 核实状态
	CreatedAt           time.Time    `json:"createdAt"`           // 创建时间
	UpdatedAt           time.Time    `json:"updatedAt"`           // 更新时间
}

// VerifyStatus 企业核实状态
type VerifyStatus uint8

const (
	VerifyStatusUnverified VerifyStatus = 0 // 未核实
	VerifyStatusVerifying  VerifyStatus = 1 // 核实中
	VerifyStatusVerified   VerifyStatus = 2 // 已核实
	VerifyStatusFailed     VerifyStatus = 3 // 核实失败
)

// CorpQueryRequest 企业查询请求
type CorpQueryRequest struct {
	Name        string `json:"name"`        // 企业名称
	UnifiedCode string `json:"unifiedCode"` // 统一社会信用代码
	LegalPerson string `json:"legalPerson"` // 法定代表人
	Page        int    `json:"page"`        // 页码
	PageSize    int    `json:"pageSize"`    // 每页大小
}

// CorpQueryResponse 企业查询响应
type CorpQueryResponse struct {
	Total int    `json:"total"` // 总数
	List  []Corp `json:"list"`  // 企业列表
}

// CorpVerifyRequest 企业核实请求
type CorpVerifyRequest struct {
	ID              uint64 `json:"id"`              // 企业ID
	UnifiedCode     string `json:"unifiedCode"`     // 统一社会信用代码
	BusinessLicense string `json:"businessLicense"` // 营业执照图片
	LegalIDCard     string `json:"legalIDCard"`     // 法定代表人身份证
}

// CorpVerifyResponse 企业核实响应
type CorpVerifyResponse struct {
	ID           uint64       `json:"id"`           // 企业ID
	VerifyStatus VerifyStatus `json:"verifyStatus"` // 核实状态
	Message      string       `json:"message"`      // 核实消息
}

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
			ID:                  1,
			Name:                "测试企业1",
			UnifiedCode:         "91110000MA12345678",
			LegalRepresentative: "张三",
			RegisteredCapital:   "1000万",
			EstablishDate:       &now,
			BusinessScope:       "技术开发、技术服务",
			RegisteredAddress:   "北京市海淀区",
			Status:              "存续",
			VerifyStatus:        VerifyStatusVerified,
			CreatedAt:           now,
			UpdatedAt:           now,
		},
		{
			ID:                  2,
			Name:                "测试企业2",
			UnifiedCode:         "91110000MA87654321",
			LegalRepresentative: "李四",
			RegisteredCapital:   "500万",
			EstablishDate:       &now,
			BusinessScope:       "销售、咨询",
			RegisteredAddress:   "上海市浦东新区",
			Status:              "存续",
			VerifyStatus:        VerifyStatusUnverified,
			CreatedAt:           now,
			UpdatedAt:           now,
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
		ID:                  1,
		Name:                "测试企业1",
		UnifiedCode:         "91110000MA12345678",
		LegalRepresentative: "张三",
		RegisteredCapital:   "1000万",
		EstablishDate:       &now,
		BusinessScope:       "技术开发、技术服务",
		RegisteredAddress:   "北京市海淀区",
		Status:              "存续",
		VerifyStatus:        VerifyStatusVerified,
		CreatedAt:           now,
		UpdatedAt:           now,
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
		ID:                  1,
		Name:                "测试企业1",
		UnifiedCode:         unifiedCode,
		LegalRepresentative: "张三",
		RegisteredCapital:   "1000万",
		EstablishDate:       &now,
		BusinessScope:       "技术开发、技术服务",
		RegisteredAddress:   "北京市海淀区",
		Status:              "存续",
		VerifyStatus:        VerifyStatusVerified,
		CreatedAt:           now,
		UpdatedAt:           now,
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
