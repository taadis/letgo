package letgo

import (
	"fmt"
	"os"
	"sort"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
)

// ConfigSource 定义配置源接口
type ConfigSource interface {
	Load() (config.Config, error)
	Priority() int // 配置源优先级（数字越大优先级越高）
}

// FileSource 本地文件配置源
type FileSource struct {
	path string
}

func NewFileSource(path string) *FileSource {
	return &FileSource{path: path}
}

func (s *FileSource) Load() (config.Config, error) {
	if _, err := os.Stat(s.path); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file %s not found", s.path)
	}
	return config.New(config.WithSource(file.NewSource(s.path))), nil
}

func (s *FileSource) Priority() int {
	return 1
}

// LoadConfig 支持多配置源加载
// sources: 配置源实例（按优先级排序）
// configStruct: 目标配置结构体指针
func LoadConfig(configStruct interface{}, sources ...ConfigSource) error {
	// 按优先级排序
	sort.Slice(sources, func(i, j int) bool {
		return sources[i].Priority() > sources[j].Priority()
	})

	// 依次加载配置源
	for _, source := range sources {
		c, err := source.Load()
		if err != nil {
			log.Errorf("failed to load config source: %v", err)
			continue
		}
		defer c.Close()

		if err := c.Scan(configStruct); err != nil {
			log.Errorf("failed to scan config: %v", err)
			continue
		}
	}

	return nil
}

func (a *App) loadLocalConfig(v any) error {
	//
	configPath := fmt.Sprintf("%s.%s", a.options.ConfigName, "yaml")
	if a.options.Env != "" {
		configPath = fmt.Sprintf("%s-%s.%s", a.options.ConfigName, a.options.Env, "yaml")
	}
	// e.g.
	// configPath = "config-dev.yaml"
	log.Infof("loading local config from %s", configPath)

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return fmt.Errorf("not found local config file=%s", configPath)
	}

	fileSource := file.NewSource(configPath)
	c := config.New(config.WithSource(fileSource))
	defer c.Close()

	if err := c.Load(); err != nil {
		return fmt.Errorf("failed to load local config: %v", err)
	}

	if err := c.Scan(&v); err != nil {
		return fmt.Errorf("failed to scan local config: %v", err)
	}
	log.Infof("Loaded local config from %s", configPath)

	return nil
}
