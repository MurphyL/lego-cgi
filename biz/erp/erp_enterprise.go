package erp

import (
	"errors"
	"time"

	"murphyl.com/lego/fns/entry"
)

// EnterpriseType 企业类型枚举
type EnterpriseType string

// 企业类型 - 国有企业、民营企业、外资企业、合资企业
const (
	EnterpriseTypeState   EnterpriseType = "state"   // 国有企业
	EnterpriseTypePrivate EnterpriseType = "private" // 民营企业
	EnterpriseTypeForeign EnterpriseType = "foreign" // 外资企业
	EnterpriseTypeJoint   EnterpriseType = "joint"   // 合资企业
)

// 企业状态枚举
type EnterpriseStatus string

// 企业状态 - 正常经营、暂停运营、已注销、吊销
const (
	EnterpriseStatusNormal    EnterpriseStatus = "normal"    // 正常运营
	EnterpriseStatusSuspended EnterpriseStatus = "suspended" // 暂停运营
	EnterpriseStatusCanceled  EnterpriseStatus = "canceled"  // 已注销
	EnterpriseStatusRevoked   EnterpriseStatus = "revoked"   // 吊销
)

type Enterprise struct {
	entry.BaseEntry
	USCC              string         // 统一社会信用代码
	FunllName         string         // 企业全称
	ShortName         string         // 企业简称（如你之前问的“企协平台”所属企业简称）
	EnterpriseType    EnterpriseType // 企业类型
	RegisteredCapital float64        `json:"registered_capital"` // 注册资本（万元）
	RegisteredAddress string         // 注册地址
	EstablishTime     time.Time      // 成立时间
	ContactPhone      string         // 联系电话
}

// 基础信息验证
func (e *Enterprise) ValidateBaseInfo() error {
	// 统一社会信用代码验证（18位）
	if e.USCC != "" && len(e.USCC) != 18 {
		return errors.New("统一社会信用代码必须为18位")
	}
	// 公司名称非空验证
	if e.FunllName == "" {
		return errors.New("公司全称不能为空")
	}
	return nil
}
