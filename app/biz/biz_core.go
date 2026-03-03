package biz

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	contract "murphyl.com/lego/app/biz/contract/handlers"
	finance "murphyl.com/lego/app/biz/finance/handlers"
)

// UseFinanceManager 财务管理模块
func UseFinanceManager(db *gorm.DB) func(router fiber.Router) {
	return func(router fiber.Router) {
		financeHandler := finance.NewFinanceHandler(db)
		financeHandler.RegisterRoutes(router)
	}
}

// UseContractManager 合同管理模块
func UseContractManager(db *gorm.DB) func(router fiber.Router) {
	return func(router fiber.Router) {
		contractHandler := contract.NewContractHandler(db)
		contractHandler.RegisterRoutes(router)
	}
}
