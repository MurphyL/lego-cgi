package finance

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
	"murphyl.com/lego/app/hrs/middleware"
)

/**
 * Finance 财务管理
 */

// Bill 账单信息
type Bill struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	BillCode    string    `gorm:"size:30;uniqueIndex" json:"bill_code"`
	ContractID  uint      `json:"contract_id"`
	TenantID    uint      `json:"tenant_id"`
	PropertyID  uint      `json:"property_id"`
	BillType    uint8     `json:"bill_type"`
	Amount      float64   `json:"amount"`
	DueDate     time.Time `json:"due_date"`
	Status      uint8     `json:"status"`
	PaidAmount  float64   `json:"paid_amount"`
	Description string    `gorm:"size:255" json:"description"`
	CreateTime  time.Time `json:"create_time"`
	UpdateTime  time.Time `json:"update_time"`
}

func (*Bill) TableName() string {
	return "hrs_bill"
}

// Payment 支付记录
type Payment struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	PaymentCode   string    `gorm:"size:30;uniqueIndex" json:"payment_code"`
	BillID        uint      `json:"bill_id"`
	TenantID      uint      `json:"tenant_id"`
	Amount        float64   `json:"amount"`
	PaymentMethod uint8     `json:"payment_method"`
	PaymentTime   time.Time `json:"payment_time"`
	TransactionID string    `gorm:"size:100" json:"transaction_id"`
	Status        uint8     `json:"status"`
	CreateTime    time.Time `json:"create_time"`
}

func (*Payment) TableName() string {
	return "hrs_payment"
}

// FinancialReport 财务报表
type FinancialReport struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	ReportName   string    `gorm:"size:100" json:"report_name"`
	ReportType   uint8     `json:"report_type"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	TotalIncome  float64   `json:"total_income"`
	TotalExpense float64   `json:"total_expense"`
	NetIncome    float64   `json:"net_income"`
	ReportURL    string    `gorm:"size:500" json:"report_url"`
	CreatorID    uint      `json:"creator_id"`
	CreateTime   time.Time `json:"create_time"`
}

func (*FinancialReport) TableName() string {
	return "hrs_financial_report"
}

// CreateBillRequest 创建账单请求
type CreateBillRequest struct {
	ContractID  uint      `json:"contract_id" validate:"required"`
	TenantID    uint      `json:"tenant_id" validate:"required"`
	PropertyID  uint      `json:"property_id" validate:"required"`
	BillType    uint8     `json:"bill_type" validate:"required"`
	Amount      float64   `json:"amount" validate:"required"`
	DueDate     time.Time `json:"due_date" validate:"required"`
	Description string    `json:"description"`
}

// UpdateBillRequest 更新账单请求
type UpdateBillRequest struct {
	Amount      float64   `json:"amount"`
	DueDate     time.Time `json:"due_date"`
	Description string    `json:"description"`
}

// CreatePaymentRequest 创建支付记录请求
type CreatePaymentRequest struct {
	BillID        uint      `json:"bill_id" validate:"required"`
	TenantID      uint      `json:"tenant_id" validate:"required"`
	Amount        float64   `json:"amount" validate:"required"`
	PaymentMethod uint8     `json:"payment_method" validate:"required"`
	PaymentTime   time.Time `json:"payment_time" validate:"required"`
	TransactionID string    `json:"transaction_id" validate:"required"`
}

// BillQueryRequest 账单查询请求
type BillQueryRequest struct {
	ContractID *uint      `json:"contract_id"`
	TenantID   *uint      `json:"tenant_id"`
	PropertyID *uint      `json:"property_id"`
	BillType   *uint8     `json:"bill_type"`
	Status     *uint8     `json:"status"`
	DueDateMin *time.Time `json:"due_date_min"`
	DueDateMax *time.Time `json:"due_date_max"`
	Page       int        `json:"page" default:"1"`
	PageSize   int        `json:"page_size" default:"10"`
}

// PaymentQueryRequest 支付记录查询请求
type PaymentQueryRequest struct {
	BillID         *uint      `json:"bill_id"`
	TenantID       *uint      `json:"tenant_id"`
	PaymentMethod  *uint8     `json:"payment_method"`
	PaymentTimeMin *time.Time `json:"payment_time_min"`
	PaymentTimeMax *time.Time `json:"payment_time_max"`
	Page           int        `json:"page" default:"1"`
	PageSize       int        `json:"page_size" default:"10"`
}

// CreateReportRequest 创建财务报表请求
type CreateReportRequest struct {
	ReportName string    `json:"report_name" validate:"required"`
	ReportType uint8     `json:"report_type" validate:"required"`
	StartDate  time.Time `json:"start_date" validate:"required"`
	EndDate    time.Time `json:"end_date" validate:"required"`
}

// FinanceHandler 财务管理处理器
type FinanceHandler struct {
	db *gorm.DB
}

// NewFinanceHandler 创建财务管理处理器
func NewFinanceHandler(db *gorm.DB) *FinanceHandler {
	return &FinanceHandler{db: db}
}

// RegisterRoutes 注册路由
func (h *FinanceHandler) RegisterRoutes(router fiber.Router) {
	// 账单管理
	router.Post("/bills", middleware.AuthMiddleware("finance:create_bill"), h.CreateBill)
	router.Get("/bills", middleware.AuthMiddleware("finance:list_bills"), h.ListBills)
	router.Get("/bills/:id", middleware.AuthMiddleware("finance:view_bill"), h.GetBill)
	router.Put("/bills/:id", middleware.AuthMiddleware("finance:update_bill"), h.UpdateBill)
	router.Delete("/bills/:id", middleware.AuthMiddleware("finance:delete_bill"), h.DeleteBill)

	// 支付管理
	router.Post("/payments", middleware.AuthMiddleware("finance:create_payment"), h.CreatePayment)
	router.Get("/payments", middleware.AuthMiddleware("finance:list_payments"), h.ListPayments)
	router.Get("/payments/:id", middleware.AuthMiddleware("finance:view_payment"), h.GetPayment)

	// 财务报表
	router.Post("/financial-reports", middleware.AuthMiddleware("finance:create_report"), h.CreateFinancialReport)
	router.Get("/financial-reports", middleware.AuthMiddleware("finance:list_reports"), h.ListFinancialReports)
	router.Get("/financial-reports/:id", middleware.AuthMiddleware("finance:view_report"), h.GetFinancialReport)
}

// CreateBill 创建账单
func (h *FinanceHandler) CreateBill(c fiber.Ctx) error {
	var req CreateBillRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// 生成账单编号
	billCode := "BILL" + time.Now().Format("20060102150405")

	// 创建账单
	bill := Bill{
		BillCode:    billCode,
		ContractID:  req.ContractID,
		TenantID:    req.TenantID,
		PropertyID:  req.PropertyID,
		BillType:    req.BillType,
		Amount:      req.Amount,
		DueDate:     req.DueDate,
		Status:      0, // 待支付
		PaidAmount:  0,
		Description: req.Description,
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
	}

	if err := h.db.Create(&bill).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create bill"})
	}

	return c.Status(fiber.StatusCreated).JSON(bill)
}

// ListBills 列出账单
func (h *FinanceHandler) ListBills(c fiber.Ctx) error {
	var req BillQueryRequest
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
	query := h.db.Model(&Bill{})

	// 应用过滤条件
	if req.ContractID != nil {
		query = query.Where("contract_id = ?", *req.ContractID)
	}
	if req.TenantID != nil {
		query = query.Where("tenant_id = ?", *req.TenantID)
	}
	if req.PropertyID != nil {
		query = query.Where("property_id = ?", *req.PropertyID)
	}
	if req.BillType != nil {
		query = query.Where("bill_type = ?", *req.BillType)
	}
	if req.Status != nil {
		query = query.Where("status = ?", *req.Status)
	}
	if req.DueDateMin != nil {
		query = query.Where("due_date >= ?", *req.DueDateMin)
	}
	if req.DueDateMax != nil {
		query = query.Where("due_date <= ?", *req.DueDateMax)
	}

	// 计算总数
	var total int64
	query.Count(&total)

	// 分页
	offset := (req.Page - 1) * req.PageSize
	var bills []Bill
	if err := query.Offset(offset).Limit(req.PageSize).Order("due_date ASC").Find(&bills).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to list bills"})
	}

	return c.JSON(fiber.Map{
		"total":     total,
		"page":      req.Page,
		"page_size": req.PageSize,
		"bills":     bills,
	})
}

// GetBill 获取账单详情
func (h *FinanceHandler) GetBill(c fiber.Ctx) error {
	id := c.Params("id")
	var bill Bill
	if err := h.db.First(&bill, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Bill not found"})
	}

	// 获取支付记录
	var payments []Payment
	h.db.Where("bill_id = ?", bill.ID).Order("payment_time DESC").Find(&payments)

	return c.JSON(fiber.Map{
		"bill":     bill,
		"payments": payments,
	})
}

// UpdateBill 更新账单
func (h *FinanceHandler) UpdateBill(c fiber.Ctx) error {
	id := c.Params("id")
	var bill Bill
	if err := h.db.First(&bill, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Bill not found"})
	}

	// 检查状态（已支付的账单不允许修改）
	if bill.Status == 1 {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Cannot update paid bill"})
	}

	var req UpdateBillRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// 更新字段
	if req.Amount > 0 {
		bill.Amount = req.Amount
	}
	if !req.DueDate.IsZero() {
		bill.DueDate = req.DueDate
	}
	if req.Description != "" {
		bill.Description = req.Description
	}
	bill.UpdateTime = time.Now()

	if err := h.db.Save(&bill).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update bill"})
	}

	return c.JSON(bill)
}

// DeleteBill 删除账单
func (h *FinanceHandler) DeleteBill(c fiber.Ctx) error {
	id := c.Params("id")
	var bill Bill
	if err := h.db.First(&bill, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Bill not found"})
	}

	// 检查状态（已支付的账单不允许删除）
	if bill.Status == 1 {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Cannot delete paid bill"})
	}

	// 开始事务
	tx := h.db.Begin()

	// 删除相关支付记录
	if err := tx.Where("bill_id = ?", bill.ID).Delete(&Payment{}).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete payments"})
	}

	// 删除账单
	if err := tx.Delete(&bill).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete bill"})
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to commit transaction"})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}

// CreatePayment 创建支付记录
func (h *FinanceHandler) CreatePayment(c fiber.Ctx) error {
	var req CreatePaymentRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// 生成支付编号
	paymentCode := "PAY" + time.Now().Format("20060102150405")

	// 检查账单是否存在
	var bill Bill
	if err := h.db.First(&bill, req.BillID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Bill not found"})
	}

	// 开始事务
	tx := h.db.Begin()

	// 创建支付记录
	payment := Payment{
		PaymentCode:   paymentCode,
		BillID:        req.BillID,
		TenantID:      req.TenantID,
		Amount:        req.Amount,
		PaymentMethod: req.PaymentMethod,
		PaymentTime:   req.PaymentTime,
		TransactionID: req.TransactionID,
		Status:        1, // 成功
		CreateTime:    time.Now(),
	}

	if err := tx.Create(&payment).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create payment"})
	}

	// 更新账单状态和已支付金额
	bill.PaidAmount += req.Amount
	if bill.PaidAmount >= bill.Amount {
		bill.Status = 1 // 已支付
		bill.PaidAmount = bill.Amount
	} else {
		bill.Status = 3 // 部分支付
	}
	bill.UpdateTime = time.Now()

	if err := tx.Save(&bill).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update bill"})
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to commit transaction"})
	}

	return c.Status(fiber.StatusCreated).JSON(payment)
}

// ListPayments 列出支付记录
func (h *FinanceHandler) ListPayments(c fiber.Ctx) error {
	var req PaymentQueryRequest
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
	query := h.db.Model(&Payment{})

	// 应用过滤条件
	if req.BillID != nil {
		query = query.Where("bill_id = ?", *req.BillID)
	}
	if req.TenantID != nil {
		query = query.Where("tenant_id = ?", *req.TenantID)
	}
	if req.PaymentMethod != nil {
		query = query.Where("payment_method = ?", *req.PaymentMethod)
	}
	if req.PaymentTimeMin != nil {
		query = query.Where("payment_time >= ?", *req.PaymentTimeMin)
	}
	if req.PaymentTimeMax != nil {
		query = query.Where("payment_time <= ?", *req.PaymentTimeMax)
	}

	// 计算总数
	var total int64
	query.Count(&total)

	// 分页
	offset := (req.Page - 1) * req.PageSize
	var payments []Payment
	if err := query.Offset(offset).Limit(req.PageSize).Order("payment_time DESC").Find(&payments).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to list payments"})
	}

	return c.JSON(fiber.Map{
		"total":     total,
		"page":      req.Page,
		"page_size": req.PageSize,
		"payments":  payments,
	})
}

// GetPayment 获取支付记录详情
func (h *FinanceHandler) GetPayment(c fiber.Ctx) error {
	id := c.Params("id")
	var payment Payment
	if err := h.db.First(&payment, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Payment not found"})
	}

	return c.JSON(payment)
}

// CreateFinancialReport 创建财务报表
func (h *FinanceHandler) CreateFinancialReport(c fiber.Ctx) error {
	var req CreateReportRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// 计算总收入（已支付的账单）
	var totalIncome float64
	h.db.Model(&Bill{}).Where("status = ? AND create_time BETWEEN ? AND ?", 1, req.StartDate, req.EndDate).Select("COALESCE(SUM(amount), 0)").Scan(&totalIncome)

	// 计算总支出（暂时设为0，实际应根据业务逻辑计算）
	totalExpense := 0.0

	// 计算净收入
	netIncome := totalIncome - totalExpense

	// 创建财务报表
	report := FinancialReport{
		ReportName:   req.ReportName,
		ReportType:   req.ReportType,
		StartDate:    req.StartDate,
		EndDate:      req.EndDate,
		TotalIncome:  totalIncome,
		TotalExpense: totalExpense,
		NetIncome:    netIncome,
		CreatorID:    1, // 暂时硬编码，实际应从会话中获取
		CreateTime:   time.Now(),
	}

	if err := h.db.Create(&report).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create financial report"})
	}

	return c.Status(fiber.StatusCreated).JSON(report)
}

// ListFinancialReports 列出财务报表
func (h *FinanceHandler) ListFinancialReports(c fiber.Ctx) error {
	var reports []FinancialReport
	if err := h.db.Order("create_time DESC").Find(&reports).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to list financial reports"})
	}

	return c.JSON(reports)
}

// GetFinancialReport 获取财务报表详情
func (h *FinanceHandler) GetFinancialReport(c fiber.Ctx) error {
	id := c.Params("id")
	var report FinancialReport
	if err := h.db.First(&report, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Financial report not found"})
	}

	return c.JSON(report)
}
