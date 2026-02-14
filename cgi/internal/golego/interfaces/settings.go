package interfaces

type AppConfig struct {
	AppTitle    string
	BindAddress string
	// 应用默认数据库
	DataSourceName string
	DriverName     string
}

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
