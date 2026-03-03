package excel

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
