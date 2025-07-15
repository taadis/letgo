package letgo

import (
	"context"
	"net/url"
	"os"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport"
)

// Option 定义配置选项函数
type Option func(*Options)

// Options 存储应用配置参数
// ref:https://github.com/go-kratos/kratos/blob/main/options.go
type Options struct {
	//
	id string
	//
	name string
	//
	version string
	//
	metadata map[string]string
	//
	endpoints []*url.URL

	//
	ctx context.Context
	//
	sigs []os.Signal

	//
	logger log.Logger
	//
	registrar registry.Registrar
	//
	registrarTimeout time.Duration
	//
	stopTimeout time.Duration
	//
	servers []transport.Server

	//koptions   []kratos.Option
	ConfigName string
	Env        string
	Port       int
	Config     interface{} // 项目的配置结构体
}

// WithID 设置应用 ID
func WithID(id string) Option {
	return func(o *Options) {
		o.id = id
	}
}

// WithName 设置应用名称
func WithName(name string) Option {
	return func(o *Options) {
		o.name = name
	}
}

// WithVersion 设置应用版本
func WithVersion(version string) Option {
	return func(o *Options) {
		o.version = version
	}
}

// WithEndpoint with service endpoint.
func WithEndpoint(endpoints ...*url.URL) Option {
	return func(o *Options) {
		o.endpoints = endpoints
	}
}

// WithContext with service context.
func WithContext(ctx context.Context) Option {
	return func(o *Options) {
		o.ctx = ctx
	}
}

// WithLogger with service logger.
func WithLogger(logger log.Logger) Option {
	return func(o *Options) {
		o.logger = logger
	}
}

// WithServer with transport servers.
func WithServer(servers ...transport.Server) Option {
	return func(o *Options) {
		o.servers = servers
	}
}

// WithSignal with exit signals.
func WithSignal(sigs ...os.Signal) Option {
	return func(o *Options) {
		o.sigs = sigs
	}
}

// WithRegistrar with service registry.
func WithRegistrar(r registry.Registrar) Option {
	return func(o *Options) {
		o.registrar = r
	}
}

// RegistrarTimeout with registrar timeout.
func RegistrarTimeout(t time.Duration) Option {
	return func(o *Options) {
		o.registrarTimeout = t
	}
}

// StopTimeout with app stop timeout.
func StopTimeout(t time.Duration) Option {
	return func(o *Options) {
		o.stopTimeout = t
	}
}

// WithConfigName 设置配置文件名称
// 比如：你的配置文件是config.yaml,这里name=config
func WithConfigName(name string) Option {
	return func(o *Options) {
		o.ConfigName = name
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

// NewOptions 创建 Options 实例，应用默认值
func NewOptions(opts ...Option) *Options {
	// todo:代码里先不使用默认值,需要的话加到命令行参数里,更贴合点
	// o := &Options{
	// 	Env:        "dev",                 // 默认环境
	// 	Port:       8080,                  // 默认端口
	// 	ConfigName: "configs/config.yaml", // 默认配置文件路径
	// }
	o := new(Options)
	for _, opt := range opts {
		opt(o)
	}
	return o
}

// LoadConfig 动态加载配置文件（支持环境隔离）
// 文件命名规则：
//   - 默认加载 config.yaml
//   - 指定环境时加载 config.{env}.yaml
//
// 使用示例：
//
//	opts := NewOptions(WithEnv("prod"))
//	opts.LoadConfig(&cfg)
// func (o *Options) LoadConfig(configStruct interface{}) error {
// 	// 1. 解析配置目录
// 	configDir := "configs"
// 	if o.ConfigPath != "" {
// 		configDir = o.ConfigPath
// 	}

// 	// 2. 动态生成文件名
// 	fileName := "config.yaml"
// 	if o.Env != "" {
// 		fileName = fmt.Sprintf("config-%s.yaml", o.Env)
// 	}

// 	// 3. 加载配置
// 	configPath := filepath.Join(configDir, fileName)
// 	c := config.New(
// 		config.WithSource(
// 			file.NewSource(configPath),
// 		),
// 	)

// 	if err := c.Load(); err != nil {
// 		return fmt.Errorf("failed to load config from %s: %v", configPath, err)
// 	}

// 	if err := c.Scan(configStruct); err != nil {
// 		return fmt.Errorf("failed to scan config: %v", err)
// 	}

// 	o.Config = configStruct
// 	return nil
// }
