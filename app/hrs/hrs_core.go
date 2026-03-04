package main

import (
	"gorm.io/driver/mysql"
	"murphyl.com/app/hrs/handlers/analytics"
	"murphyl.com/app/hrs/handlers/property"
	"murphyl.com/app/hrs/handlers/tenant"

	"murphyl.com/lego/biz/contract"
	"murphyl.com/lego/biz/finance"
	"murphyl.com/lego/biz/general"
	"murphyl.com/lego/biz/iam"
	"murphyl.com/lego/cgi"
	"murphyl.com/lego/dal"
	"murphyl.com/lego/fns/shared"
)

// hrs 模块是人力资源系统模块，使用Fiber框架

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
	dao := dal.New(mysql.Open, cnf.dsn)
	app := cgi.NewLegoApp(cnf)
	app.Mount("/account", iam.NewAccountHandler(dao))
	app.Mount("/system", general.NewSystemDictHandler(dao))
	app.Mount("/finance", finance.NewFinanceHandler(dao))
	app.Mount("/contract", contract.NewContractHandler(dao))
	// 挂载租户管理路由
	app.Mount("/tenant", tenant.NewTenantHandler(dao))
	// 挂载房源管理路由
	app.Mount("/hrs", property.NewPropertyHandler(dao))
	// 挂载数据分析路由
	app.Mount("/hrs/rpt", analytics.NewAnalyticsHandler(dao))
	// 运行服务
	app.Serve(cnf.addr)
}

func loadConfig() *AppConfig {
	appConfig := &AppConfig{}
	shared.LoadProperty(&appConfig.title, "LEGO_APP_TITLE", AppTitle, "应用标题")
	shared.LoadProperty(&appConfig.addr, "LEGO_BIND_ADDR", BindAddr, "应用绑定地址")
	shared.LoadProperty(&appConfig.dsn, "GO_DSN_MYSQL", DataSourceName, "数据库连接地址")
	return appConfig
}

func (c AppConfig) AppTitle() string {
	return c.title
}

func (c AppConfig) DataSourceName() string {
	return c.dsn
}
