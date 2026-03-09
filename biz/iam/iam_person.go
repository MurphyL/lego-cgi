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
