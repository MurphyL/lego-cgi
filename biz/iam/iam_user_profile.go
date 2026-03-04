package iam

// UpdateProfileRequest 更新用户资料请求
type UpdateProfileRequest struct {
	Email  string `json:"email"`  // 邮箱
	Mobile string `json:"mobile"` // 手机号
	Avatar string `json:"avatar"` // 头像URL
}

// UpdateProfileResponse 更新用户资料响应
type UpdateProfileResponse struct {
	Success bool    `json:"success"`
	User    Account `json:"user"`
}

// UserProfileResponse 用户信息响应
type UserProfileResponse struct {
	User       Account    `json:"user"`
	PersonInfo PersonInfo `json:"personInfo"`
}
