package iam

import (
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

type loginHandler struct {
	db *gorm.DB
}

func (ah *loginHandler) RegisterRoutes(router fiber.Router) {
	router.Post("/logout", ah.LogoutHandler)
}

// LogoutHandler 登出处理函数
func (ah *loginHandler) LogoutHandler(c fiber.Ctx) error {
	// 实际应用中可能需要将token加入黑名单
	return c.Status(fiber.StatusOK).JSON(LogoutResponse{
		Success: true,
	})
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
