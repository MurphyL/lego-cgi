package cate

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v3"
	"murphyl.com/lego/dal"
)

// cate 模块是分类管理模块，用于管理各种类型的分类
// 主要功能包括：分类的创建、更新、删除、查询、树形结构展示等

// Category 分类定义
type Category struct {
	ID          uint64         `json:"id"`
	ParentID    *uint64        `json:"parentId"` // 父分类ID，根分类为nil
	Code        string         `json:"code"`
	Name        string         `json:"name"`
	Description string         `json:"description,omitempty"`
	Level       int            `json:"level"`  // 分类层级
	Path        string         `json:"path"`   // 分类路径，如：1/2/3
	Weight      int            `json:"weight"` // 排序权重
	Status      dal.StatusEnum `json:"status"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
}

// CategoryWithChildren 带子分类的分类定义
type CategoryWithChildren struct {
	Category
	Children []CategoryWithChildren `json:"children,omitempty"`
}

// IsValid 检查分类是否有效
func (c *Category) IsValid() bool {
	return c.Status.IsEnabled()
}

// CategoryRequest 分类请求
type CategoryRequest struct {
	ParentID    *uint64 `json:"parentId"`
	Code        string  `json:"code"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Weight      int     `json:"weight"`
	Status      int     `json:"status"`
}

// CreateCategoryHandler 创建分类
func CreateCategoryHandler(c fiber.Ctx) error {
	var req CategoryRequest
	body := c.Body()
	if err := json.Unmarshal(body, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// 实际应用中应该调用服务层创建分类
	// category, err := categoryService.CreateCategory(req)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	// 模拟分类创建
	category := Category{
		ID:          1,
		ParentID:    req.ParentID,
		Code:        req.Code,
		Name:        req.Name,
		Description: req.Description,
		Level:       1,
		Path:        "1",
		Weight:      req.Weight,
		Status:      dal.StatusEnum(req.Status),
	}

	return c.Status(fiber.StatusOK).JSON(category)
}

// UpdateCategoryHandler 更新分类
func UpdateCategoryHandler(c fiber.Ctx) error {
	var req CategoryRequest
	body := c.Body()
	if err := json.Unmarshal(body, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// 实际应用中应该调用服务层更新分类
	// category, err := categoryService.UpdateCategory(req)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	// 模拟分类更新
	category := Category{
		ID:          1,
		ParentID:    req.ParentID,
		Code:        req.Code,
		Name:        req.Name,
		Description: req.Description,
		Level:       1,
		Path:        "1",
		Weight:      req.Weight,
		Status:      dal.StatusEnum(req.Status),
	}

	return c.Status(fiber.StatusOK).JSON(category)
}

// DeleteCategoryHandler 删除分类
func DeleteCategoryHandler(c fiber.Ctx) error {
	// id := c.Params("id")

	// 实际应用中应该调用服务层删除分类
	// err := categoryService.DeleteCategory(id)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
}

// GetCategoryHandler 获取分类
func GetCategoryHandler(c fiber.Ctx) error {
	// id := c.Params("id")

	// 实际应用中应该调用服务层获取分类
	// category, err := categoryService.GetCategory(id)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	// 模拟分类获取
	category := Category{
		ID:          1,
		ParentID:    nil,
		Code:        "root",
		Name:        "根分类",
		Description: "系统根分类",
		Level:       1,
		Path:        "1",
		Weight:      100,
		Status:      dal.StatusEnabled,
	}

	return c.Status(fiber.StatusOK).JSON(category)
}

// ListCategoriesHandler 列出分类
func ListCategoriesHandler(c fiber.Ctx) error {
	// 实际应用中应该调用服务层列出分类
	// categories, err := categoryService.ListCategories()
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	// 模拟分类列表
	categories := []Category{
		{
			ID:          1,
			ParentID:    nil,
			Code:        "root",
			Name:        "根分类",
			Description: "系统根分类",
			Level:       1,
			Path:        "1",
			Weight:      100,
			Status:      dal.StatusEnabled,
		},
		{
			ID:          2,
			ParentID:    uintPtr(1),
			Code:        "electronics",
			Name:        "电子产品",
			Description: "电子产品分类",
			Level:       2,
			Path:        "1/2",
			Weight:      90,
			Status:      dal.StatusEnabled,
		},
		{
			ID:          3,
			ParentID:    uintPtr(2),
			Code:        "phones",
			Name:        "手机",
			Description: "手机分类",
			Level:       3,
			Path:        "1/2/3",
			Weight:      80,
			Status:      dal.StatusEnabled,
		},
	}

	return c.Status(fiber.StatusOK).JSON(categories)
}

// GetCategoryTreeHandler 获取分类树
func GetCategoryTreeHandler(c fiber.Ctx) error {
	// 实际应用中应该调用服务层获取分类树
	// categoryTree, err := categoryService.GetCategoryTree()
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	// 模拟分类树
	categoryTree := CategoryWithChildren{
		Category: Category{
			ID:          1,
			ParentID:    nil,
			Code:        "root",
			Name:        "根分类",
			Description: "系统根分类",
			Level:       1,
			Path:        "1",
			Weight:      100,
			Status:      dal.StatusEnabled,
		},
		Children: []CategoryWithChildren{
			{
				Category: Category{
					ID:          2,
					ParentID:    uintPtr(1),
					Code:        "electronics",
					Name:        "电子产品",
					Description: "电子产品分类",
					Level:       2,
					Path:        "1/2",
					Weight:      90,
					Status:      dal.StatusEnabled,
				},
				Children: []CategoryWithChildren{
					{
						Category: Category{
							ID:          3,
							ParentID:    uintPtr(2),
							Code:        "phones",
							Name:        "手机",
							Description: "手机分类",
							Level:       3,
							Path:        "1/2/3",
							Weight:      80,
							Status:      dal.StatusEnabled,
						},
					},
					{
						Category: Category{
							ID:          4,
							ParentID:    uintPtr(2),
							Code:        "laptops",
							Name:        "笔记本电脑",
							Description: "笔记本电脑分类",
							Level:       3,
							Path:        "1/2/4",
							Weight:      70,
							Status:      dal.StatusEnabled,
						},
					},
				},
			},
			{
				Category: Category{
					ID:          5,
					ParentID:    uintPtr(1),
					Code:        "clothing",
					Name:        "服装",
					Description: "服装分类",
					Level:       2,
					Path:        "1/5",
					Weight:      85,
					Status:      dal.StatusEnabled,
				},
			},
		},
	}

	return c.Status(fiber.StatusOK).JSON(categoryTree)
}

// uintPtr 返回uint64的指针
func uintPtr(u uint64) *uint64 {
	return &u
}
