package wecom

import (
	"bytes"
	"fmt"
	"net/http"

	"murphyl.com/lego/fns"
	"murphyl.com/lego/fns/oauth2"
)

const (
	oauthTokenWebhook = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%v&corpsecret=%v"
)

type WecomOauth2Client fns.PerformClient

func NewWecomOauth2Client(grantType oauth2.GrantType, clientId, clientSecret string) *WecomOauth2Client {
	endpoint := fmt.Sprintf(oauthTokenWebhook, clientId, clientSecret)
	return &WecomOauth2Client{
		HttpClient: http.DefaultClient,
		PerformAgent: func(b []byte) (*http.Request, error) {
			httpReq, err := http.NewRequest(http.MethodGet, endpoint, bytes.NewReader(b))
			if err != nil {
				return nil, fmt.Errorf("创建HTTP请求出错：%v", err.Error())
			}
			return httpReq, nil
		},
	}
}

func (c *WecomOauth2Client) GetAccessToken() (oauth2.GrantedAccessToken, error) {
	return nil, nil
}
