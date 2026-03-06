package oauth2

import "net/http"

/**
OAuth 2.0 规定了四种获得令牌的流程：
- 授权码（authorization-code）
- 隐藏式（implicit）
- 密码式（password）：
- 客户端凭证（client credentials）
*/

type GrantType string

const (
	GrantTypeAuthorizationCode GrantType = "authorization_code"
	GrantTypeImplicit          GrantType = "implicit"
	GrantTypePassword          GrantType = "password"
	GrantTypeClientCredentials GrantType = "client_credentials"
)

func NewClient(grantType GrantType, clientId, clientSecret string) *Client {
	return &Client{GrantType: grantType, ClientId: clientId, ClientSecret: clientSecret}
}

type Client struct {
	GrantType    GrantType `json:"grant_type"`
	ClientId     string    `json:"client_id"`
	ClientSecret string    `json:"client_secret"`
}

func (c *Client) Apply(makeRequest func() *http.Request) {
	req := makeRequest()
	req.Header.Set("Authorization", "Basic "+c.ClientId+":"+c.ClientSecret)
}
