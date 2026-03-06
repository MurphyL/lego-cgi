package ecp

import (
	"time"

	"murphyl.com/lego/fns/entry"
)

// EnterpriseType 企业类型枚举
type EnterpriseType string

// 企业类型
const (
	EnterpriseTypePrivate EnterpriseType = "private" // 民营企业
	EnterpriseTypeState   EnterpriseType = "state"   // 国有企业
	EnterpriseTypeForeign EnterpriseType = "foreign" // 外资企业
)



// 企业状态
const (
	EnterpriseStatusNormal    entry.StatusEnum = 210 // 正常运营
	EnterpriseStatusSuspended entry.StatusEnum = 220 // 暂停运营
	EnterpriseStatusCanceled  entry.StatusEnum = 230 // 已注销
)

type Enterprise struct {
	entry.BaseEntry
	FunllName      string           // 企业全称
	ShortName      string           // 企业简称（如你之前问的“企协平台”所属企业简称）
	EnterpriseType EnterpriseType   // 企业类型
	Status         entry.StatusEnum // 企业状态
	EstablishTime  time.Time        // 成立时间
	Address        string           // 注册地址
	Contact        string           // 联系人
	Phone          string           // 联系电话
}
