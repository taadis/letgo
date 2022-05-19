package service

import (
	"fmt"
	"log"

	"github.com/taadis/letgo/app/di/store"
)

type Service struct {
	Store store.Store
}

func (s *Service) GetSome(id int) (int, error) {
	// 执行一些逻辑,比如参数验证
	if id <= 0 {
		return 0, fmt.Errorf("invalid id:%d", id)
	}

	// 使用依赖项store.Store的Get()方法从数据库获取结果
	result, err := s.Store.Get(id)
	if err != nil {
		return 0, err
	}

	// 执行一些逻辑,比如结果处理
	result++
	return result, nil
}

func (s *Service) CheckConnect() error {
	// 使用依赖项store.Store的Ping()方法从数据库获取结果
	err := s.Store.Ping()
	if err != nil {
		// 可记录错误日志等
		log.Printf("ping error:%+v", err)
		return err
	}

	return nil
}
