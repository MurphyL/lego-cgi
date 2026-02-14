package main

import (
	"flag"

	"murphyl.com/lego/cgi/internal/golego"
	"murphyl.com/lego/cgi/internal/golego/handler/account"
)

const (
	Version = "development"
)

var appConfig *golego.AppContext = &golego.AppContext{}

func init() {
	flag.StringVar(&appConfig.AppTitle, "title", golego.GetEnv("LEGO_APP_TITLE", "接口网关"), "应用标题")
	flag.StringVar(&appConfig.BindAddress, "addr", golego.GetEnv("LEGO_BIND_ADDR", ":4044"), "服务绑定地址")
	flag.StringVar(&appConfig.DataSourceName, "dsn", "", "MySQL数据库连接")
}

func main() {
	app := golego.NewLegoApp(appConfig)
	app.RetrieveOne("/user/profile", account.GetProfilefunc)
	app.Listen(appConfig.BindAddress)
}
