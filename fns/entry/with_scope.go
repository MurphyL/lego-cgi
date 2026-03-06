package entry

// ResourceScope 资源范围
type ResourceScope string

// ResourceScope 枚举
const (
	ResourceScopeGlobal ResourceScope = "global" // 全局资源范围
	ResourceScopeRole   ResourceScope = "role"   // 角色资源范围
	ResourceScopeUser   ResourceScope = "user"   // 用户资源范围
)
