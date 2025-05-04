package mocks

import (
	domain "cmd/api/internal/domain/entities"
	"context"

	"github.com/stretchr/testify/mock"
)

type RecetaServiceMock struct {
	mock.Mock
}

func (m *RecetaServiceMock) Create(ctx context.Context, receta domain.Receta) (*string, error) {
	args := m.Called(receta)
	return nil, args.Error(1)
}

func (m *RecetaServiceMock) Get(ctx context.Context, id string) (domain.Receta, error) {
	args := m.Called(id)
	return domain.Receta{}, args.Error(1)
}

func (m *RecetaServiceMock) AddIngrediente(ctx context.Context, id string, ingredienteCantidad domain.IngredienteCantidad) error {
	args := m.Called(id, ingredienteCantidad)
	return args.Error(0)
}

func (m *RecetaServiceMock) RemoveIngrediente(ctx context.Context, id string, ingrediente_id string) error {
	args := m.Called(id, ingrediente_id)
	return args.Error(1)
}
