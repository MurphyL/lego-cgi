package dal

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"murphyl.com/lego/dal/connectors"
	"murphyl.com/lego/misc"
)

var sugarLogger = misc.NewSugarLogger()

func New(productKind string, dsn string) DataAccessLayer {
	switch productKind {
	case "mysql", "sqlite":
		repo := &GormRepo{ctx: context.Background()}
		switch productKind {
		case "mysql":
			repo.gorm = connectors.OpenMySqlConnection(dsn)
		default:
			panic("不支持的数据库类型：" + productKind)
		}
		return repo
	default:
		panic("不支持的数据访问产品：" + productKind)
	}

}

type DataAccessLayer interface {
	fiber.Service
}

type GormRepo struct {
	ctx  context.Context
	gorm gorm.Dialector
}

func (a GormRepo) Start(ctx context.Context) error {
	return nil
}

func (a GormRepo) String() string {
	sugarLogger.Info("DataAccessLayer Start:", a.ctx.Value("key"))
	return "default"
}

func (a GormRepo) State(ctx context.Context) (string, error) {
	return "", nil
}

func (a GormRepo) Terminate(ctx context.Context) error {
	sugarLogger.Info("DataAccessLayer Terminate")
	return nil
}
