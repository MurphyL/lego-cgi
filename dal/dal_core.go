package dal

import (
	"fmt"

	"gorm.io/gorm"
)

// dal 模块是数据访问层模块，定义了DataAccessLayer接口，用于数据访问操作
// 主要功能包括：数据的增删改查、分页查询、事务处理等

func New(dial func(string) gorm.Dialector, dsn string) *gorm.DB {
	db, err := gorm.Open(dial(dsn))
	if err != nil {
		panic(fmt.Errorf("创建Gorm实例出错：%v", err.Error()))
	}
	return db
}
