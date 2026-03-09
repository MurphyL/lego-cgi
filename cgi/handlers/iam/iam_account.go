package iam

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
	"murphyl.com/lego/fns/sugar"
)

var sugarLogger = sugar.NewSugarLogger()

// NewAccountHandler 创建账户处理器
func NewAccountHandler(dao *gorm.DB) *accountHandler {
	return &accountHandler{db: dao}
}

type accountHandler struct {
	db *gorm.DB
}

func (ah *accountHandler) RegisterRoutes(router fiber.Router) {
}

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
