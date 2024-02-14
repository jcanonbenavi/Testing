package repository

import (
	"doubles/internal/city"

	"github.com/stretchr/testify/mock"
)

func NewMock() *Mock {
	return &Mock{
		FuncSaveCity: func(c *city.City) {},
	}
}

type Mock struct {
	mock.Mock
	FuncSaveCity func(c *city.City)
}

func (m *Mock) SaveCity(c *city.City) (err error) {
	// Call the method of the mock and get the output
	output := m.Called(c)
	// Call the function if it is not nil
	if m.FuncSaveCity != nil {
		m.FuncSaveCity(c)
	}
	// Error is on the first position of the output
	err = output.Error(0)
	return err
}
