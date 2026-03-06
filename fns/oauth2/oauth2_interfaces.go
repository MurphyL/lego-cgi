package oauth2

type GrantedAccessToken interface {
	AccessToken() string
	TokenType() string
	ExpiresIn() int
}
