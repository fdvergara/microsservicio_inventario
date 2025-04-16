package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type DatabaseMock struct {
	mock.Mock
}

func (m *DatabaseMock) Get(ctx context.Context, id string, result interface{}) error {
	args := m.Called(id, result)
	return args.Error(1)
}

func (m *DatabaseMock) Create(ctx context.Context, structure interface{}) error {
	args := m.Called(structure)
	return args.Error(1)
}

func (m *DatabaseMock) Update(ctx context.Context, id string, update interface{}) error {
	args := m.Called(id, update)
	return args.Error(1)
}

func (m *DatabaseMock) Delete(ctx context.Context, id string) error {
	args := m.Called(id)
	return args.Error(1)
}
