package system

import (
	"github.com/gofiber/fiber/v3"

	"murphyl.com/lego/cgi/support"
	"murphyl.com/lego/udf"
)

var sugarLogger = udf.NewSugarLogger()

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
	dao := support.DefaultDataAccessLayer(c)
	sugarLogger.Infoln("获取字典类型：", dao)
	return c.JSON(fiber.Map{"code": "200", "data": []DictType{}})
}
