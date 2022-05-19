package store

import "github.com/stretchr/testify/mock"

type MockStore struct {
	mock.Mock
}

func (s *MockStore) Get(id int) (int, error) {
	called := s.Called(id)
	return called.Get(0).(int), called.Error(1)
}

func (s *MockStore) Ping() error {
	return s.Called().Error(1)
}
