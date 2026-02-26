package connectors

import (
	"fmt"
	"net/url"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// root:123456@tcp(localhost:3306)/tizi365?charset=utf8&parseTime=True&loc=Local
type MySqlConn struct {
	Username string
	Password string
	Protocol string
	Address  string
	Database string
}

func NewMySqlConnection() gorm.Dialector {
	return mysql.Open("")
}

func (conn *MySqlConn) DatasourceName() string {
	u := url.URL{
		User: url.UserPassword(conn.Username, conn.Password),
		Host: fmt.Sprintf("%s(%s)", conn.Protocol, conn.Address),
		Path: conn.Database,
	}
	return u.String()
}
