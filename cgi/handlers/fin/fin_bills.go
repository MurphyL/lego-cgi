package fin

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
	"murphyl.com/lego/fns/entry"
)

/* 账单 */
// BillTypeEnum 账单类型
type BillTypeEnum string

const (
	TypeIncome BillTypeEnum = "income" // 收入
	TypeExpend BillTypeEnum = "expend" // 支出
)

// NewFinanceHandler 创建财务管理处理器
func NewFinanceHandler(dao *gorm.DB) *financeHandler {
	return &financeHandler{db: dao}
}

// financeHandler 财务管理处理器
type financeHandler struct {
	db *gorm.DB
}

// RegisterRoutes 注册路由
func (h *financeHandler) RegisterRoutes(router fiber.Router) {
	// 账单管理
}

// Bill 账单信息
type Bill struct {
	entry.BaseEntry
	BillCode    string    `gorm:"size:30;uniqueIndex" json:"bill_code"`
	BillType    uint8     `json:"bill_type"`
	Amount      float64   `json:"amount"`
	DueDate     time.Time `json:"due_date"`
	PaidAmount  float64   `json:"paid_amount"`
	Description string    `gorm:"size:255" json:"description"`
}

func (*Bill) TableName() string {
	return "fin_bill"
}
