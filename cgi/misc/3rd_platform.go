package misc

import "net/http"

/**
 * 第三方支持：如开发平台，通用接口等
 */

// 调用HTTP接口
func ApplyHttpRequest(request *http.Request) (*http.Response, error) {
	return nil, nil
}

// 调用预置功能
func NewPreformAgent[R any, P any](request func(P) (R, error)) func(P) (R, error) {
	return func(args P) (R, error) {
		return request(args)
	}
}
