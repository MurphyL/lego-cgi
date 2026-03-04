package contract

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
	"murphyl.com/app/hrs/middleware"
	"murphyl.com/lego/cgi"
	"murphyl.com/lego/dal"
)

/**
 * Contract 合同
 */

// Contract 合同基础信息
type Contract struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	ContractCode    string    `gorm:"size:30;uniqueIndex" json:"contract_code"`
	PropertyID      uint      `json:"property_id"`
	TenantID        uint      `json:"tenant_id"`
	TemplateID      uint      `json:"template_id"`
	ContractName    string    `gorm:"size:100" json:"contract_name"`
	StartDate       time.Time `json:"start_date"`
	EndDate         time.Time `json:"end_date"`
	RentAmount      float64   `json:"rent_amount"`
	DepositAmount   float64   `json:"deposit_amount"`
	PaymentMethod   uint8     `json:"payment_method"`
	Status          uint8     `json:"status"`
	ContractFileURL string    `gorm:"size:500" json:"contract_file_url"`
	CreatorID       uint      `json:"creator_id"`
	CreateTime      time.Time `json:"create_time"`
	UpdateTime      time.Time `json:"update_time"`
}

func (*Contract) TableName() string {
	return "hrs_contract"
}

// ContractTemplate 合同模板
type ContractTemplate struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	TemplateName    string    `gorm:"size:100" json:"template_name"`
	TemplateContent string    `gorm:"type:text" json:"template_content"`
	Status          uint8     `json:"status"`
	CreateTime      time.Time `json:"create_time"`
	UpdateTime      time.Time `json:"update_time"`
}

func (*ContractTemplate) TableName() string {
	return "hrs_contract_template"
}

// ContractClause 合同条款
type ContractClause struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	ContractID    uint      `json:"contract_id"`
	ClauseType    string    `gorm:"size:50" json:"clause_type"`
	ClauseContent string    `gorm:"type:text" json:"clause_content"`
	CreateTime    time.Time `json:"create_time"`
}

func (*ContractClause) TableName() string {
	return "hrs_contract_clause"
}

// ContractSignature 合同签章
type ContractSignature struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	ContractID   uint      `json:"contract_id"`
	SignerType   uint8     `json:"signer_type"`
	SignerID     uint      `json:"signer_id"`
	SignerName   string    `gorm:"size:50" json:"signer_name"`
	SignatureURL string    `gorm:"size:500" json:"signature_url"`
	SignTime     time.Time `json:"sign_time"`
}

func (*ContractSignature) TableName() string {
	return "hrs_contract_signature"
}

// ContractStatusLog 合同状态变更日志
type ContractStatusLog struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	ContractID   uint      `json:"contract_id"`
	OldStatus    uint8     `json:"old_status"`
	NewStatus    uint8     `json:"new_status"`
	ChangeReason string    `gorm:"size:255" json:"change_reason"`
	OperatorID   uint      `json:"operator_id"`
	CreateTime   time.Time `json:"create_time"`
}

func (*ContractStatusLog) TableName() string {
	return "hrs_contract_status_log"
}

// ContractRisk 合同风险预警
type ContractRisk struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	ContractID  uint      `json:"contract_id"`
	RiskType    uint8     `json:"risk_type"`
	RiskLevel   uint8     `json:"risk_level"`
	RiskMessage string    `gorm:"size:255" json:"risk_message"`
	Status      uint8     `json:"status"`
	CreateTime  time.Time `json:"create_time"`
	UpdateTime  time.Time `json:"update_time"`
}

func (*ContractRisk) TableName() string {
	return "hrs_contract_risk"
}

// CreateContractRequest 创建合同请求
type CreateContractRequest struct {
	PropertyID    uint      `json:"property_id" validate:"required"`
	TenantID      uint      `json:"tenant_id" validate:"required"`
	TemplateID    uint      `json:"template_id" validate:"required"`
	ContractName  string    `json:"contract_name" validate:"required"`
	StartDate     time.Time `json:"start_date" validate:"required"`
	EndDate       time.Time `json:"end_date" validate:"required"`
	RentAmount    float64   `json:"rent_amount" validate:"required"`
	DepositAmount float64   `json:"deposit_amount" validate:"required"`
	PaymentMethod uint8     `json:"payment_method" validate:"required"`
	Clauses       []Clause  `json:"clauses"`
}

// Clause 条款信息
type Clause struct {
	ClauseType    string `json:"clause_type" validate:"required"`
	ClauseContent string `json:"clause_content" validate:"required"`
}

// UpdateContractRequest 更新合同请求
type UpdateContractRequest struct {
	ContractName  string    `json:"contract_name"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	RentAmount    float64   `json:"rent_amount"`
	DepositAmount float64   `json:"deposit_amount"`
	PaymentMethod uint8     `json:"payment_method"`
	Clauses       []Clause  `json:"clauses"`
}

// ChangeContractStatusRequest 变更合同状态请求
type ChangeContractStatusRequest struct {
	NewStatus    uint8  `json:"new_status" validate:"required"`
	ChangeReason string `json:"change_reason" validate:"required"`
}

// CreateSignatureRequest 创建签章请求
type CreateSignatureRequest struct {
	SignerType   uint8  `json:"signer_type" validate:"required"`
	SignerID     uint   `json:"signer_id" validate:"required"`
	SignerName   string `json:"signer_name" validate:"required"`
	SignatureURL string `json:"signature_url" validate:"required"`
}

// ContractQueryRequest 合同查询请求
type ContractQueryRequest struct {
	PropertyID   *uint      `json:"property_id"`
	TenantID     *uint      `json:"tenant_id"`
	Status       *uint8     `json:"status"`
	StartDateMin *time.Time `json:"start_date_min"`
	StartDateMax *time.Time `json:"start_date_max"`
	Page         int        `json:"page" default:"1"`
	PageSize     int        `json:"page_size" default:"10"`
}

// ContractHandler 合同处理器
type ContractHandler struct {
	db *gorm.DB
}

// NewContractHandler 创建合同处理器
func NewContractHandler() *ContractHandler {
	return &ContractHandler{}
}

// GetDataAccessLayer 获取数据访问层
func (h *ContractHandler) GetDataAccessLayer(ctx fiber.Ctx) dal.DataAccessLayer {
	return cgi.DefaultDataAccessLayer(ctx)
}

// RegisterRoutes 注册路由
func (h *ContractHandler) RegisterRoutes(router fiber.Router) {
	// 合同模板管理
	router.Post("/contract-templates", middleware.AuthMiddleware("contract:create_template"), h.CreateContractTemplate)
	router.Get("/contract-templates", middleware.AuthMiddleware("contract:list_templates"), h.ListContractTemplates)
	router.Get("/contract-templates/:id", middleware.AuthMiddleware("contract:view_template"), h.GetContractTemplate)
	router.Put("/contract-templates/:id", middleware.AuthMiddleware("contract:update_template"), h.UpdateContractTemplate)
	router.Delete("/contract-templates/:id", middleware.AuthMiddleware("contract:delete_template"), h.DeleteContractTemplate)

	// 合同管理
	router.Post("/contracts", middleware.AuthMiddleware("contract:create"), h.CreateContract)
	router.Get("/contracts", middleware.AuthMiddleware("contract:list"), h.ListContracts)
	router.Get("/contracts/:id", middleware.AuthMiddleware("contract:view"), h.GetContract)
	router.Put("/contracts/:id", middleware.AuthMiddleware("contract:update"), h.UpdateContract)
	router.Delete("/contracts/:id", middleware.AuthMiddleware("contract:delete"), h.DeleteContract)
	router.Put("/contracts/:id/status", middleware.AuthMiddleware("contract:update_status"), h.ChangeContractStatus)

	// 合同签章
	router.Post("/contracts/:id/signatures", middleware.AuthMiddleware("contract:create_signature"), h.CreateSignature)
	router.Get("/contracts/:id/signatures", middleware.AuthMiddleware("contract:list_signatures"), h.ListSignatures)

	// 合同风险预警
	router.Get("/contracts/:id/risks", middleware.AuthMiddleware("contract:list_risks"), h.ListContractRisks)
	router.Put("/contract-risks/:id/status", middleware.AuthMiddleware("contract:update_risk_status"), h.UpdateRiskStatus)
}

// CreateContractTemplate 创建合同模板
func (h *ContractHandler) CreateContractTemplate(c fiber.Ctx) error {
	var req struct {
		TemplateName    string `json:"template_name" validate:"required"`
		TemplateContent string `json:"template_content" validate:"required"`
	}
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	template := ContractTemplate{
		TemplateName:    req.TemplateName,
		TemplateContent: req.TemplateContent,
		Status:          1, // 启用
		CreateTime:      time.Now(),
		UpdateTime:      time.Now(),
	}

	if err := h.db.Create(&template).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create contract template"})
	}

	return c.Status(fiber.StatusCreated).JSON(template)
}

// ListContractTemplates 列出合同模板
func (h *ContractHandler) ListContractTemplates(c fiber.Ctx) error {
	var templates []ContractTemplate
	if err := h.db.Where("status = ?", 1).Find(&templates).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to list contract templates"})
	}

	return c.JSON(templates)
}

// GetContractTemplate 获取合同模板
func (h *ContractHandler) GetContractTemplate(c fiber.Ctx) error {
	id := c.Params("id")
	var template ContractTemplate
	if err := h.db.First(&template, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Contract template not found"})
	}

	return c.JSON(template)
}

// UpdateContractTemplate 更新合同模板
func (h *ContractHandler) UpdateContractTemplate(c fiber.Ctx) error {
	id := c.Params("id")
	var template ContractTemplate
	if err := h.db.First(&template, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Contract template not found"})
	}

	var req struct {
		TemplateName    string `json:"template_name"`
		TemplateContent string `json:"template_content"`
		Status          uint8  `json:"status"`
	}
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if req.TemplateName != "" {
		template.TemplateName = req.TemplateName
	}
	if req.TemplateContent != "" {
		template.TemplateContent = req.TemplateContent
	}
	template.Status = req.Status
	template.UpdateTime = time.Now()

	if err := h.db.Save(&template).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update contract template"})
	}

	return c.JSON(template)
}

// DeleteContractTemplate 删除合同模板
func (h *ContractHandler) DeleteContractTemplate(c fiber.Ctx) error {
	id := c.Params("id")
	var template ContractTemplate
	if err := h.db.First(&template, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Contract template not found"})
	}

	if err := h.db.Delete(&template).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete contract template"})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}

// CreateContract 创建合同
func (h *ContractHandler) CreateContract(c fiber.Ctx) error {
	var req CreateContractRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// 生成合同编号
	contractCode := "CON" + time.Now().Format("20060102150405")

	// 创建合同
	contract := Contract{
		ContractCode:  contractCode,
		PropertyID:    req.PropertyID,
		TenantID:      req.TenantID,
		TemplateID:    req.TemplateID,
		ContractName:  req.ContractName,
		StartDate:     req.StartDate,
		EndDate:       req.EndDate,
		RentAmount:    req.RentAmount,
		DepositAmount: req.DepositAmount,
		PaymentMethod: req.PaymentMethod,
		Status:        0, // 待签署
		CreatorID:     1, // 暂时硬编码，实际应从会话中获取
		CreateTime:    time.Now(),
		UpdateTime:    time.Now(),
	}

	// 开始事务
	tx := h.db.Begin()
	if err := tx.Create(&contract).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create contract"})
	}

	// 处理条款
	for _, clause := range req.Clauses {
		contractClause := ContractClause{
			ContractID:    contract.ID,
			ClauseType:    clause.ClauseType,
			ClauseContent: clause.ClauseContent,
			CreateTime:    time.Now(),
		}
		if err := tx.Create(&contractClause).Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create contract clause"})
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to commit transaction"})
	}

	return c.Status(fiber.StatusCreated).JSON(contract)
}

// ListContracts 列出合同
func (h *ContractHandler) ListContracts(c fiber.Ctx) error {
	var req ContractQueryRequest
	if err := c.Bind().Query(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// 设置默认值
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	// 构建查询
	query := h.db.Model(&Contract{})

	// 应用过滤条件
	if req.PropertyID != nil {
		query = query.Where("property_id = ?", *req.PropertyID)
	}
	if req.TenantID != nil {
		query = query.Where("tenant_id = ?", *req.TenantID)
	}
	if req.Status != nil {
		query = query.Where("status = ?", *req.Status)
	}
	if req.StartDateMin != nil {
		query = query.Where("start_date >= ?", *req.StartDateMin)
	}
	if req.StartDateMax != nil {
		query = query.Where("start_date <= ?", *req.StartDateMax)
	}

	// 计算总数
	var total int64
	query.Count(&total)

	// 分页
	offset := (req.Page - 1) * req.PageSize
	var contracts []Contract
	if err := query.Offset(offset).Limit(req.PageSize).Order("create_time DESC").Find(&contracts).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to list contracts"})
	}

	return c.JSON(fiber.Map{
		"total":     total,
		"page":      req.Page,
		"page_size": req.PageSize,
		"contracts": contracts,
	})
}

// GetContract 获取合同详情
func (h *ContractHandler) GetContract(c fiber.Ctx) error {
	id := c.Params("id")
	var contract Contract
	if err := h.db.First(&contract, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Contract not found"})
	}

	// 获取条款
	var clauses []ContractClause
	h.db.Where("contract_id = ?", contract.ID).Find(&clauses)

	// 获取签章
	var signatures []ContractSignature
	h.db.Where("contract_id = ?", contract.ID).Find(&signatures)

	// 获取状态变更记录
	var statusLogs []ContractStatusLog
	h.db.Where("contract_id = ?", contract.ID).Order("create_time DESC").Find(&statusLogs)

	// 获取风险预警
	var risks []ContractRisk
	h.db.Where("contract_id = ?", contract.ID).Find(&risks)

	return c.JSON(fiber.Map{
		"contract":    contract,
		"clauses":     clauses,
		"signatures":  signatures,
		"status_logs": statusLogs,
		"risks":       risks,
	})
}

// UpdateContract 更新合同
func (h *ContractHandler) UpdateContract(c fiber.Ctx) error {
	id := c.Params("id")
	var contract Contract
	if err := h.db.First(&contract, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Contract not found"})
	}

	// 检查状态（已签署或履行中的合同不允许修改）
	if contract.Status >= 1 {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Cannot update contract in current status"})
	}

	var req UpdateContractRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// 更新字段
	if req.ContractName != "" {
		contract.ContractName = req.ContractName
	}
	if !req.StartDate.IsZero() {
		contract.StartDate = req.StartDate
	}
	if !req.EndDate.IsZero() {
		contract.EndDate = req.EndDate
	}
	if req.RentAmount > 0 {
		contract.RentAmount = req.RentAmount
	}
	if req.DepositAmount > 0 {
		contract.DepositAmount = req.DepositAmount
	}
	contract.PaymentMethod = req.PaymentMethod
	contract.UpdateTime = time.Now()

	// 开始事务
	tx := h.db.Begin()

	// 保存合同信息
	if err := tx.Save(&contract).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update contract"})
	}

	// 更新条款
	if len(req.Clauses) > 0 {
		// 删除旧条款
		if err := tx.Where("contract_id = ?", contract.ID).Delete(&ContractClause{}).Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete old clauses"})
		}

		// 添加新条款
		for _, clause := range req.Clauses {
			contractClause := ContractClause{
				ContractID:    contract.ID,
				ClauseType:    clause.ClauseType,
				ClauseContent: clause.ClauseContent,
				CreateTime:    time.Now(),
			}
			if err := tx.Create(&contractClause).Error; err != nil {
				tx.Rollback()
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create contract clause"})
			}
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to commit transaction"})
	}

	return c.JSON(contract)
}

// DeleteContract 删除合同
func (h *ContractHandler) DeleteContract(c fiber.Ctx) error {
	id := c.Params("id")
	var contract Contract
	if err := h.db.First(&contract, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Contract not found"})
	}

	// 检查状态（仅待签署状态的合同允许删除）
	if contract.Status != 0 {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Cannot delete contract in current status"})
	}

	// 开始事务
	tx := h.db.Begin()

	// 删除相关数据
	if err := tx.Where("contract_id = ?", contract.ID).Delete(&ContractClause{}).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete contract clauses"})
	}

	if err := tx.Where("contract_id = ?", contract.ID).Delete(&ContractSignature{}).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete contract signatures"})
	}

	if err := tx.Where("contract_id = ?", contract.ID).Delete(&ContractStatusLog{}).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete contract status logs"})
	}

	if err := tx.Where("contract_id = ?", contract.ID).Delete(&ContractRisk{}).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete contract risks"})
	}

	// 删除合同
	if err := tx.Delete(&contract).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete contract"})
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to commit transaction"})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}

// ChangeContractStatus 变更合同状态
func (h *ContractHandler) ChangeContractStatus(c fiber.Ctx) error {
	id := c.Params("id")
	var contract Contract
	if err := h.db.First(&contract, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Contract not found"})
	}

	var req ChangeContractStatusRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// 记录旧状态
	oldStatus := contract.Status

	// 更新状态
	contract.Status = req.NewStatus
	contract.UpdateTime = time.Now()

	// 开始事务
	tx := h.db.Begin()

	// 保存合同状态
	if err := tx.Save(&contract).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update contract status"})
	}

	// 记录状态变更日志
	statusLog := ContractStatusLog{
		ContractID:   contract.ID,
		OldStatus:    oldStatus,
		NewStatus:    req.NewStatus,
		ChangeReason: req.ChangeReason,
		OperatorID:   1, // 暂时硬编码，实际应从会话中获取
		CreateTime:   time.Now(),
	}
	if err := tx.Create(&statusLog).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create status log"})
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to commit transaction"})
	}

	return c.JSON(contract)
}

// CreateSignature 创建签章
func (h *ContractHandler) CreateSignature(c fiber.Ctx) error {
	contractID := c.Params("id")
	var contract Contract
	if err := h.db.First(&contract, contractID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Contract not found"})
	}

	var req CreateSignatureRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	signature := ContractSignature{
		ContractID:   contract.ID,
		SignerType:   req.SignerType,
		SignerID:     req.SignerID,
		SignerName:   req.SignerName,
		SignatureURL: req.SignatureURL,
		SignTime:     time.Now(),
	}

	if err := h.db.Create(&signature).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create signature"})
	}

	// 检查是否所有必要的签章都已完成
	// 这里可以根据业务逻辑判断，例如需要房东和租户都签章后，合同状态变为已签署

	return c.Status(fiber.StatusCreated).JSON(signature)
}

// ListSignatures 列出签章
func (h *ContractHandler) ListSignatures(c fiber.Ctx) error {
	contractID := c.Params("id")
	var contract Contract
	if err := h.db.First(&contract, contractID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Contract not found"})
	}

	var signatures []ContractSignature
	if err := h.db.Where("contract_id = ?", contract.ID).Find(&signatures).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to list signatures"})
	}

	return c.JSON(signatures)
}

// ListContractRisks 列出合同风险
func (h *ContractHandler) ListContractRisks(c fiber.Ctx) error {
	contractID := c.Params("id")
	var contract Contract
	if err := h.db.First(&contract, contractID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Contract not found"})
	}

	var risks []ContractRisk
	if err := h.db.Where("contract_id = ?", contract.ID).Find(&risks).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to list contract risks"})
	}

	return c.JSON(risks)
}

// UpdateRiskStatus 更新风险状态
func (h *ContractHandler) UpdateRiskStatus(c fiber.Ctx) error {
	riskID := c.Params("id")
	var risk ContractRisk
	if err := h.db.First(&risk, riskID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Risk not found"})
	}

	var req struct {
		Status uint8 `json:"status" validate:"required"`
	}
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	risk.Status = req.Status
	risk.UpdateTime = time.Now()

	if err := h.db.Save(&risk).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update risk status"})
	}

	return c.JSON(risk)
}
