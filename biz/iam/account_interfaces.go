package iam

// PersonInfo 公民信息
type PersonInfo struct {
	Id         uint64 `json:"id"`
	RealName   string `json:"realName"`
	IdCardType string `json:"idCardType"` // 证件类型
	IdCardNo   string `json:"idCardNo"`
}

func (a PersonInfo) TableName() string {
	return "base_person"
}

// Account 可登录账号
type Account struct {
	ID       uint   `gorm:"primarykey" json:"id"`
	PersonID uint   `gorm:"uniqueIndex" json:"personId"`
	Username string `gorm:"uniqueIndex" json:"username"`
	Password string `json:"-"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
}

// 映射表名
func (a Account) TableName() string {
	return "sys_account"
}

// LoginMethod 登录方式
type LoginMethod string

// 登录类型
const (
	LoginMethodPassword     LoginMethod = "password"      // 密码登录
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

// UserProfileResponse 用户信息响应
type UserProfileResponse struct {
	User       Account    `json:"user"`
	PersonInfo PersonInfo `json:"personInfo"`
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

// CaptchaRequest 获取验证码请求
type CaptchaRequest struct {
	Type string `json:"type"` // 验证码类型：login, reset, register
}

// CaptchaResponse 获取验证码响应
type CaptchaResponse struct {
	Key  string `json:"key"`
	Data string `json:"data"` // 验证码图片数据或验证码代码
}
