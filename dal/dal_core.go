package dal

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"murphyl.com/lego/dal/connectors"
	"murphyl.com/lego/udf"
)

const (
	namespace = "dal"
)

var sugarLogger = udf.NewSugarLogger()

func RefKey(objectKey string) string {
	return udf.ObjectKey(namespace, objectKey)
}

func New(objectKey, productKind string, dsn string) DataAccessLayer {
	switch productKind {
	case "mysql", "sqlite":
		repo := &GormRepo{key: RefKey(objectKey)}
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
	key  string
}

func (r GormRepo) Start(ctx context.Context) error {
	return nil
}

func (r GormRepo) String() string {
	return r.key
}

func (r GormRepo) State(ctx context.Context) (string, error) {
	return "", nil
}

func (r GormRepo) Terminate(ctx context.Context) error {
	return nil
}
