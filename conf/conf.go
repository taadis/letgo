package conf

// DatabaseConf
type DatabaseConf struct {
	// mysql
	DriverName string `json:"driverName"`

	// format "username:password@(ip:port)/database?charset=utf8&parseTime=True&loc=Local"
	DataSourceName string `json:"dataSourceName"`
}

// Database
func Database() *DatabaseConf {
	return &DatabaseConf{
		DriverName:     "mysql",
		DataSourceName: "root:root..@/letgo?charset=utf8&parseTime=True&loc=Local",
	}
}
