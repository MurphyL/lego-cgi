package dal

import (
	"gorm.io/gorm"
	"murphyl.com/lego/dal/connectors"
)

func New(productKind string) DataAccessLayer {
	switch productKind {
	case "mysql":
		return connectors.NewMySqlConnection()
	default:
		return nil
	}
}

type DataAccessLayer interface{}

type GormRepo struct {
	gormDial *gorm.Dialector
}
