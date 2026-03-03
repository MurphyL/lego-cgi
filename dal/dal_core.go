package dal

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"murphyl.com/lego/dal/drivers"
	"murphyl.com/lego/pkg/shared"
	"murphyl.com/lego/pkg/sugar"
)

const (
	namespace = "dal"
)

var sugarLogger = sugar.NewSugarLogger()

func RefKey(objectKey string) string {
	return shared.ObjectKey(namespace, objectKey)
}

func New(objectKey, productKind string, dsn string) DataAccessLayer {
	switch productKind {
	case drivers.RdbmsMySql:
		var conn gorm.Dialector
		switch productKind {
		case drivers.RdbmsMySql:
			conn = drivers.NewMySqlConnection(dsn)
		default:
			panic(fmt.Errorf("不支持的数据库类型：%v", productKind))
		}
		dao, err := gorm.Open(conn)
		if err != nil {
			panic(fmt.Errorf("创建Gorm实例出错：%v", err.Error()))
		}
		return drivers.NewGromRepo(RefKey(objectKey), *dao)
	default:
		panic("不支持的数据访问产品：" + productKind)
	}

}

type DataAccessLayer interface {
	fiber.Service

	RetrieveOne(dest interface{}, conds ...interface{}) error
	RetrieveAll(dest interface{}, conds ...interface{}) error
}
