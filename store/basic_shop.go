package store

//
type BasicShop struct {
	Id           string `gorm:"column:id;primary_key"`
	Name         string `gorm:"column:name"`
	PlatformCode string `gorm:"column:platform_code"`
	AppInfo      string `gorm:"column:app_info"`
	Enabled      string `gorm:"column:enabled"`
}

// TableName
func (*BasicShop) TableName() string {
	return "basic_shop"
}
