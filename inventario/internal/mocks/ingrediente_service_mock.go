package mocks

import (
	domain "cmd/api/internal/domain/entities"
	"context"

	"github.com/stretchr/testify/mock"
)

type IngredienteServiceMock struct {
	mock.Mock
}

func (m *IngredienteServiceMock) Get(ctx context.Context, id string) (domain.Ingrediente, error) {
	args := m.Called(id)
	return domain.Ingrediente{}, args.Error(1)

}

func (m *IngredienteServiceMock) Create(ctx context.Context, ingrediente domain.Ingrediente) (*string, error) {
	args := m.Called(ingrediente)
	return nil, args.Error(1)
}

func (m *IngredienteServiceMock) Update(ctx context.Context, id string, cantidad float64) error {
	args := m.Called(id, cantidad)
	return args.Error(0)
}

func (m *IngredienteServiceMock) Delete(ctx context.Context, id string) error {
	args := m.Called(id)
	return args.Error(1)
}
