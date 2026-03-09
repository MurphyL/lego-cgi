package etc

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
	"murphyl.com/lego/cgi"
	"murphyl.com/lego/fns/entry"
)

// 系统管理模块
func NewSystemHandler(dao *gorm.DB) *systemHandler {
	return &systemHandler{dao: dao}
}

type systemHandler struct {
	dao *gorm.DB
}

func (h *systemHandler) RegisterRoutes(router fiber.Router) {
	// 删除数据字典类型
	router.Delete("/dict/types/:dictCode", func(c fiber.Ctx) error {
		return cgi.DeleteOne[struct{ DictCode string }, DictType](c, h.dao)
	})
	// 获取数据字典类型详情
	router.Get("/dict/types/:dictCode", func(c fiber.Ctx) error {
		return cgi.RetrieveOne[struct{ DictCode string }, DictType](c, h.dao)
	})
	// 查询数据字典类型列表
	router.Get("/dict/types", func(c fiber.Ctx) error {
		return cgi.RetrieveAll[struct{ DictCode string }, DictType](c, h.dao)
	})
	// 查询数据字典项
	router.Get("/dict/items", func(c fiber.Ctx) error {
		return cgi.RetrieveAll[struct{ DictCode, ItemValue string }, DictItem](c, h.dao)
	})
	// 获取数据字典项详情
	router.Get("/dict/items/:id", func(c fiber.Ctx) error {
		return cgi.RetrieveOne[struct{ Id string }, DictItem](c, h.dao)
	})
	// 创建数据字典项
	router.Post("/dict/items", func(c fiber.Ctx) error {
		return cgi.CreateOne[DictItem](c, h.dao)
	})
}

// 数据字典类型
type DictType struct {
	entry.BaseEntry
	DictCode    string `json:"dictCode"`
	DictName    string `json:"dictName"`
	Description string `json:"description"`
	Sort        int    `json:"sort"`
}

// 数据字典项
type DictItem struct {
	entry.BaseEntry
	DictCode  string `json:"dictCode"`
	ItemLabel string `json:"itemLabel"`
	ItemValue string `json:"itemValue"`
	Remark    string `json:"remark"`
	Sort      int    `json:"sort"`
}

// 数据字典组
type DictGroup struct {
	DictCode string     `json:"dictCode"`
	DictName string     `json:"dictName"`
	Items    []DictItem `json:"items"`
}

func (i DictItem) TableName() string {
	return "sys_dict_item"
}

func (t DictType) TableName() string {
	return "sys_dict_type"
}
