package letgo

// Option 定义配置选项函数
type Option func(*Options)

// Options 存储应用配置参数
type Options struct {
	Name       string
	Version    string
	ID         string
	ConfigPath string
	Env        string
	Port       int
	Config     interface{} // 项目的配置结构体
}

// WithName 设置应用名称
func WithName(name string) Option {
	return func(o *Options) {
		o.Name = name
	}
}

// WithVersion 设置应用版本
func WithVersion(version string) Option {
	return func(o *Options) {
		o.Version = version
	}
}

// WithID 设置应用 ID
func WithID(id string) Option {
	return func(o *Options) {
		o.ID = id
	}
}

// WithConfigPath 设置配置文件路径
func WithConfigPath(path string) Option {
	return func(o *Options) {
		o.ConfigPath = path
	}
}

// WithEnv 设置运行环境
func WithEnv(env string) Option {
	return func(o *Options) {
		o.Env = env
	}
}

// WithPort 设置服务端口
func WithPort(port int) Option {
	return func(o *Options) {
		o.Port = port
	}
}

// WithConfig 设置配置结构体
func WithConfig(config interface{}) Option {
	return func(o *Options) {
		o.Config = config
	}
}

// NewOptions 创建 Options 实例，应用默认值
func NewOptions(opts ...Option) *Options {
	o := &Options{
		Env:        "dev",                 // 默认环境
		Port:       8080,                  // 默认端口
		ConfigPath: "configs/config.yaml", // 默认配置文件路径
	}
	for _, opt := range opts {
		opt(o)
	}
	return o
}
