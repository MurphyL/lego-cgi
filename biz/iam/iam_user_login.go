package iam

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

// LoginMethod 登录方式
type LoginMethod string

// 登录类型
const (
	LoginMethodPassword     LoginMethod = "password"      // 密码登录
	LoginMethodLDAP         LoginMethod = "ldap"          // LDAP登录
	LoginMethodEmail        LoginMethod = "email_code"    // 邮箱验证码登录
	LoginMethodPhone        LoginMethod = "phone_code"    // 手机验证码登录
	LoginMethodWechatQrcode LoginMethod = "wechat_qrcode" // 微信二维码登录
	LoginMethodAlipayQrcode LoginMethod = "alipay_qrcode" // 支付宝二维码登录
)

// PasswordActionType 密码登录的相关操作
type PasswordActionType string

// 密码登录操作
const (
	PasswordActionTypeLogin    PasswordActionType = "login"    // 登录
	PasswordActionTypeReset    PasswordActionType = "reset"    // 重置密码
	PasswordActionTypeRegister PasswordActionType = "register" // 注册
)

// PasswordLoginArgs 登录或者重置密码
type PasswordLoginArgs struct {
	Action      PasswordActionType `json:"action"`
	Username    string             `json:"username"`
	Password    string             `json:"password"`
	CaptchaCode string             `json:"captchaCode"`
	CaptchaKey  string             `json:"captchaKey"`
}

func (a *PasswordLoginArgs) ValidRequest() bool {
	return a.Username != "" && a.Password != "" && a.CaptchaCode != "" && a.CaptchaKey != ""
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token     string  `json:"token"`
	ExpiresAt int64   `json:"expiresAt"`
	User      Account `json:"user"`
}

// LogoutResponse 登出响应
type LogoutResponse struct {
	Success bool `json:"success"`
}

type LoginHandler struct {
	db *gorm.DB
}

func (ah *LoginHandler) RegisterRoutes(router fiber.Router) {
	router.Post("/login", ah.LoginHandler)
	router.Post("/logout", ah.LogoutHandler)
}

// LoginHandler 登录处理函数
func (ah *LoginHandler) LoginHandler(c fiber.Ctx) error {
	var req PasswordLoginArgs
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

	// 验证用户名和密码（实际应用中应该从数据库中查询）
	// 这里只是模拟验证
	if req.Username != "admin" || req.Password != "password" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid username or password"})
	}

	// 生成JWT token
	token, expiresAt, err := generateToken(req.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	// 模拟用户信息
	user := Account{
		ID:       1,
		PersonID: 1,
		Username: req.Username,
		Mobile:   "13800138000",
		Email:    "admin@example.com",
		Avatar:   "https://example.com/avatar.jpg", // 模拟头像URL
	}

	return c.Status(fiber.StatusOK).JSON(LoginResponse{
		Token:     token,
		ExpiresAt: expiresAt,
		User:      user,
	})
}

// LogoutHandler 登出处理函数
func (ah *LoginHandler) LogoutHandler(c fiber.Ctx) error {
	// 实际应用中可能需要将token加入黑名单
	return c.Status(fiber.StatusOK).JSON(LogoutResponse{
		Success: true,
	})
}
