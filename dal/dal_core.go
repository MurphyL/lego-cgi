package dal

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"murphyl.com/lego/dal/drivers"
	"murphyl.com/lego/pkg/shared"
	"murphyl.com/lego/pkg/sugar"
)

// dal 模块是数据访问层模块，定义了DataAccessLayer接口，用于数据访问操作
// 主要功能包括：数据的增删改查、分页查询、事务处理等

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
		return drivers.NewGromRepo(RefKey(objectKey), dao)
	default:
		panic("不支持的数据访问产品：" + productKind)
	}

}

type DataAccessLayer interface {
	fiber.Service

	RetrieveOne(dest interface{}, conds ...interface{}) error
	RetrieveAll(dest interface{}, conds ...interface{}) error
	Create(dest interface{}) error
	Update(dest interface{}, conds ...interface{}) error
	Delete(dest interface{}, conds ...interface{}) error
	Count(dest interface{}, conds ...interface{}) (int64, error)
	Page(dest interface{}, page, pageSize int, conds ...interface{}) error
	Transaction(fn func(tx *gorm.DB) error) error
}
