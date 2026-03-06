package property

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
	"murphyl.com/app/prs/middleware"
)

/**
 * Property 房产
 */

// Property 房产基础信息
type Property struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	PropertyCode   string    `gorm:"size:30;uniqueIndex" json:"property_code"`
	PropertyTitle  string    `gorm:"size:100" json:"property_title"`
	OwnerName      string    `gorm:"size:50" json:"owner_name"`
	PropertyCertNo string    `gorm:"size:50;uniqueIndex" json:"property_cert_no"`
	Address        string    `gorm:"size:255" json:"address"`
	Area           float64   `json:"area"`
	RoomType       string    `gorm:"size:20" json:"room_type"`
	Orientation    string    `gorm:"size:10" json:"orientation"`
	Decoration     uint8     `json:"decoration"`
	RoomCount      int       `json:"room_count"`
	Status         uint8     `json:"status"`
	Price          float64   `json:"price"`
	Description    string    `gorm:"type:text" json:"description"`
	CreatorID      uint      `json:"creator_id"`
	CreateTime     time.Time `json:"create_time"`
	UpdateTime     time.Time `json:"update_time"`
}

func (*Property) TableName() string {
	return "hrs_property"
}

// PropertyImage 房源图片
type PropertyImage struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	PropertyID uint      `json:"property_id"`
	ImageURL   string    `gorm:"size:500" json:"image_url"`
	ImageType  uint8     `json:"image_type"`
	SortOrder  int       `json:"sort_order"`
	CreateTime time.Time `json:"create_time"`
}

func (*PropertyImage) TableName() string {
	return "hrs_property_image"
}

// PropertyTag 房源标签
type PropertyTag struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	PropertyID uint      `json:"property_id"`
	TagName    string    `gorm:"size:50" json:"tag_name"`
	CreateTime time.Time `json:"create_time"`
}

func (*PropertyTag) TableName() string {
	return "hrs_property_tag"
}

// PropertyStatusLog 房源状态变更日志
type PropertyStatusLog struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	PropertyID   uint      `json:"property_id"`
	OldStatus    uint8     `json:"old_status"`
	NewStatus    uint8     `json:"new_status"`
	ChangeReason string    `gorm:"size:255" json:"change_reason"`
	OperatorID   uint      `json:"operator_id"`
	CreateTime   time.Time `json:"create_time"`
}

func (*PropertyStatusLog) TableName() string {
	return "hrs_property_status_log"
}

// PropertyViewing 房源带看记录
type PropertyViewing struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	PropertyID  uint      `json:"property_id"`
	TenantID    *uint     `json:"tenant_id"`
	ViewerName  string    `gorm:"size:50" json:"viewer_name"`
	ViewerPhone string    `gorm:"size:20" json:"viewer_phone"`
	ViewTime    time.Time `json:"view_time"`
	Feedback    string    `gorm:"type:text" json:"feedback"`
	NextPlan    string    `gorm:"size:255" json:"next_plan"`
	AgentID     uint      `json:"agent_id"`
	CreateTime  time.Time `json:"create_time"`
}

func (*PropertyViewing) TableName() string {
	return "hrs_property_viewing"
}

// CreatePropertyRequest 创建房源请求
type CreatePropertyRequest struct {
	PropertyTitle  string   `json:"property_title" validate:"required"`
	OwnerName      string   `json:"owner_name" validate:"required"`
	PropertyCertNo string   `json:"property_cert_no" validate:"required"`
	Address        string   `json:"address" validate:"required"`
	Area           float64  `json:"area" validate:"required"`
	RoomType       string   `json:"room_type" validate:"required"`
	Orientation    string   `json:"orientation" validate:"required"`
	Decoration     uint8    `json:"decoration"`
	RoomCount      int      `json:"room_count" validate:"required"`
	Price          float64  `json:"price" validate:"required"`
	Description    string   `json:"description"`
	Images         []Image  `json:"images"`
	Tags           []string `json:"tags"`
}

// Image 图片信息
type Image struct {
	ImageURL  string `json:"image_url" validate:"required"`
	ImageType uint8  `json:"image_type"`
	SortOrder int    `json:"sort_order"`
}

// UpdatePropertyRequest 更新房源请求
type UpdatePropertyRequest struct {
	PropertyTitle string   `json:"property_title"`
	OwnerName     string   `json:"owner_name"`
	Address       string   `json:"address"`
	Area          float64  `json:"area"`
	RoomType      string   `json:"room_type"`
	Orientation   string   `json:"orientation"`
	Decoration    uint8    `json:"decoration"`
	RoomCount     int      `json:"room_count"`
	Price         float64  `json:"price"`
	Description   string   `json:"description"`
	Images        []Image  `json:"images"`
	Tags          []string `json:"tags"`
}

// ChangePropertyStatusRequest 变更房源状态请求
type ChangePropertyStatusRequest struct {
	NewStatus    uint8  `json:"new_status" validate:"required"`
	ChangeReason string `json:"change_reason" validate:"required"`
}

// PropertyQueryRequest 房源查询请求
type PropertyQueryRequest struct {
	AreaMin    *float64 `json:"area_min"`
	AreaMax    *float64 `json:"area_max"`
	PriceMin   *float64 `json:"price_min"`
	PriceMax   *float64 `json:"price_max"`
	RoomType   string   `json:"room_type"`
	Decoration *uint8   `json:"decoration"`
	Status     *uint8   `json:"status"`
	Address    string   `json:"address"`
	Page       int      `json:"page" default:"1"`
	PageSize   int      `json:"page_size" default:"10"`
}

// PropertyHandler 房源处理器
type PropertyHandler struct {
	db *gorm.DB
}

// NewPropertyHandler 创建房源处理器
func NewPropertyHandler(dao *gorm.DB) func(router fiber.Router) {
	return func(router fiber.Router) {
		h := &PropertyHandler{db: dao}
		h.RegisterRoutes(router)
	}
}

// RegisterRoutes 注册路由
func (h *PropertyHandler) RegisterRoutes(router fiber.Router) {
	router.Post("/properties", middleware.AuthMiddleware("property:create"), h.CreateProperty)
	router.Get("/properties", middleware.AuthMiddleware("property:list"), h.ListProperties)
	router.Get("/properties/:id", middleware.AuthMiddleware("property:view"), h.GetProperty)
	router.Put("/properties/:id", middleware.AuthMiddleware("property:update"), h.UpdateProperty)
	router.Delete("/properties/:id", middleware.AuthMiddleware("property:delete"), h.DeleteProperty)
	router.Put("/properties/:id/status", middleware.AuthMiddleware("property:update_status"), h.ChangePropertyStatus)
	router.Post("/properties/:id/viewings", middleware.AuthMiddleware("property:create_viewing"), h.CreateViewing)
	router.Get("/properties/:id/viewings", middleware.AuthMiddleware("property:list_viewings"), h.ListViewings)
}

// CreateProperty 创建房源
func (h *PropertyHandler) CreateProperty(c fiber.Ctx) error {
	var req CreatePropertyRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// 生成房源编号
	propertyCode := "PROP" + time.Now().Format("20060102150405")

	// 创建房源
	property := Property{
		PropertyCode:   propertyCode,
		PropertyTitle:  req.PropertyTitle,
		OwnerName:      req.OwnerName,
		PropertyCertNo: req.PropertyCertNo,
		Address:        req.Address,
		Area:           req.Area,
		RoomType:       req.RoomType,
		Orientation:    req.Orientation,
		Decoration:     req.Decoration,
		RoomCount:      req.RoomCount,
		Status:         0, // 待租
		Price:          req.Price,
		Description:    req.Description,
		CreatorID:      1, // 暂时硬编码，实际应从会话中获取
		CreateTime:     time.Now(),
		UpdateTime:     time.Now(),
	}

	// 开始事务
	tx := h.db.Begin()
	if err := tx.Create(&property).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create property"})
	}

	// 处理图片
	for _, img := range req.Images {
		image := PropertyImage{
			PropertyID: property.ID,
			ImageURL:   img.ImageURL,
			ImageType:  img.ImageType,
			SortOrder:  img.SortOrder,
			CreateTime: time.Now(),
		}
		if err := tx.Create(&image).Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create property image"})
		}
	}

	// 处理标签
	for _, tagName := range req.Tags {
		tag := PropertyTag{
			PropertyID: property.ID,
			TagName:    tagName,
			CreateTime: time.Now(),
		}
		if err := tx.Create(&tag).Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create property tag"})
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to commit transaction"})
	}

	return c.Status(fiber.StatusCreated).JSON(property)
}

// ListProperties 列出房源
func (h *PropertyHandler) ListProperties(c fiber.Ctx) error {
	var req PropertyQueryRequest
	if err := c.Bind().Query(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// 设置默认值
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	// 构建查询
	query := h.db.Model(&Property{})

	// 应用过滤条件
	if req.AreaMin != nil {
		query = query.Where("area >= ?", *req.AreaMin)
	}
	if req.AreaMax != nil {
		query = query.Where("area <= ?", *req.AreaMax)
	}
	if req.PriceMin != nil {
		query = query.Where("price >= ?", *req.PriceMin)
	}
	if req.PriceMax != nil {
		query = query.Where("price <= ?", *req.PriceMax)
	}
	if req.RoomType != "" {
		query = query.Where("room_type = ?", req.RoomType)
	}
	if req.Decoration != nil {
		query = query.Where("decoration = ?", *req.Decoration)
	}
	if req.Status != nil {
		query = query.Where("status = ?", *req.Status)
	}
	if req.Address != "" {
		query = query.Where("address LIKE ?", "%"+req.Address+"%")
	}

	// 计算总数
	var total int64
	query.Count(&total)

	// 分页
	offset := (req.Page - 1) * req.PageSize
	var properties []Property
	if err := query.Offset(offset).Limit(req.PageSize).Order("create_time DESC").Find(&properties).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to list properties"})
	}

	return c.JSON(fiber.Map{
		"total":      total,
		"page":       req.Page,
		"page_size":  req.PageSize,
		"properties": properties,
	})
}

// GetProperty 获取房源详情
func (h *PropertyHandler) GetProperty(c fiber.Ctx) error {
	id := c.Params("id")
	var property Property
	if err := h.db.First(&property, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Property not found"})
	}

	// 获取图片
	var images []PropertyImage
	h.db.Where("property_id = ?", property.ID).Order("sort_order").Find(&images)

	// 获取标签
	var tags []PropertyTag
	h.db.Where("property_id = ?", property.ID).Find(&tags)

	// 获取状态变更记录
	var statusLogs []PropertyStatusLog
	h.db.Where("property_id = ?", property.ID).Order("create_time DESC").Find(&statusLogs)

	return c.JSON(fiber.Map{
		"property":    property,
		"images":      images,
		"tags":        tags,
		"status_logs": statusLogs,
	})
}

// UpdateProperty 更新房源
func (h *PropertyHandler) UpdateProperty(c fiber.Ctx) error {
	id := c.Params("id")
	var property Property
	if err := h.db.First(&property, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Property not found"})
	}

	// 检查权限（谁创建谁修改）
	// 实际应从会话中获取当前用户ID并与property.CreatorID比较

	// 检查状态（已租或已备案房源仅允许修改非核心字段）
	if property.Status == 1 { // 已租
		// 仅允许修改非核心字段
		var req struct {
			Description string  `json:"description"`
			Price       float64 `json:"price"`
		}
		if err := c.Bind().JSON(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
		}

		property.Description = req.Description
		property.Price = req.Price
		property.UpdateTime = time.Now()
	} else {
		// 允许修改所有字段
		var req UpdatePropertyRequest
		if err := c.Bind().JSON(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
		}

		// 更新核心字段
		if req.PropertyTitle != "" {
			property.PropertyTitle = req.PropertyTitle
		}
		if req.OwnerName != "" {
			property.OwnerName = req.OwnerName
		}
		if req.Address != "" {
			property.Address = req.Address
		}
		if req.Area > 0 {
			property.Area = req.Area
		}
		if req.RoomType != "" {
			property.RoomType = req.RoomType
		}
		if req.Orientation != "" {
			property.Orientation = req.Orientation
		}
		property.Decoration = req.Decoration
		if req.RoomCount > 0 {
			property.RoomCount = req.RoomCount
		}
		if req.Price > 0 {
			property.Price = req.Price
		}
		if req.Description != "" {
			property.Description = req.Description
		}
		property.UpdateTime = time.Now()

		// 开始事务
		tx := h.db.Begin()

		// 更新图片
		if len(req.Images) > 0 {
			// 删除旧图片
			if err := tx.Where("property_id = ?", property.ID).Delete(&PropertyImage{}).Error; err != nil {
				tx.Rollback()
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete old images"})
			}

			// 添加新图片
			for _, img := range req.Images {
				image := PropertyImage{
					PropertyID: property.ID,
					ImageURL:   img.ImageURL,
					ImageType:  img.ImageType,
					SortOrder:  img.SortOrder,
					CreateTime: time.Now(),
				}
				if err := tx.Create(&image).Error; err != nil {
					tx.Rollback()
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create property image"})
				}
			}
		}

		// 更新标签
		if len(req.Tags) > 0 {
			// 删除旧标签
			if err := tx.Where("property_id = ?", property.ID).Delete(&PropertyTag{}).Error; err != nil {
				tx.Rollback()
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete old tags"})
			}

			// 添加新标签
			for _, tagName := range req.Tags {
				tag := PropertyTag{
					PropertyID: property.ID,
					TagName:    tagName,
					CreateTime: time.Now(),
				}
				if err := tx.Create(&tag).Error; err != nil {
					tx.Rollback()
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create property tag"})
				}
			}
		}

		// 提交事务
		if err := tx.Commit().Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to commit transaction"})
		}
	}

	// 保存房源信息
	if err := h.db.Save(&property).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update property"})
	}

	return c.JSON(property)
}

// DeleteProperty 删除房源
func (h *PropertyHandler) DeleteProperty(c fiber.Ctx) error {
	id := c.Params("id")
	var property Property
	if err := h.db.First(&property, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Property not found"})
	}

	// 检查状态（仅对"待审核"或"已下架"状态房源开放删除）
	if property.Status != 0 && property.Status != 3 { // 假设3是已下架
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Cannot delete property in current status"})
	}

	// 开始事务
	tx := h.db.Begin()

	// 删除相关数据
	if err := tx.Where("property_id = ?", property.ID).Delete(&PropertyImage{}).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete property images"})
	}

	if err := tx.Where("property_id = ?", property.ID).Delete(&PropertyTag{}).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete property tags"})
	}

	if err := tx.Where("property_id = ?", property.ID).Delete(&PropertyStatusLog{}).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete property status logs"})
	}

	if err := tx.Where("property_id = ?", property.ID).Delete(&PropertyViewing{}).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete property viewings"})
	}

	// 删除房源
	if err := tx.Delete(&property).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete property"})
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to commit transaction"})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}

// ChangePropertyStatus 变更房源状态
func (h *PropertyHandler) ChangePropertyStatus(c fiber.Ctx) error {
	id := c.Params("id")
	var property Property
	if err := h.db.First(&property, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Property not found"})
	}

	var req ChangePropertyStatusRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// 记录旧状态
	oldStatus := property.Status

	// 更新状态
	property.Status = req.NewStatus
	property.UpdateTime = time.Now()

	// 开始事务
	tx := h.db.Begin()

	// 保存房源状态
	if err := tx.Save(&property).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update property status"})
	}

	// 记录状态变更日志
	statusLog := PropertyStatusLog{
		PropertyID:   property.ID,
		OldStatus:    oldStatus,
		NewStatus:    req.NewStatus,
		ChangeReason: req.ChangeReason,
		OperatorID:   1, // 暂时硬编码，实际应从会话中获取
		CreateTime:   time.Now(),
	}
	if err := tx.Create(&statusLog).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create status log"})
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to commit transaction"})
	}

	return c.JSON(property)
}

// CreateViewing 创建带看记录
func (h *PropertyHandler) CreateViewing(c fiber.Ctx) error {
	propertyID := c.Params("id")
	var property Property
	if err := h.db.First(&property, propertyID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Property not found"})
	}

	var req struct {
		TenantID    *uint     `json:"tenant_id"`
		ViewerName  string    `json:"viewer_name" validate:"required"`
		ViewerPhone string    `json:"viewer_phone" validate:"required"`
		ViewTime    time.Time `json:"view_time" validate:"required"`
		Feedback    string    `json:"feedback"`
		NextPlan    string    `json:"next_plan"`
	}

	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	viewing := PropertyViewing{
		PropertyID:  property.ID,
		TenantID:    req.TenantID,
		ViewerName:  req.ViewerName,
		ViewerPhone: req.ViewerPhone,
		ViewTime:    req.ViewTime,
		Feedback:    req.Feedback,
		NextPlan:    req.NextPlan,
		AgentID:     1, // 暂时硬编码，实际应从会话中获取
		CreateTime:  time.Now(),
	}

	if err := h.db.Create(&viewing).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create viewing record"})
	}

	return c.Status(fiber.StatusCreated).JSON(viewing)
}

// ListViewings 列出带看记录
func (h *PropertyHandler) ListViewings(c fiber.Ctx) error {
	propertyID := c.Params("id")
	var property Property
	if err := h.db.First(&property, propertyID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Property not found"})
	}

	var viewings []PropertyViewing
	if err := h.db.Where("property_id = ?", property.ID).Order("view_time DESC").Find(&viewings).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to list viewing records"})
	}

	return c.JSON(viewings)
}
