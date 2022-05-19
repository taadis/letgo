package service

import (
	"errors"
	"testing"

	"github.com/taadis/letgo/app/di/store"
)

func TestServiceSuccess(t *testing.T) {
	// create a new instance of the mock store
	mockStore := new(store.MockStore)
	// 在"On()"方法中,我们希望接下来"Get()"方法,并设置参数为2,并将返回参数定义为7,nil
	mockStore.On("Get", 2).Return(7, nil)
	// 接下来,创建一个新的Service实例,并使用mockStore作为"store.Store"接口的实际依赖项
	s := Service{mockStore}
	// 然后调用GetSome()方法
	_, err := s.GetSome(2)
	// 之前为mockStore的定义的期望在这里声明
	mockStore.AssertExpectations(t)
	// 最后,断言没有任何错误
	if err != nil {
		t.Errorf("error should be nil but got:%+v", err)
	}
}

func TestServiceParamsInvalid(t *testing.T) {
	mockStore := new(store.MockStore)
	// 本例中模拟参数为-1,这将导致服务中的验证失败
	mockStore.On("Get", -1).Return(5, nil)
	s := Service{mockStore}
	_, err := s.GetSome(-1)
	//mockStore.AssertExpectations(t)
	if err.Error() != "invalid id:-1" {
		t.Errorf("error should be 'invalid id:-1' but got %v", err)
	}
}

func TestStoreError(t *testing.T) {
	mockStore := new(store.MockStore)
	// 本例模拟store发生错误
	mockStore.On("Get", 2).Return(0, errors.New("some error"))
	s := Service{mockStore}
	_, err := s.GetSome(2)
	if err.Error() != "some error" {
		t.Errorf("error should be 'some error' but got %v", err)
	}
}
