package shared

// ResourceScope 资源范围
type ResourceScope string

// ResourceScope 枚举
const (
	// GlobalResourceScope 全局资源范围
	ResourceScopeGlobal ResourceScope = "global"
	// RoleResourceScope 角色资源范围
	ResourceScopeRole ResourceScope = "role"
	// UserResourceScope 用户资源范围
	ResourceScopeUser ResourceScope = "user"
)

type ScopeEntry struct {
	Scope ResourceScope `json:"scope"`
}

// Status 状态枚举
type Status uint8

// Status 枚举
const (
	StatusDisabled Status = 0 // 禁用
	StatusEnabled  Status = 1 // 启用
	StatusDeleted  Status = 2 // 逻辑删除
)

// IsEnabled 是否启用
func (s Status) IsEnabled() bool {
	return s == StatusEnabled
}

// IsDisabled 是否禁用
func (s Status) IsDisabled() bool {
	return s == StatusDisabled
}

// IsDeleted 是否逻辑删除
func (s Status) IsDeleted() bool {
	return s == StatusDeleted
}
