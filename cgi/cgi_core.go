package cgi

import (
	"context"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"gorm.io/gorm"
)

type LegoApp struct {
	*fiber.App
	*gorm.DB
	ctx context.Context
}

type AppContext interface {
	AppTitle() string
	BindAddress() string
}

func NewLegoApp(appConfig AppContext) *LegoApp {
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
	la.App = fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		AppName:       appConfig.AppTitle(),
	})
	// 配置日志输出到文件
	file, _ := os.Create("app.log")
	// defer file.Close() // TODO close fh
	la.App.Use(logger.New(logger.Config{
		Stream: file, // 替代 v2 的 Output
		Format: "${time} ${status} ${method} ${path}\n",
	}))
	return la
}

func (la *LegoApp) UseService(servicePtr fiber.Service) {
	lac := la.App.Config()
	lac.Services = append(lac.Services, servicePtr)
}

func (la *LegoApp) RetrieveOne(endpoint string, handler func(c fiber.Ctx) error) {
	la.App.Get("/api"+endpoint, handler)
}
