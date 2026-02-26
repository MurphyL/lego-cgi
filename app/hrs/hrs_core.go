package main

import (
	"murphyl.com/lego/cgi"
	"murphyl.com/lego/cgi/domain/account"
	"murphyl.com/lego/cgi/misc"
)

type AppConfig struct {
	title string
	addr  string
	// 应用默认数据库
	dsn string
}

func main() {
	cnf := loadConfig()
	app := cgi.NewLegoApp(cnf)
	app.RetrieveOne("/user/profile", account.GetProfilefunc)
	app.Listen(cnf.BindAddress())
}

func loadConfig() *AppConfig {
	appConfig := &AppConfig{}
	misc.LoadProperty(&appConfig.title, "LEGO_APP_TITLE", "接口网关", "应用标题")
	misc.LoadProperty(&appConfig.addr, "LEGO_BIND_ADDR", ":4044", "应用绑定地址")
	misc.LoadProperty(&appConfig.dsn, "DATASOURCE_NAME", "", "数据库连接地址")
	return appConfig
}

func (c AppConfig) AppTitle() string {
	return c.title
}

func (c AppConfig) BindAddress() string {
	return ":4000"
}

func (c AppConfig) DataSourceName() string {
	return ":4000"
}
