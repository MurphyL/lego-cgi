package handlers

import (
	"github.com/gofiber/fiber/v3"

	"murphyl.com/lego/dal"
	"murphyl.com/lego/misc"
)

var sugarLogger = misc.NewSugarLogger()

type DictType struct {
	DictCode string
	DictName string
}

type DictItem struct {
	DictCode  string
	ItemLabel string
	ItemValue string
}

type DictGroup struct {
	Items []DictItem
}

func GetDictTypeHandler(c fiber.Ctx) error {
	dao, _ := fiber.GetService[dal.DataAccessLayer](c.App().State(), "default")
	sugarLogger.Infoln("获取字典类型：", dao)
	return c.JSON(fiber.Map{"code": "200", "data": []DictType{}})
}
