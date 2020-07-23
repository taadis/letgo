package store

//
type BasicPlatform struct {
	Id      string `gorm:"column:id;primary_key"`
	Name    string `gorm:"column:name"`
	Code    string `gorm:"column:code"`
	Enabled bool   `gorm:"column:enabled"`
}

// TableName
func (*BasicPlatform) TableName() string {
	return "basic_platform"
}
