package main

import (
	"flag"

	"murphyl.com/lego/cgi/internal/golego"
	"murphyl.com/lego/cgi/internal/golego/interfaces"
)

var (
	Version   = "development"
	BuildTime = ""
	GitHash   = ""
)

var appConfig *interfaces.AppConfig

func init() {
	appConfig = &interfaces.AppConfig{}
	flag.StringVar(&appConfig.AppTitle, "title", "接口网关", "应用标题")
	flag.StringVar(&appConfig.BindAddress, "addr", golego.GetEnv("GO_DSN_MYSQL", ":4044"), "服务绑定地址")
	flag.StringVar(&appConfig.DataSourceName, "dsn", golego.GetEnv("GO_DSN_MYSQL", ""), "数据库连接")
}

func main() {
	app := golego.NewLegoApp(appConfig)
	app.Listen(appConfig.BindAddress)
}
