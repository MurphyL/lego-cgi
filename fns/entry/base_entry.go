package entry

import "time"

type BaseEntry struct {
	ID       uint64     `json:"id" gorm:"primaryKey"`
	Status   StatusEnum `json:"status" gorm:"index,default:1"`
	CreateAt time.Time  `json:"createAt" gorm:"autoCreateTime"`
	CreateBy uint64
	UpdateAt time.Time `json:"updateTime" gorm:"autoUpdateTime"`
	UpdateBy uint64
}
