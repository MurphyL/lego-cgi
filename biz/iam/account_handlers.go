package iam

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

// 密钥，实际应用中应该从配置文件或环境变量中获取
var jwtSecret = []byte("your-secret-key")

// LoginHandler 登录处理函数
func LoginHandler(c fiber.Ctx) error {
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
	}

	return c.Status(fiber.StatusOK).JSON(LoginResponse{
		Token:     token,
		ExpiresAt: expiresAt,
		User:      user,
	})
}

// LogoutHandler 登出处理函数
func LogoutHandler(c fiber.Ctx) error {
	// 实际应用中可能需要将token加入黑名单
	return c.Status(fiber.StatusOK).JSON(LogoutResponse{
		Success: true,
	})
}

// GetUserProfileHandler 获取用户信息处理函数
func GetUserProfileHandler(c fiber.Ctx) error {
	// 实际应用中应该从token中获取用户信息
	// 这里只是模拟用户信息
	user := Account{
		ID:       1,
		PersonID: 1,
		Username: "admin",
		Mobile:   "13800138000",
		Email:    "admin@example.com",
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
func ResetPasswordHandler(c fiber.Ctx) error {
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
