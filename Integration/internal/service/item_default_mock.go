package service

import (
	"testing/integration/internal"

	"github.com/stretchr/testify/mock"
)

type ItemDefaultMock struct {
	mock.Mock
}

func NewItemDefaultMock() *ItemDefaultMock {
	return &ItemDefaultMock{}
}

func (item *ItemDefaultMock) FindById(id int) (i internal.Item, err error) {
	args := item.Called(id)
	return args.Get(0).(internal.Item), args.Error(1)

}
