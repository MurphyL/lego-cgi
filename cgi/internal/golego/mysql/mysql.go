package mysql

import (
	"fmt"
	"net/url"
)

// root:123456@tcp(localhost:3306)/tizi365?charset=utf8&parseTime=True&loc=Local
type MySqlConn struct {
	Username string
	Password string
	Protocol string
	Address  string
	Database string
}

func (conn *MySqlConn) DatasourceName() string {
	u := url.URL{
		User: url.UserPassword(conn.Username, conn.Password),
		Host: fmt.Sprintf("%s(%s)", conn.Protocol, conn.Address),
		Path: conn.Database,
	}
	return u.String()
}
