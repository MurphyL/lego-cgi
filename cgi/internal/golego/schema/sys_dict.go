package interfaces

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
