package main

import (
	"murphyl.com/app/hrs/handlers/analytics"
	"murphyl.com/app/hrs/handlers/property"
	"murphyl.com/app/hrs/handlers/tenant"
	"murphyl.com/lego/biz"
	"murphyl.com/lego/cgi"
	"murphyl.com/lego/dal"
	"murphyl.com/lego/pkg/shared"
)

// hrs 模块是人力资源系统模块，使用Fiber框架
// 主要功能包括：身份管理、系统字典管理等

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
	dao := dal.New(shared.DeafultKey, "mysql", cnf.dsn)
	app := cgi.NewLegoApp(cnf, cgi.UseFiberService(dao))
	app.Mount("/account", biz.UseIdentifyManager)
	app.Mount("/system", biz.UseSystemDictManager)
	app.Mount("/finance", biz.UseFinanceManager(dao.DB()))
	app.Mount("/contract", biz.UseContractManager(dao.DB()))

	// 挂载房源管理路由
	propertyHandler := property.NewPropertyHandler(dao.DB())
	propertyHandler.RegisterRoutes(app.Group("/api"))

	// 挂载租户管理路由
	tenantHandler := tenant.NewTenantHandler(dao.DB())
	tenantHandler.RegisterRoutes(app.Group("/api"))

	// 挂载数据分析路由
	analyticsHandler := analytics.NewAnalyticsHandler(dao.DB())
	analyticsHandler.RegisterRoutes(app.Group("/api"))

	app.Serve(cnf.BindAddress())
}

func loadConfig() *AppConfig {
	appConfig := &AppConfig{}
	shared.LoadProperty(&appConfig.title, "LEGO_APP_TITLE", AppTitle, "应用标题")
	shared.LoadProperty(&appConfig.addr, "LEGO_BIND_ADDR", BindAddr, "应用绑定地址")
	shared.LoadProperty(&appConfig.dsn, "DATASOURCE_NAME", DataSourceName, "数据库连接地址")
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
