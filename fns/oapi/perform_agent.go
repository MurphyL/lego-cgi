package oapi

/**
 * 第三方支持：如开发平台，通用接口等
 */

type PerformAgent[R any, P any] func(P) (R, error)
