package main

import (
	"murphyl.com/app/idm"
	"murphyl.com/app/sys"
	"murphyl.com/lego/cgi"
	"murphyl.com/lego/dal"
	"murphyl.com/lego/misc"
)

var (
	AppTitle       = "房源智管"
	BindAddr       = ":4000"
	DataSourceName = ""
)

type AppConfig struct {
	title string
	addr  string
	// 应用默认数据库
	dsn string
}

func main() {
	cnf := loadConfig()
	dao := dal.New("mysql", cnf.dsn)
	app := cgi.NewLegoApp(cnf, cgi.UseFiberService(dao))
	app.Mount("/account", idm.UseIdentifyManager)
	app.Mount("/system", sys.UseSystemDictManager)
	app.Serve(cnf.BindAddress())
}

func loadConfig() *AppConfig {
	appConfig := &AppConfig{}
	misc.LoadProperty(&appConfig.title, "LEGO_APP_TITLE", AppTitle, "应用标题")
	misc.LoadProperty(&appConfig.addr, "LEGO_BIND_ADDR", BindAddr, "应用绑定地址")
	misc.LoadProperty(&appConfig.dsn, "DATASOURCE_NAME", DataSourceName, "数据库连接地址")
	return appConfig
}

func (c AppConfig) AppTitle() string {
	return c.title
}

func (c AppConfig) BindAddress() string {
	return c.addr
}

func (c AppConfig) DataSourceName() string {
	return c.dsn
}
