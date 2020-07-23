package store

//
type BasicPlatform struct {
	Code    string `gorm:"column:id;primary_key"`
	Name    string `gorm:"column:name"`
	Enabled bool   `gorm:"column:enabled"`
}

// TableName
func (*BasicPlatform) TableName() string {
	return "basic_platform"
}
