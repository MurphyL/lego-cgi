package iam

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// iam 模块是身份与访问管理模块，包含用户管理、RBAC权限控制、租户管理等功能
// 主要功能包括：用户登录、登出、获取用户信息、重置密码、获取验证码等

// Account 可登录账号
type Account struct {
	ID       uint   `gorm:"primarykey" json:"id"`
	PersonID uint   `gorm:"uniqueIndex" json:"personId"`
	Username string `gorm:"uniqueIndex" json:"username"`
	Password string `json:"-"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"` // 头像URL
}

// 映射表名
func (a Account) TableName() string {
	return "sys_account"
}

// CaptchaRequest 获取验证码请求
type CaptchaRequest struct {
	Type string `json:"type"` // 验证码类型：login, reset, register
}

// CaptchaResponse 获取验证码响应
type CaptchaResponse struct {
	Key  string `json:"key"`
	Data string `json:"data"` // 验证码图片数据或验证码代码
}

// 密钥，实际应用中应该从配置文件或环境变量中获取
var jwtSecret = []byte("your-secret-key")

// NewAccountHandler 创建账户处理器
func NewAccountHandler(dao *gorm.DB) func(router fiber.Router) {
	return func(router fiber.Router) {
		h := &AccountHandler{db: dao}
		h.RegisterRoutes(router)
	}
}

type AccountHandler struct {
	db *gorm.DB
}

func (ah *AccountHandler) RegisterRoutes(router fiber.Router) {
	router.Get("/profile", ah.GetUserProfileHandler)
	router.Put("/profile", ah.UpdateProfileHandler)
	router.Post("/reset-password", ah.ResetPasswordHandler)
}

// GetUserProfileHandler 获取用户信息处理函数
func (ah *AccountHandler) GetUserProfileHandler(c fiber.Ctx) error {
	// 实际应用中应该从token中获取用户信息
	// 这里只是模拟用户信息
	user := Account{
		ID:       1,
		PersonID: 1,
		Username: "admin",
		Mobile:   "13800138000",
		Email:    "admin@example.com",
		Avatar:   "https://example.com/avatar.jpg", // 模拟头像URL
	}

	personInfo := PersonInfo{
		Id:         1,
		RealName:   "管理员",
		IdCardType: "身份证",
		IdCardNo:   "110101199001011234",
	}

	return c.Status(fiber.StatusOK).JSON(UserProfileResponse{
		User:       user,
		PersonInfo: personInfo,
	})
}

// ResetPasswordHandler 重置密码处理函数
func (ah *AccountHandler) ResetPasswordHandler(c fiber.Ctx) error {
	var req ResetPasswordRequest
	body := c.Body()
	if err := json.Unmarshal(body, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if !req.ValidRequest() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request parameters"})
	}

	// 验证验证码（实际应用中应该从缓存中获取验证码进行验证）
	// if !verifyCaptcha(req.CaptchaKey, req.CaptchaCode) {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid captcha"})
	// }

	// 验证旧密码（实际应用中应该从数据库中查询）
	// 这里只是模拟验证
	if req.Username != "admin" || req.OldPassword != "password" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid username or old password"})
	}

	// 更新密码（实际应用中应该更新数据库）

	return c.Status(fiber.StatusOK).JSON(ResetPasswordResponse{
		Success: true,
	})
}

// CaptchaHandler 获取验证码处理函数
func CaptchaHandler(c fiber.Ctx) error {
	var req CaptchaRequest
	body := c.Body()
	if err := json.Unmarshal(body, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// 生成验证码（实际应用中应该生成图片验证码）
	// 这里只是模拟生成验证码
	key := "captcha-key-" + time.Now().String()
	data := "1234" // 实际应用中应该生成随机验证码

	// 将验证码存储到缓存中（实际应用中应该使用redis等缓存）
	// cache.Set(key, data, 5*time.Minute)

	return c.Status(fiber.StatusOK).JSON(CaptchaResponse{
		Key:  key,
		Data: data,
	})
}

// UpdateProfileHandler 更新用户资料处理函数
func (ah *AccountHandler) UpdateProfileHandler(c fiber.Ctx) error {
	var req UpdateProfileRequest
	body := c.Body()
	if err := json.Unmarshal(body, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// 实际应用中应该从token中获取用户信息，然后更新数据库
	// 这里只是模拟更新用户资料
	user := Account{
		ID:       1,
		PersonID: 1,
		Username: "admin",
		Mobile:   req.Mobile,
		Email:    req.Email,
		Avatar:   req.Avatar,
	}

	return c.Status(fiber.StatusOK).JSON(UpdateProfileResponse{
		Success: true,
		User:    user,
	})
}

// ResetPasswordRequest 重置密码请求
type ResetPasswordRequest struct {
	Username    string `json:"username"`
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
	CaptchaCode string `json:"captchaCode"`
	CaptchaKey  string `json:"captchaKey"`
}

func (r *ResetPasswordRequest) ValidRequest() bool {
	return r.Username != "" && r.OldPassword != "" && r.NewPassword != "" && r.CaptchaCode != "" && r.CaptchaKey != ""
}

// ResetPasswordResponse 重置密码响应
type ResetPasswordResponse struct {
	Success bool `json:"success"`
}

// generateToken 生成JWT token
func generateToken(username string) (string, int64, error) {
	expiresAt := time.Now().Add(24 * time.Hour).Unix()

	claims := jwt.MapClaims{
		"username": username,
		"exp":      expiresAt,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", 0, err
	}

	return tokenString, expiresAt, nil
}

// verifyCaptcha 验证验证码
func verifyCaptcha(key, code string) bool {
	// 实际应用中应该从缓存中获取验证码进行验证
	return true
}
