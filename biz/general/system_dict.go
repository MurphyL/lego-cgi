package general

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
	"murphyl.com/lego/cgi"
	"murphyl.com/lego/dal"
	"murphyl.com/lego/fns/sugar"
)

var sugarLogger = sugar.NewSugarLogger()

// 数据字典模块，主要功能包括：字典类型管理、字典项管理、字典组管理等

type DictType struct {
	dal.BaseEntry
	DictCode    string `json:"dictCode"`
	DictName    string `json:"dictName"`
	Description string `json:"description"`
	Sort        int    `json:"sort"`
}

type DictItem struct {
	dal.BaseEntry
	DictCode  string `json:"dictCode"`
	ItemLabel string `json:"itemLabel"`
	ItemValue string `json:"itemValue"`
	Remark    string `json:"remark"`
	Sort      int    `json:"sort"`
}

type DictGroup struct {
	DictCode string     `json:"dictCode"`
	DictName string     `json:"dictName"`
	Items    []DictItem `json:"items"`
}

// DictTypeRequest 字典类型请求
type DictTypeRequest struct {
	DictCode    string `json:"dictCode"`
	DictName    string `json:"dictName"`
	Description string `json:"description"`
	Status      int    `json:"status"`
	Sort        int    `json:"sort"`
}

func (i DictItem) TableName() string {
	return "sys_dict_item"
}

func (t DictType) TableName() string {
	return "sys_dict_type"
}

type systemDictHandler struct {
	db *gorm.DB
}

func NewSystemDictHandler(dao *gorm.DB) func(router fiber.Router) {
	return func(router fiber.Router) {
		h := &systemDictHandler{db: dao}
		h.RegisterRoutes(router)
	}
}

func (h *systemDictHandler) RegisterRoutes(router fiber.Router) {
	router.Post("/dict/types", h.CreateDictTypeHandler)
	router.Put("/dict/types/:dictCode", h.UpdateDictTypeHandler)
	router.Delete("/dict/types/:dictCode", h.DeleteDictTypeHandler)
	router.Get("/dict/types/:dictCode", h.GetDictTypeHandler)
	// 查询字典类型
	router.Get("/dict/types", func(c fiber.Ctx) error {
		return cgi.RetrieveEntries[struct{ dictCode string }, DictType](c, h.db)
	})
	// 查询字典项
	router.Get("/dict/items", func(c fiber.Ctx) error {
		return cgi.RetrieveEntries[struct{ DictCode, ItemValue string }, DictItem](c, h.db)
	})
	// 查询字典项
	router.Get("/dict/items/:id", func(c fiber.Ctx) error {
		return cgi.RetrieveEntry[struct{ Id string }, DictItem](c, h.db)
	})
	// 创建字典项
	router.Post("/dict/items", func(c fiber.Ctx) error {
		return cgi.CreateEntry[DictItem](c, h.db)
	})
}

// CreateDictTypeHandler 创建字典类型
func (h *systemDictHandler) CreateDictTypeHandler(c fiber.Ctx) error {
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
		Sort:        req.Sort,
	}

	return c.Status(fiber.StatusOK).JSON(dictType)
}

// UpdateDictTypeHandler 更新字典类型
func (h *systemDictHandler) UpdateDictTypeHandler(c fiber.Ctx) error {
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
		Sort:        req.Sort,
	}

	return c.Status(fiber.StatusOK).JSON(dictType)
}

// DeleteDictTypeHandler 删除字典类型
func (h *systemDictHandler) DeleteDictTypeHandler(c fiber.Ctx) error {
	// dictCode := c.Params("dictCode")

	// 实际应用中应该调用服务层删除字典类型
	// err := dictService.DeleteDictType(dictCode)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
}

// GetDictTypeHandler 获取字典类型
func (h *systemDictHandler) GetDictTypeHandler(c fiber.Ctx) error {
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
		Sort:        1,
	}

	return c.Status(fiber.StatusOK).JSON(dictType)
}

// CreateDictItemHandler 创建字典项
func (h *systemDictHandler) CreateDictItemHandler(c fiber.Ctx) error {
	var req DictItem
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
		Sort:      req.Sort,
	}

	return c.Status(fiber.StatusOK).JSON(dictItem)
}

// UpdateDictItemHandler 更新字典项
func (h *systemDictHandler) UpdateDictItemHandler(c fiber.Ctx) error {
	var req DictItem
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
		Sort:      req.Sort,
	}

	return c.Status(fiber.StatusOK).JSON(dictItem)
}

// DeleteDictItemHandler 删除字典项
func (h *systemDictHandler) DeleteDictItemHandler(c fiber.Ctx) error {
	// id := c.Params("id")

	// 实际应用中应该调用服务层删除字典项
	// err := dictService.DeleteDictItem(id)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
}

// GetDictItemHandler 获取字典项
func (h *systemDictHandler) GetDictItemHandler(c fiber.Ctx) error {
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
		Sort:      1,
	}

	return c.Status(fiber.StatusOK).JSON(dictItem)
}

// GetDictGroupHandler 获取字典组
func (h *systemDictHandler) GetDictGroupHandler(c fiber.Ctx) error {
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
				Sort:      1,
			},
			{
				DictCode:  dictCode,
				ItemLabel: "选项2",
				ItemValue: "2",
				Sort:      2,
			},
		},
	}

	return c.Status(fiber.StatusOK).JSON(dictGroup)
}
