package repository

import "doubles/internal/city"

// Stub is a stub for the RepositoryWriter interface.
func NewStub() *Stub {
	return &Stub{}
}

type Stub struct {
	FuncSaveCity func(c *city.City) (err error)
}

func (s *Stub) SaveCity(c *city.City) (err error) {
	err = s.FuncSaveCity(c)
	return
}
