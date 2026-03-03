package system

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v3"

	"murphyl.com/lego/cgi"
	"murphyl.com/lego/pkg/requests"
	"murphyl.com/lego/pkg/sugar"
)

var sugarLogger = sugar.NewSugarLogger()

// CreateDictTypeHandler 创建字典类型
func CreateDictTypeHandler(c fiber.Ctx) error {
	var req DictTypeRequest
	body := c.Body()
	if err := json.Unmarshal(body, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// 实际应用中应该调用服务层创建字典类型
	// dictType, err := dictService.CreateDictType(req)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	// 模拟创建字典类型
	dictType := DictType{
		DictCode:    req.DictCode,
		DictName:    req.DictName,
		Description: req.Description,
		Status:      Status(req.Status),
		Sort:        req.Sort,
	}

	return c.Status(fiber.StatusOK).JSON(dictType)
}

// UpdateDictTypeHandler 更新字典类型
func UpdateDictTypeHandler(c fiber.Ctx) error {
	var req DictTypeRequest
	body := c.Body()
	if err := json.Unmarshal(body, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// 实际应用中应该调用服务层更新字典类型
	// dictType, err := dictService.UpdateDictType(req)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	// 模拟更新字典类型
	dictType := DictType{
		DictCode:    req.DictCode,
		DictName:    req.DictName,
		Description: req.Description,
		Status:      Status(req.Status),
		Sort:        req.Sort,
	}

	return c.Status(fiber.StatusOK).JSON(dictType)
}

// DeleteDictTypeHandler 删除字典类型
func DeleteDictTypeHandler(c fiber.Ctx) error {
	// dictCode := c.Params("dictCode")

	// 实际应用中应该调用服务层删除字典类型
	// err := dictService.DeleteDictType(dictCode)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
}

// GetDictTypeHandler 获取字典类型
func GetDictTypeHandler(c fiber.Ctx) error {
	dictCode := c.Params("dictCode")

	// 实际应用中应该调用服务层获取字典类型
	// dictType, err := dictService.GetDictType(dictCode)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	// 模拟获取字典类型
	dictType := DictType{
		DictCode:    dictCode,
		DictName:    "测试字典",
		Description: "测试字典类型",
		Status:      StatusEnabled,
		Sort:        1,
	}

	return c.Status(fiber.StatusOK).JSON(dictType)
}

// ListDictTypesHandler 列出字典类型
func ListDictTypesHandler(c fiber.Ctx) error {
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

// CreateDictItemHandler 创建字典项
func CreateDictItemHandler(c fiber.Ctx) error {
	var req DictItemRequest
	body := c.Body()
	if err := json.Unmarshal(body, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// 实际应用中应该调用服务层创建字典项
	// dictItem, err := dictService.CreateDictItem(req)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	// 模拟创建字典项
	dictItem := DictItem{
		DictCode:  req.DictCode,
		ItemLabel: req.ItemLabel,
		ItemValue: req.ItemValue,
		Status:    Status(req.Status),
		Sort:      req.Sort,
	}

	return c.Status(fiber.StatusOK).JSON(dictItem)
}

// UpdateDictItemHandler 更新字典项
func UpdateDictItemHandler(c fiber.Ctx) error {
	var req DictItemRequest
	body := c.Body()
	if err := json.Unmarshal(body, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// 实际应用中应该调用服务层更新字典项
	// dictItem, err := dictService.UpdateDictItem(req)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	// 模拟更新字典项
	dictItem := DictItem{
		DictCode:  req.DictCode,
		ItemLabel: req.ItemLabel,
		ItemValue: req.ItemValue,
		Status:    Status(req.Status),
		Sort:      req.Sort,
	}

	return c.Status(fiber.StatusOK).JSON(dictItem)
}

// DeleteDictItemHandler 删除字典项
func DeleteDictItemHandler(c fiber.Ctx) error {
	// id := c.Params("id")

	// 实际应用中应该调用服务层删除字典项
	// err := dictService.DeleteDictItem(id)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
}

// GetDictItemHandler 获取字典项
func GetDictItemHandler(c fiber.Ctx) error {
	// id := c.Params("id")

	// 实际应用中应该调用服务层获取字典项
	// dictItem, err := dictService.GetDictItem(id)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	// 模拟获取字典项
	dictItem := DictItem{
		DictCode:  "test",
		ItemLabel: "测试项",
		ItemValue: "test",
		Status:    StatusEnabled,
		Sort:      1,
	}

	return c.Status(fiber.StatusOK).JSON(dictItem)
}

// ListDictItemsHandler 列出字典项
func ListDictItemsHandler(c fiber.Ctx) error {
	dictCode := c.Query("dictCode")

	// 实际应用中应该调用服务层列出字典项
	// dictItems, err := dictService.ListDictItems(dictCode)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	// 模拟列出字典项
	dictItems := []DictItem{
		{
			DictCode:  dictCode,
			ItemLabel: "选项1",
			ItemValue: "1",
			Status:    StatusEnabled,
			Sort:      1,
		},
		{
			DictCode:  dictCode,
			ItemLabel: "选项2",
			ItemValue: "2",
			Status:    StatusEnabled,
			Sort:      2,
		},
	}

	return c.Status(fiber.StatusOK).JSON(dictItems)
}

// GetDictGroupHandler 获取字典组
func GetDictGroupHandler(c fiber.Ctx) error {
	dictCode := c.Params("dictCode")

	// 实际应用中应该调用服务层获取字典组
	// dictGroup, err := dictService.GetDictGroup(dictCode)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	// 模拟获取字典组
	dictGroup := DictGroup{
		DictCode: dictCode,
		DictName: "测试字典",
		Items: []DictItem{
			{
				DictCode:  dictCode,
				ItemLabel: "选项1",
				ItemValue: "1",
				Status:    StatusEnabled,
				Sort:      1,
			},
			{
				DictCode:  dictCode,
				ItemLabel: "选项2",
				ItemValue: "2",
				Status:    StatusEnabled,
				Sort:      2,
			},
		},
	}

	return c.Status(fiber.StatusOK).JSON(dictGroup)
}
