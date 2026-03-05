package tag

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
)

// TagRequest 标签请求
type TagRequest struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        int    `json:"type"`
	Weight      int    `json:"weight"`
	Status      int    `json:"status"`
}

func NewSystemTag(code, name, desc string) *Tag {
	return NewTag(TypeSystem, code, name, desc)
}

func NewTag(tagType Type, code, name, desc string) *Tag {
	return &Tag{Type: tagType, Code: code, Name: name, Description: desc}
}

// CreateTagHandler 创建标签
func CreateTagHandler(c fiber.Ctx) error {
	var req TagRequest
	body := c.Body()
	if err := json.Unmarshal(body, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// 实际应用中应该调用服务层创建标签
	// tag, err := tagService.CreateTag(req)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	// 模拟标签创建
	tag := Tag{
		ID:          1,
		Code:        req.Code,
		Name:        req.Name,
		Description: req.Description,
		Type:        Type(req.Type),
		Weight:      req.Weight,
	}

	return c.Status(fiber.StatusOK).JSON(tag)
}

// UpdateTagHandler 更新标签
func UpdateTagHandler(c fiber.Ctx) error {
	var req TagRequest
	body := c.Body()
	if err := json.Unmarshal(body, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// 实际应用中应该调用服务层更新标签
	// tag, err := tagService.UpdateTag(req)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	// 模拟标签更新
	tag := Tag{
		ID:          1,
		Code:        req.Code,
		Name:        req.Name,
		Description: req.Description,
		Type:        Type(req.Type),
		Weight:      req.Weight,
	}

	return c.Status(fiber.StatusOK).JSON(tag)
}

// DeleteTagHandler 删除标签
func DeleteTagHandler(c fiber.Ctx) error {
	// id := c.Params("id")

	// 实际应用中应该调用服务层删除标签
	// err := tagService.DeleteTag(id)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
}

// GetTagHandler 获取标签
func GetTagHandler(c fiber.Ctx) error {
	// id := c.Params("id")

	// 实际应用中应该调用服务层获取标签
	// tag, err := tagService.GetTag(id)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	// 模拟标签获取
	tag := Tag{
		ID:          1,
		Code:        "system",
		Name:        "系统标签",
		Description: "系统默认标签",
		Type:        TypeSystem,
		Weight:      100,
	}

	return c.Status(fiber.StatusOK).JSON(tag)
}

// ListTagsHandler 列出标签
func ListTagsHandler(c fiber.Ctx) error {
	// 实际应用中应该调用服务层列出标签
	// tags, err := tagService.ListTags()
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	// 模拟标签列表
	tags := []Tag{}

	return c.Status(fiber.StatusOK).JSON(tags)
}
