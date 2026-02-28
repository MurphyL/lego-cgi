package drivers

import (
	"context"

	"gorm.io/gorm"
)

func NewGromRepo(key string, gorm gorm.DB) *GormRepo {
	return &GormRepo{key: key, gorm: gorm}
}

type GormRepo struct {
	gorm gorm.DB
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

func (r GormRepo) RetrieveOne(dest interface{}, conds ...interface{}) error {
	return r.gorm.Take(dest, conds...).Error
}

func (r GormRepo) RetrieveAll(dest interface{}, conds ...interface{}) error {
	return r.gorm.Find(dest, conds...).Error
}
