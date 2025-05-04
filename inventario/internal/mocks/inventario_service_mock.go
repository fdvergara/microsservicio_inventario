package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type InventarioServiceMock struct {
	mock.Mock
}

func (m *InventarioServiceMock) VerificarDisponibilidad(ctx context.Context, ingredientes map[string]float64) (bool, map[string]float64, error) {
	args := m.Called(ingredientes)
	return args.Bool(0), args.Get(1).(map[string]float64), args.Error(2)
}

func (m *InventarioServiceMock) UpdateInventario(ctx context.Context, id string, cantidad float64) error {
	args := m.Called(id, cantidad)
	return args.Error(0)
}
