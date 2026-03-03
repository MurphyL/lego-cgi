package system

type ResourceScope string

const (
	Global ResourceScope = "global"
)

type ScopeEntry struct {
	Scope ResourceScope
}

// Status 状态枚举
type Status uint8

const (
	StatusDisabled Status = 0 // 禁用
	StatusEnabled  Status = 1 // 启用
	StatusDeleted  Status = 2 // 逻辑删除
)

func (s Status) IsEnabled() bool {
	return s == StatusEnabled
}

func (s Status) IsDisabled() bool {
	return s == StatusDisabled
}

func (s Status) IsDeleted() bool {
	return s == StatusDeleted
}
