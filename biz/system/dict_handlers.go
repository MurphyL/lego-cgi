package system

import (
	"fmt"

	"github.com/gofiber/fiber/v3"

	"murphyl.com/lego/cgi"
	"murphyl.com/lego/pkg/requests"
	"murphyl.com/lego/pkg/sugar"
)

var sugarLogger = sugar.NewSugarLogger()

func SearchDictTypeHandler(c fiber.Ctx) error {
	dao := cgi.DefaultDataAccessLayer(c)
	records := make([]DictType, 0)
	if err := dao.RetrieveAll(&records); err == nil {
		return c.JSON(requests.NewSuccessResult(records))
	} else {
		sugarLogger.Error("查询数据字典类型列表出错：", err.Error())
		wraped := requests.NewResultViaError(fmt.Errorf("查询数据字典类型列表出错：%s", err.Error()))
		return c.Status(fiber.StatusInternalServerError).JSON(wraped)
	}
}
