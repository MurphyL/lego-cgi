package perform

/**
 * 第三方支持：如开发平台，通用接口等
 */

// NewPreformAgent 调用预置功能
func NewPreformAgent[R any, P any](request func(P) (R, error)) func(P) (R, error) {
	return func(args P) (R, error) {
		return request(args)
	}
}
