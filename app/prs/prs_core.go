package main

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"murphyl.com/lego/cgi"
	"murphyl.com/lego/cgi/handlers/etc"
	"murphyl.com/lego/cgi/handlers/fin"
	"murphyl.com/lego/cgi/handlers/iam"
	"murphyl.com/lego/cgi/handlers/tag"
	"murphyl.com/lego/fns"
)

// prs 房屋租赁系统，使用Fiber框架

var (
	AppTitle       = "房源智管"
	AppVersion     = "0.0.1"
	BindAddr       = ":4000"
	DataSourceName = ""
)

type appConfig struct {
	title string
	addr  string
	// 应用默认数据库
	dsn string
}

func main() {
	// 加载配置
	cnf := loadConfig()
	// 初始化数据库连接
	dao := cgi.NewLegoRepo(mysql.Open, cnf.dsn)
	// 初始化应用及上下文
	app := cgi.NewLegoApp(cnf)
	// 挂载账户管理路由
	app.Handle("/api", &externalHandler{dao: dao.DB})
	// 挂载系统配置路由
	// app.Mount("/system", etc.NewSystemDictHandler(dao))
	// 挂载财务管理路由
	// app.Mount("/fin", fin.NewFinanceHandler(dao))
	// 挂在合同管理路由
	// app.Mount("/erp", erp.NewContractHandler(dao))
	// 挂载房源管理路由
	// app.Mount("/prs", property.NewPropertyHandler(dao))
	// 挂载数据分析路由
	// app.Mount("/prs/rpt", analytics.NewAnalyticsHandler(dao))
	// 运行服务
	app.Serve(cnf.addr)
}

func loadConfig() *appConfig {
	appConfig := &appConfig{}
	fns.LoadProperty(&appConfig.title, "LEGO_APP_TITLE", AppTitle)
	fns.LoadProperty(&appConfig.addr, "LEGO_BIND_ADDR", BindAddr)
	fns.LoadProperty(&appConfig.dsn, "GO_DSN_MYSQL", DataSourceName)
	return appConfig
}
func (c appConfig) AppPhrase() string {
	return "prs"
}
func (c appConfig) AppVersion() string {
	return AppVersion
}
func (c appConfig) AppTitle() string {
	return c.title
}

type externalHandler struct{ dao *gorm.DB }

func (h *externalHandler) RegisterRoutes(router fiber.Router) {
	etc.NewSystemHandler(h.dao).RegisterRoutes(router)
	iam.NewAccountHandler(h.dao).RegisterRoutes(router)
	tag.NewTagHandler(h.dao).RegisterRoutes(router)
	fin.NewFinanceHandler(h.dao).RegisterRoutes(router)
}
