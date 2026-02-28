package wework

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=813e2278-7858-432f-8ecb-1ac0c1a41bac

const (
	WeworkChatGroupRobotWebook = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send"
)

type ChatGroupRobot func(ctx context.Context, data any) (*http.Request, error)

// 群机器人/HTTP请求报文
func NewChatGroupRobot(robotKey string) (ChatGroupRobot, error) {
	robotWebhook, err := url.Parse(WeworkChatGroupRobotWebook)
	if err != nil {
		return nil, fmt.Errorf("解析URL出错：%v", err.Error())
	}
	robotWebhook.RawQuery = url.Values{"key": []string{robotKey}}.Encode()
	robotEndpoint := robotWebhook.String()
	return func(ctx context.Context, data any) (*http.Request, error) {
		var body []byte
		switch v := data.(type) {
		case []byte:
			body = v
		case string:
			body = []byte(v)
		default:
			tmp, err := json.Marshal(v)
			if err != nil {
				return nil, err
			}
			body = tmp
		}
		if req, err := http.NewRequestWithContext(ctx, http.MethodPost, robotEndpoint, bytes.NewReader(body)); err == nil {
			return req, nil
		} else {
			return nil, fmt.Errorf("构造报文出错：%v", err.Error())
		}
	}, nil
}
