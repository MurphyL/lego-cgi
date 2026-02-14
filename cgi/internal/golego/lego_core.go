package golego

import (
	"context"
	"os"
	"strings"

	"github.com/gofiber/fiber/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"murphyl.com/lego/cgi/internal/golego/interfaces"
)

type LegoApp struct {
	*fiber.App
	*gorm.DB
	ctx context.Context
}

func NewLegoApp(appConfig *interfaces.AppConfig) *LegoApp {
	conf := &interfaces.AppConfig{}
	app := &LegoApp{ctx: context.Background()}
	app.App = fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		AppName:       conf.AppTitle,
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

func NewLegoRepo(ctx context.Context) *gorm.DB {
	dsn := ctx.Value(UniqueId("REPO", "DEFAULT")).(string)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func (app *LegoApp) UseDomain(domain interfaces.Domain) {

}
