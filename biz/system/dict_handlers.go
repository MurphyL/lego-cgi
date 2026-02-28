package system

import (
	"github.com/gofiber/fiber/v3"

	"murphyl.com/lego/cgi"
	"murphyl.com/lego/udf"
)

var sugarLogger = udf.NewSugarLogger()

func SearchDictTypeHandler(c fiber.Ctx) error {
	dao := cgi.DefaultDataAccessLayer(c)
	records := make([]DictType, 0)
	if err := dao.RetrieveAll(&records); err == nil {
		return c.JSON(fiber.Map{"code": "200", "data": records})
	} else {
		sugarLogger.Error("查询数据字典类型列表出错：", err.Error())
		return c.JSON(fiber.Map{"code": "500", "message": err.Error()})
	}
}
