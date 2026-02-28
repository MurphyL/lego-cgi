package system

type ResourceScope string

const (
	Global ResourceScope = "global"
)

type ScopeEntry struct {
	Scope ResourceScope
}
