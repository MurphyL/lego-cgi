package entry

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
