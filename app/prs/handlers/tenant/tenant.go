package tenant

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
	"murphyl.com/app/prs/middleware"
)

/**
 * Tenant 租户
 */

// Tenant 租户基础信息
type Tenant struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	TenantCode    string    `gorm:"size:30;uniqueIndex" json:"tenant_code"`
	Name          string    `gorm:"size:50" json:"name"`
	Gender        uint8     `json:"gender"`
	Age           int       `json:"age"`
	IDCard        string    `gorm:"size:18;uniqueIndex" json:"id_card"`
	Phone         string    `gorm:"size:20" json:"phone"`
	Phone2        string    `gorm:"size:20" json:"phone2"`
	Email         string    `gorm:"size:100" json:"email"`
	Occupation    string    `gorm:"size:100" json:"occupation"`
	FamilyMembers int       `json:"family_members"`
	CreditScore   *int      `json:"credit_score"`
	Level         uint8     `json:"level"`
	Status        uint8     `json:"status"`
	CreatorID     uint      `json:"creator_id"`
	CreateTime    time.Time `json:"create_time"`
	UpdateTime    time.Time `json:"update_time"`
}

func (*Tenant) TableName() string {
	return "hrs_tenant"
}

// TenantQualification 租户资质文件
type TenantQualification struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	TenantID     uint      `json:"tenant_id"`
	FileType     uint8     `json:"file_type"`
	FileURL      string    `gorm:"size:500" json:"file_url"`
	VerifyStatus uint8     `json:"verify_status"`
	VerifyResult string    `gorm:"size:255" json:"verify_result"`
	CreateTime   time.Time `json:"create_time"`
}

func (*TenantQualification) TableName() string {
	return "hrs_tenant_qualification"
}

// TenantFollowup 租户跟进记录
type TenantFollowup struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	TenantID   uint       `json:"tenant_id"`
	Stage      uint8      `json:"stage"`
	Content    string     `gorm:"type:text" json:"content"`
	NextAction string     `gorm:"size:255" json:"next_action"`
	NextTime   *time.Time `json:"next_time"`
	AgentID    uint       `json:"agent_id"`
	CreateTime time.Time  `json:"create_time"`
	UpdateTime time.Time  `json:"update_time"`
}

func (*TenantFollowup) TableName() string {
	return "hrs_tenant_followup"
}

// TenantCommunication 租户沟通记录
type TenantCommunication struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	TenantID          uint      `json:"tenant_id"`
	CommunicationType uint8     `json:"communication_type"`
	CommunicationTime time.Time `json:"communication_time"`
	Content           string    `gorm:"type:text" json:"content"`
	Duration          *int      `json:"duration"`
	RecordingURL      string    `gorm:"size:500" json:"recording_url"`
	AgentID           uint      `json:"agent_id"`
	CreateTime        time.Time `json:"create_time"`
}

func (*TenantCommunication) TableName() string {
	return "hrs_tenant_communication"
}

// CreateTenantRequest 创建租户请求
type CreateTenantRequest struct {
	Name           string          `json:"name" validate:"required"`
	Gender         uint8           `json:"gender"`
	Age            int             `json:"age" validate:"required"`
	IDCard         string          `json:"id_card" validate:"required"`
	Phone          string          `json:"phone" validate:"required"`
	Phone2         string          `json:"phone2"`
	Email          string          `json:"email"`
	Occupation     string          `json:"occupation"`
	FamilyMembers  int             `json:"family_members"`
	CreditScore    *int            `json:"credit_score"`
	Qualifications []Qualification `json:"qualifications"`
}

// Qualification 资质文件信息
type Qualification struct {
	FileType uint8  `json:"file_type" validate:"required"`
	FileURL  string `json:"file_url" validate:"required"`
}

// UpdateTenantRequest 更新租户请求
type UpdateTenantRequest struct {
	Name           string          `json:"name"`
	Gender         uint8           `json:"gender"`
	Age            int             `json:"age"`
	Phone          string          `json:"phone"`
	Phone2         string          `json:"phone2"`
	Email          string          `json:"email"`
	Occupation     string          `json:"occupation"`
	FamilyMembers  int             `json:"family_members"`
	CreditScore    *int            `json:"credit_score"`
	Qualifications []Qualification `json:"qualifications"`
}

// TenantQueryRequest 租户查询请求
type TenantQueryRequest struct {
	Name     string `json:"name"`
	IDCard   string `json:"id_card"`
	Phone    string `json:"phone"`
	Level    *uint8 `json:"level"`
	Status   *uint8 `json:"status"`
	Page     int    `json:"page" default:"1"`
	PageSize int    `json:"page_size" default:"10"`
}

// CreateFollowupRequest 创建跟进记录请求
type CreateFollowupRequest struct {
	Stage      uint8      `json:"stage" validate:"required"`
	Content    string     `json:"content" validate:"required"`
	NextAction string     `json:"next_action"`
	NextTime   *time.Time `json:"next_time"`
}

// CreateCommunicationRequest 创建沟通记录请求
type CreateCommunicationRequest struct {
	CommunicationType uint8     `json:"communication_type" validate:"required"`
	CommunicationTime time.Time `json:"communication_time" validate:"required"`
	Content           string    `json:"content" validate:"required"`
	Duration          *int      `json:"duration"`
	RecordingURL      string    `json:"recording_url"`
}

// TenantHandler 租户处理器
type TenantHandler struct {
	db *gorm.DB
}

// NewTenantHandler 创建租户处理器
func NewTenantHandler(dao *gorm.DB) func(router fiber.Router) {
	return func(router fiber.Router) {
		h := &TenantHandler{db: dao}
		h.RegisterRoutes(router)
	}
}

// RegisterRoutes 注册路由
func (h *TenantHandler) RegisterRoutes(router fiber.Router) {
	router.Post("/tenants", middleware.AuthMiddleware("tenant:create"), h.CreateTenant)
	router.Get("/tenants", middleware.AuthMiddleware("tenant:list"), h.ListTenants)
	router.Get("/tenants/:id", middleware.AuthMiddleware("tenant:view"), h.GetTenant)
	router.Put("/tenants/:id", middleware.AuthMiddleware("tenant:update"), h.UpdateTenant)
	router.Delete("/tenants/:id", middleware.AuthMiddleware("tenant:delete"), h.DeleteTenant)
	router.Post("/tenants/:id/followups", middleware.AuthMiddleware("tenant:create_followup"), h.CreateFollowup)
	router.Get("/tenants/:id/followups", middleware.AuthMiddleware("tenant:list_followups"), h.ListFollowups)
	router.Post("/tenants/:id/communications", middleware.AuthMiddleware("tenant:create_communication"), h.CreateCommunication)
	router.Get("/tenants/:id/communications", middleware.AuthMiddleware("tenant:list_communications"), h.ListCommunications)
}

// CreateTenant 创建租户
func (h *TenantHandler) CreateTenant(c fiber.Ctx) error {
	var req CreateTenantRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// 生成租户编号
	tenantCode := "TEN" + time.Now().Format("20060102150405")

	// 计算客户等级
	level := uint8(0) // 默认为普通
	if req.CreditScore != nil && *req.CreditScore >= 700 {
		level = 1 // VIP
	}

	// 创建租户
	tenant := Tenant{
		TenantCode:    tenantCode,
		Name:          req.Name,
		Gender:        req.Gender,
		Age:           req.Age,
		IDCard:        req.IDCard,
		Phone:         req.Phone,
		Phone2:        req.Phone2,
		Email:         req.Email,
		Occupation:    req.Occupation,
		FamilyMembers: req.FamilyMembers,
		CreditScore:   req.CreditScore,
		Level:         level,
		Status:        1, // 正常
		CreatorID:     1, // 暂时硬编码，实际应从会话中获取
		CreateTime:    time.Now(),
		UpdateTime:    time.Now(),
	}

	// 开始事务
	tx := h.db.Begin()
	if err := tx.Create(&tenant).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create tenant"})
	}

	// 处理资质文件
	for _, q := range req.Qualifications {
		qualification := TenantQualification{
			TenantID:     tenant.ID,
			FileType:     q.FileType,
			FileURL:      q.FileURL,
			VerifyStatus: 0, // 待核验
			CreateTime:   time.Now(),
		}
		if err := tx.Create(&qualification).Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create tenant qualification"})
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to commit transaction"})
	}

	return c.Status(fiber.StatusCreated).JSON(tenant)
}

// ListTenants 列出租户
func (h *TenantHandler) ListTenants(c fiber.Ctx) error {
	var req TenantQueryRequest
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
	query := h.db.Model(&Tenant{})

	// 应用过滤条件
	if req.Name != "" {
		query = query.Where("name LIKE ?", "%"+req.Name+"%")
	}
	if req.IDCard != "" {
		query = query.Where("id_card = ?", req.IDCard)
	}
	if req.Phone != "" {
		query = query.Where("phone = ?", req.Phone)
	}
	if req.Level != nil {
		query = query.Where("level = ?", *req.Level)
	}
	if req.Status != nil {
		query = query.Where("status = ?", *req.Status)
	}

	// 计算总数
	var total int64
	query.Count(&total)

	// 分页
	offset := (req.Page - 1) * req.PageSize
	var tenants []Tenant
	if err := query.Offset(offset).Limit(req.PageSize).Order("create_time DESC").Find(&tenants).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to list tenants"})
	}

	return c.JSON(fiber.Map{
		"total":     total,
		"page":      req.Page,
		"page_size": req.PageSize,
		"tenants":   tenants,
	})
}

// GetTenant 获取租户详情
func (h *TenantHandler) GetTenant(c fiber.Ctx) error {
	id := c.Params("id")
	var tenant Tenant
	if err := h.db.First(&tenant, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tenant not found"})
	}

	// 获取资质文件
	var qualifications []TenantQualification
	h.db.Where("tenant_id = ?", tenant.ID).Find(&qualifications)

	// 获取跟进记录
	var followups []TenantFollowup
	h.db.Where("tenant_id = ?", tenant.ID).Order("create_time DESC").Find(&followups)

	// 获取沟通记录
	var communications []TenantCommunication
	h.db.Where("tenant_id = ?", tenant.ID).Order("communication_time DESC").Find(&communications)

	return c.JSON(fiber.Map{
		"tenant":         tenant,
		"qualifications": qualifications,
		"followups":      followups,
		"communications": communications,
	})
}

// UpdateTenant 更新租户
func (h *TenantHandler) UpdateTenant(c fiber.Ctx) error {
	id := c.Params("id")
	var tenant Tenant
	if err := h.db.First(&tenant, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tenant not found"})
	}

	var req UpdateTenantRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// 更新字段
	if req.Name != "" {
		tenant.Name = req.Name
	}
	tenant.Gender = req.Gender
	if req.Age > 0 {
		tenant.Age = req.Age
	}
	if req.Phone != "" {
		tenant.Phone = req.Phone
	}
	tenant.Phone2 = req.Phone2
	tenant.Email = req.Email
	tenant.Occupation = req.Occupation
	if req.FamilyMembers > 0 {
		tenant.FamilyMembers = req.FamilyMembers
	}
	if req.CreditScore != nil {
		tenant.CreditScore = req.CreditScore
		// 更新客户等级
		if *req.CreditScore >= 700 {
			tenant.Level = 1 // VIP
		} else {
			tenant.Level = 0 // 普通
		}
	}
	tenant.UpdateTime = time.Now()

	// 开始事务
	tx := h.db.Begin()

	// 保存租户信息
	if err := tx.Save(&tenant).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update tenant"})
	}

	// 更新资质文件
	if len(req.Qualifications) > 0 {
		// 删除旧资质文件
		if err := tx.Where("tenant_id = ?", tenant.ID).Delete(&TenantQualification{}).Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete old qualifications"})
		}

		// 添加新资质文件
		for _, q := range req.Qualifications {
			qualification := TenantQualification{
				TenantID:     tenant.ID,
				FileType:     q.FileType,
				FileURL:      q.FileURL,
				VerifyStatus: 0, // 待核验
				CreateTime:   time.Now(),
			}
			if err := tx.Create(&qualification).Error; err != nil {
				tx.Rollback()
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create tenant qualification"})
			}
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to commit transaction"})
	}

	return c.JSON(tenant)
}

// DeleteTenant 删除租户
func (h *TenantHandler) DeleteTenant(c fiber.Ctx) error {
	id := c.Params("id")
	var tenant Tenant
	if err := h.db.First(&tenant, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tenant not found"})
	}

	// 开始事务
	tx := h.db.Begin()

	// 删除相关数据
	if err := tx.Where("tenant_id = ?", tenant.ID).Delete(&TenantQualification{}).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete tenant qualifications"})
	}

	if err := tx.Where("tenant_id = ?", tenant.ID).Delete(&TenantFollowup{}).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete tenant followups"})
	}

	if err := tx.Where("tenant_id = ?", tenant.ID).Delete(&TenantCommunication{}).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete tenant communications"})
	}

	// 删除租户
	if err := tx.Delete(&tenant).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete tenant"})
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to commit transaction"})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}

// CreateFollowup 创建跟进记录
func (h *TenantHandler) CreateFollowup(c fiber.Ctx) error {
	tenantID := c.Params("id")
	var tenant Tenant
	if err := h.db.First(&tenant, tenantID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tenant not found"})
	}

	var req CreateFollowupRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	followup := TenantFollowup{
		TenantID:   tenant.ID,
		Stage:      req.Stage,
		Content:    req.Content,
		NextAction: req.NextAction,
		NextTime:   req.NextTime,
		AgentID:    1, // 暂时硬编码，实际应从会话中获取
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	if err := h.db.Create(&followup).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create followup record"})
	}

	return c.Status(fiber.StatusCreated).JSON(followup)
}

// ListFollowups 列出跟进记录
func (h *TenantHandler) ListFollowups(c fiber.Ctx) error {
	tenantID := c.Params("id")
	var tenant Tenant
	if err := h.db.First(&tenant, tenantID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tenant not found"})
	}

	var followups []TenantFollowup
	if err := h.db.Where("tenant_id = ?", tenant.ID).Order("create_time DESC").Find(&followups).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to list followup records"})
	}

	return c.JSON(followups)
}

// CreateCommunication 创建沟通记录
func (h *TenantHandler) CreateCommunication(c fiber.Ctx) error {
	tenantID := c.Params("id")
	var tenant Tenant
	if err := h.db.First(&tenant, tenantID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tenant not found"})
	}

	var req CreateCommunicationRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	communication := TenantCommunication{
		TenantID:          tenant.ID,
		CommunicationType: req.CommunicationType,
		CommunicationTime: req.CommunicationTime,
		Content:           req.Content,
		Duration:          req.Duration,
		RecordingURL:      req.RecordingURL,
		AgentID:           1, // 暂时硬编码，实际应从会话中获取
		CreateTime:        time.Now(),
	}

	if err := h.db.Create(&communication).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create communication record"})
	}

	return c.Status(fiber.StatusCreated).JSON(communication)
}

// ListCommunications 列出沟通记录
func (h *TenantHandler) ListCommunications(c fiber.Ctx) error {
	tenantID := c.Params("id")
	var tenant Tenant
	if err := h.db.First(&tenant, tenantID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tenant not found"})
	}

	var communications []TenantCommunication
	if err := h.db.Where("tenant_id = ?", tenant.ID).Order("communication_time DESC").Find(&communications).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to list communication records"})
	}

	return c.JSON(communications)
}
