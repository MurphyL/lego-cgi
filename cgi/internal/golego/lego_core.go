package golego

import (
	"context"
	"os"
	"strings"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type LegoApp struct {
	*fiber.App
	*gorm.DB
	ctx context.Context
}

type AppContext struct {
	AppTitle    string
	BindAddress string
	// 应用默认数据库
	DataSourceName string
}

func NewLegoApp(appConfig *AppContext) *LegoApp {
	app := &LegoApp{ctx: context.Background()}
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
	app.App = fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		AppName:       appConfig.AppTitle,
	})
	return app
}

func UniqueId(items ...string) string {
	return strings.Join(items, ":")
}

func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = defaultValue
	}
	return value
}

func (app *LegoApp) RetrieveOne(endpoint string, handler func(c fiber.Ctx) error) {
	app.App.Get("/api"+endpoint, handler)
}
