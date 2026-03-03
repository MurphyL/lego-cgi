package corp

import (
	"time"
)

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
