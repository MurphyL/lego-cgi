package cgi

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

// dal 模块是数据访问层模块，定义了DataAccessLayer接口，用于数据访问操作
func NewLegoRepo(dial func(string) gorm.Dialector, dsn string) *gorm.DB {
	conf := &gorm.Config{}
	dao, err := gorm.Open(dial(dsn), conf)
	if err != nil {
		panic(fmt.Errorf("创建Gorm实例出错：%v", err.Error()))
	}
	db, _ := dao.DB()
	displayDSN, _ := strings.CutPrefix(dsn, "@")
	sugarLogger.Infoln("数据库信息：", displayDSN)
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	db.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	db.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	db.SetConnMaxLifetime(time.Hour)
	return dao
}
