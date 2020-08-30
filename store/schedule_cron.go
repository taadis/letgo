package store

// ScheduleCron
type ScheduleCron struct {
	Id      string `gorm:"column:id;primary_key"`
	Name    string `gorm:"column:name"`
	Value   string `gorm:"column:value"`
	Enabled bool   `gorm:"column:enabled"`
}

// TableName
func (*ScheduleCron) TableName() string {
	return "schedule_cron"
}
