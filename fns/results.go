package fns

import (
	"fmt"
)

type Result struct {
	Payload any    `json:"payload"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

func NewResult(ok bool, payload any, message string) *Result {
	return &Result{Success: ok, Payload: payload, Message: message}
}

func NewSuccessResult(payload any) *Result {
	return NewResult(true, payload, "OK")
}

func NewResultViaError(err error) *Result {
	return NewResult(false, nil, fmt.Sprintf("未知错误：%s", err.Error()))
}

func NewResultViaMessage(ok bool, msg string) *Result {
	return NewResult(ok, nil, msg)
}
