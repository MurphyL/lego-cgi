package system

import (
	"time"
)

// system 模块是系统管理模块，包含数据字典管理等功能
// 主要功能包括：字典类型管理、字典项管理、字典组管理等

type DictType struct {
	ID          uint64    `json:"id"`
	DictCode    string    `json:"dictCode"`
	DictName    string    `json:"dictName"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	Sort        int       `json:"sort"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type DictItem struct {
	ID        uint64    `json:"id"`
	DictCode  string    `json:"dictCode"`
	ItemLabel string    `json:"itemLabel"`
	ItemValue string    `json:"itemValue"`
	Status    Status    `json:"status"`
	Sort      int       `json:"sort"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type DictGroup struct {
	DictCode string     `json:"dictCode"`
	DictName string     `json:"dictName"`
	Items    []DictItem `json:"items"`
}

// DictTypeRequest 字典类型请求
type DictTypeRequest struct {
	DictCode    string `json:"dictCode"`
	DictName    string `json:"dictName"`
	Description string `json:"description"`
	Status      int    `json:"status"`
	Sort        int    `json:"sort"`
}

// DictItemRequest 字典项请求
type DictItemRequest struct {
	DictCode  string `json:"dictCode"`
	ItemLabel string `json:"itemLabel"`
	ItemValue string `json:"itemValue"`
	Status    int    `json:"status"`
	Sort      int    `json:"sort"`
}

func (i DictItem) TableName() string {
	return "sys_dict_item"
}

func (t DictType) TableName() string {
	return "sys_dict_type"
}
