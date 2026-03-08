package dal

import (
	"fmt"
	"net/url"
	"time"

	"gorm.io/gorm"
	"murphyl.com/lego/fns/sugar"
)

var sugarLogger = sugar.NewSugarLogger()

// dal 模块是数据访问层模块，定义了DataAccessLayer接口，用于数据访问操作
func New(dial func(string) gorm.Dialector, dsn string) *gorm.DB {
	dao, err := gorm.Open(dial(dsn))
	if err != nil {
		panic(fmt.Errorf("创建Gorm实例出错：%v", err.Error()))
	}
	dsnURL, _ := url.Parse(dsn)
	db, _ := dao.DB()
	sugarLogger.Infoln("数据库信息：", dao.Name(), dsnURL)
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	db.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	db.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	db.SetConnMaxLifetime(time.Hour)
	return dao
}
