package drivers

import (
	"context"

	"gorm.io/gorm"
)

func NewGormRepo(key string, gorm *gorm.DB) *GormRepo {
	return &GormRepo{key: key, gorm: gorm}
}

type GormRepo struct {
	gorm *gorm.DB
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

func (r GormRepo) Create(dest interface{}) error {
	return r.gorm.Create(dest).Error
}

func (r GormRepo) Update(dest interface{}, conds ...interface{}) error {
	if len(conds) > 0 {
		return r.gorm.Model(dest).Where(conds[0], conds[1:]...).Updates(dest).Error
	}
	return r.gorm.Model(dest).Updates(dest).Error
}

func (r GormRepo) Delete(dest interface{}, conds ...interface{}) error {
	return r.gorm.Delete(dest, conds...).Error
}

func (r GormRepo) Count(dest interface{}, conds ...interface{}) (int64, error) {
	var count int64
	db := r.gorm.Model(dest)
	if len(conds) > 0 {
		db = db.Where(conds[0], conds[1:]...)
	}
	err := db.Count(&count).Error
	return count, err
}

func (r GormRepo) Page(dest interface{}, page, pageSize int, conds ...interface{}) error {
	offset := (page - 1) * pageSize
	db := r.gorm
	if len(conds) > 0 {
		db = db.Where(conds[0], conds[1:]...)
	}
	return db.Offset(offset).Limit(pageSize).Find(dest).Error
}

func (r GormRepo) Transaction(fn func(tx *gorm.DB) error) error {
	return r.gorm.Transaction(fn)
}

// DB 获取数据库连接
func (r GormRepo) DB() *gorm.DB {
	return r.gorm
}
