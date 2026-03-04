package cgi

import (
	"context"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v3"

	"murphyl.com/lego/pkg/sugar"
)

// cgi 模块是CGI相关模块，提供了LegoApp结构体，用于创建和管理Fiber应用程序
// 主要功能包括：创建应用、挂载路由、启动服务等

var sugarLogger = sugar.NewSugarLogger()

type LegoApp struct {
	app *fiber.App
	ctx context.Context
}

type AppContext interface {
	AppTitle() string
	BindAddress() string
}

type AppHandler interface {
	RegisterRoutes(router fiber.Router)
}

func NewLegoApp(appConfig AppContext, opts ...LegoOption) *LegoApp {
	ac := &fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		AppName:       appConfig.AppTitle(),
	}
	// 应用可选配置
	for _, opt := range opts {
		opt(ac)
	}
	// 应用服务
	la := &LegoApp{app: fiber.New(*ac), ctx: context.Background()}
	// 注册关闭前钩子
	la.app.Hooks().OnPreShutdown(func() error {
		sugarLogger.Infoln("Server is shutting down...")
		return nil
	})
	return la
}

type LegoOption = func(cfg *fiber.Config)

func UseFiberService(service fiber.Service) LegoOption {
	return func(cfg *fiber.Config) {
		cfg.Services = append(cfg.Services, service)
	}
}

func (la *LegoApp) RouterGroup(path string, handlers ...any) {
	la.app.Group(path, handlers...)
}

func (la *LegoApp) Mount(url string, useRouterGroup func(router fiber.Router)) {
	useRouterGroup(la.app.Group(path.Join("/api", url)))
}

// Group 创建路由组
func (la *LegoApp) Group(prefix string, handlers ...any) fiber.Router {
	return la.app.Group(prefix, handlers...)
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
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
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
