package cgi

import (
	"context"
	"os"
	"os/signal"
	"path"
	"time"

	"github.com/gofiber/fiber/v3"

	"murphyl.com/lego/misc"
)

var sugarLogger = misc.NewSugarLogger()

type LegoApp struct {
	app *fiber.App
	ctx context.Context
}

type AppContext interface {
	AppTitle() string
	BindAddress() string
}

func NewLegoApp(appConfig AppContext, opts ...LegoOption) *LegoApp {
	la := &LegoApp{ctx: context.Background()}
	/**
	var err error
	// 数据库
	app.DB, err = gorm.Open(mysql.Open(appConfig.DataSourceName), &gorm.Config{
		AllowGlobalUpdate: false,
	})
	if err != nil {
		panic("连接默认数据库出错：" + err.Error())
	}
	sqlDb, _ := app.DB.DB()
	sqlDb.SetMaxIdleConns(15)
	sqlDb.SetMaxOpenConns(25)
	sqlDb.SetConnMaxLifetime(5 * time.Minute)
	**/
	// 应用服务
	la.app = fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		AppName:       appConfig.AppTitle(),
	})
	// 应用可选配置
	for _, opt := range opts {
		opt(la)
	}
	// 注册关闭前钩子
	la.app.Hooks().OnPreShutdown(func() error {
		sugarLogger.Infoln("Server is shutting down...")
		return nil
	})
	return la
}

type LegoOption = func(cfg *LegoApp)

func UseFiberService(service fiber.Service) LegoOption {
	return func(la *LegoApp) {
		cfg := la.app.Config()
		cfg.Services = append(cfg.Services, service)
	}
}

func (la *LegoApp) RouterGroup(path string, handlers ...any) {
	la.app.Group(path, handlers...)
}

func (la *LegoApp) Mount(url string, useRouterGroup func(router fiber.Router)) {
	useRouterGroup(la.app.Group(path.Join("/api", url)))
}

func (la *LegoApp) Serve(addr string) {
	// 启动服务器协程
	go func() {
		if err := la.app.Listen(addr); err != nil {
			sugarLogger.Info("Server Shutdown:", err.Error())
		}
	}()
	sugarLogger.Info("Server started:", addr)
	// 监听中断信号并触发优雅关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit
	// 创建带超时的上下文，限制最长等待30秒
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// 优雅关闭
	if err := la.app.ShutdownWithContext(ctx); err != nil {
		sugarLogger.Info("Server failed:", err)
	}
	sugarLogger.Info("Server exited")
}
