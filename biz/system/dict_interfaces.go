package system

type DictType struct {
	DictCode string
	DictName string
}

type DictItem struct {
	DictCode  string
	ItemLabel string
	ItemValue string
}

type DictGroup struct {
	Items []DictItem
}

func (i DictItem) TableName() string {
	return "sys_dict_item"
}
