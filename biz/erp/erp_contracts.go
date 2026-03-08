package erp

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
	"murphyl.com/lego/cgi"
	"murphyl.com/lego/fns/entry"
	"murphyl.com/lego/fns/sugar"
)

var sugarLogger = sugar.NewSugarLogger()

// ContractHandler 合同处理器
type ContractHandler struct {
	db *gorm.DB
}

// NewContractHandler 创建合同处理器
func NewContractHandler(dao *gorm.DB) func(router fiber.Router) {
	return func(router fiber.Router) {
		sugarLogger.Info("注册合同管理模块")
		h := &ContractHandler{db: dao}
		h.RegisterRoutes(router)
	}
}

// RegisterRoutes 注册路由
func (h *ContractHandler) RegisterRoutes(router fiber.Router) {
	// 合同管理
	router.Get("/contracts/:id", func(c fiber.Ctx) error {
		return cgi.RetrieveOne[Contract, Contract](c, h.db)
	})
}

// Contract 合同基础信息
type Contract struct {
	entry.BaseEntry
	ContractCode    string    `gorm:"size:30;uniqueIndex" json:"contract_code"`
	ContractName    string    `gorm:"size:100" json:"contract_name"`
	StartDate       time.Time `json:"start_date"`
	EndDate         time.Time `json:"end_date"`
	Status          uint8     `json:"status"`
	ContractFileURL string    `gorm:"size:500" json:"contract_file_url"`
}

func (*Contract) TableName() string {
	return "lego_contract"
}
