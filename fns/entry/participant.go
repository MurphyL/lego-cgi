package entry

type Participant interface {
	// 主体的类型
	PrincipalType() string
	// 主体的Id
	PrincipalId() uint64
}
