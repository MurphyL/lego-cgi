package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"murphyl.com/app/prs/handlers/analytics"
	"murphyl.com/app/prs/handlers/property"
	"murphyl.com/app/prs/handlers/tenant"

	"murphyl.com/lego/biz/conf"
	"murphyl.com/lego/biz/erp"
	"murphyl.com/lego/biz/fin"
	"murphyl.com/lego/biz/iam"
	"murphyl.com/lego/cgi"
	"murphyl.com/lego/dal"
	"murphyl.com/lego/fns"
)

// hrs 模块是人力资源系统模块，使用Fiber框架

var (
	AppTitle       = "房源智管"
	AppVersion     = "0.0.1"
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
	// 加载配置
	cnf := loadConfig()
	// 初始化数据库连接
	dao := dal.New(mysql.Open, cnf.dsn)
	// 初始化应用及上下文
	app := cgi.NewLegoApp(cnf)
	// 挂载账户管理路由
	app.Mount("/account", iam.NewAccountHandler(dao))
	// 挂载系统配置路由
	app.Mount("/system", conf.NewSystemDictHandler(dao))
	// 挂载财务管理路由
	app.Mount("/finance", fin.NewFinanceHandler(dao))
	// 挂在合同管理路由
	app.Mount("/contract", erp.NewContractHandler(dao))
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
	fns.LoadProperty(&appConfig.title, "LEGO_APP_TITLE", AppTitle)
	fns.LoadProperty(&appConfig.addr, "LEGO_BIND_ADDR", BindAddr)
	fns.LoadProperty(&appConfig.dsn, "GO_DSN_MYSQL", DataSourceName)
	return appConfig
}

func (c AppConfig) AppTitle() string {
	return fmt.Sprintf("%s v%s", c.title, AppVersion)
}
