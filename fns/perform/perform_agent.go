package perform

/**
 * 第三方支持：如开发平台，通用接口等
 */

type PerformAgent[R any, P any] func(P) (R, error)

// NewPreformAgent 调用预置功能
func NewPreformAgent[R any, P any](request func(P) (R, error)) PerformAgent[R, P] {
	return func(args P) (R, error) {
		return request(args)
	}
}
