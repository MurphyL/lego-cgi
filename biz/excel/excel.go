package excel

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

// excel 模块是通用的Excel导出模块，用于导出各种数据到Excel文件
// 主要功能包括：导出数据到Excel、设置Excel样式等

// ExportRequest Excel导出请求
type ExportRequest struct {
	FileName   string        `json:"fileName"`   // 文件名
	SheetName  string        `json:"sheetName"`  // 工作表名称
	Headers    []string      `json:"headers"`    // 表头
	Data       []interface{} `json:"data"`       // 数据
	SheetIndex int           `json:"sheetIndex"` // 工作表索引
}

// ExportResponse Excel导出响应
type ExportResponse struct {
	FileName string `json:"fileName"` // 文件名
	FileUrl  string `json:"fileUrl"`  // 文件URL
	Size     int64  `json:"size"`     // 文件大小
}

func NewExcelHandler(dao *gorm.DB) func(router fiber.Router) {
	return func(router fiber.Router) {
		router.Post("/export", ExportExcelHandler)
	}
}

// ExportExcelHandler 导出Excel文件
func ExportExcelHandler(c fiber.Ctx) error {
	var req ExportRequest
	body := c.Body()
	if err := json.Unmarshal(body, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// 创建Excel文件
	f := excelize.NewFile()

	// 设置工作表名称
	sheetName := req.SheetName
	if sheetName == "" {
		sheetName = "Sheet1"
	}

	// 创建工作表
	index, err := f.NewSheet(sheetName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create sheet"})
	}

	// 设置默认工作表
	f.SetActiveSheet(index)

	// 写入表头
	for i, header := range req.Headers {
		cell := fmt.Sprintf("%c1", 'A'+i)
		f.SetCellValue(sheetName, cell, header)
	}

	// 写入数据
	for i, row := range req.Data {
		// 检查数据类型
		switch v := row.(type) {
		case map[string]interface{}:
			// 处理map类型数据
			for j, header := range req.Headers {
				cell := fmt.Sprintf("%c%d", 'A'+j, i+2)
				if value, ok := v[header]; ok {
					f.SetCellValue(sheetName, cell, value)
				}
			}
		case []interface{}:
			// 处理切片类型数据
			for j, value := range v {
				if j < len(req.Headers) {
					cell := fmt.Sprintf("%c%d", 'A'+j, i+2)
					f.SetCellValue(sheetName, cell, value)
				}
			}
		default:
			// 其他类型数据
			cell := fmt.Sprintf("A%d", i+2)
			f.SetCellValue(sheetName, cell, v)
		}
	}

	// 设置文件名
	fileName := req.FileName
	if fileName == "" {
		fileName = fmt.Sprintf("export_%s.xlsx", time.Now().Format("20060102150405"))
	}

	// 保存Excel文件
	if err := f.SaveAs(fileName); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save excel file"})
	}

	// 构造响应
	response := ExportResponse{
		FileName: fileName,
		FileUrl:  "/download/" + fileName,
		Size:     0, // 实际应用中应该计算文件大小
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// ExportDemoHandler 导出示例Excel文件
func ExportDemoHandler(c fiber.Ctx) error {
	// 创建Excel文件
	f := excelize.NewFile()

	// 设置工作表名称
	sheetName := "Demo"
	index, err := f.NewSheet(sheetName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create sheet"})
	}

	// 设置默认工作表
	f.SetActiveSheet(index)

	// 写入表头
	headers := []string{"Name", "Age", "Email", "Phone"}
	for i, header := range headers {
		cell := fmt.Sprintf("%c1", 'A'+i)
		f.SetCellValue(sheetName, cell, header)
	}

	// 写入示例数据
	data := [][]interface{}{
		{"John Doe", 30, "john@example.com", "123-456-7890"},
		{"Jane Smith", 25, "jane@example.com", "987-654-3210"},
		{"Bob Johnson", 35, "bob@example.com", "555-123-4567"},
	}

	for i, row := range data {
		for j, value := range row {
			cell := fmt.Sprintf("%c%d", 'A'+j, i+2)
			f.SetCellValue(sheetName, cell, value)
		}
	}

	// 设置文件名
	fileName := fmt.Sprintf("demo_%s.xlsx", time.Now().Format("20060102150405"))

	// 保存Excel文件
	if err := f.SaveAs(fileName); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save excel file"})
	}

	// 构造响应
	response := ExportResponse{
		FileName: fileName,
		FileUrl:  "/download/" + fileName,
		Size:     0, // 实际应用中应该计算文件大小
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
