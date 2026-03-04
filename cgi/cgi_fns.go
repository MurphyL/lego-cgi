package cgi

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func RetrieveEntries[Q any, T any](c fiber.Ctx, db *gorm.DB) error {
	var query = new(Q)
	if err := c.Bind().All(query); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}
	var records []T
	if err := db.Where(query).Find(&records).Error; err != nil {
		sugarLogger.Error("查询列表出错：", err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(records)
}

func RetrieveEntry[Q any, T any](c fiber.Ctx, db *gorm.DB) error {
	var query = new(Q)
	if err := c.Bind().All(query); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}
	var record T
	if err := db.Where(query).Take(&record).Error; err != nil {
		sugarLogger.Error("查询记录出错：", err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(record)
}

func CreateEntry[T any](c fiber.Ctx, db *gorm.DB) error {
	var payload = new(T)
	if err := c.Bind().Body(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}
	if err := db.Create(&payload).Error; err != nil {
		sugarLogger.Error("创建记录出错：", err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(payload)
}
