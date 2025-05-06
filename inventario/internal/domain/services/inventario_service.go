package services

import (
	"context"
	"errors"
)

type InventarioService interface {
	VerificarDisponibilidad(ctx context.Context, ingredientes map[string]float64) (bool, map[string]float64, error)
	UpdateInventario(ctx context.Context, id string, cantidad float64) error
}

type inventarioService struct {
	ingrediente IngredienteService
}

func NewInventarioService(ingrediente IngredienteService) InventarioService {
	return &inventarioService{ingrediente: ingrediente}
}

func (s *inventarioService) VerificarDisponibilidad(ctx context.Context, ingredientes map[string]float64) (bool, map[string]float64, error) {
	detalle := make(map[string]float64)
	for id, cantidadRequerida := range ingredientes {
		ingrediente, err := s.ingrediente.Get(ctx, id)
		if err != nil {
			return false, nil, err
		}
		if ingrediente.Cantidad < cantidadRequerida {
			detalle[ingrediente.Nombre] = cantidadRequerida - ingrediente.Cantidad
		}
	}
	return len(detalle) == 0, detalle, nil
}

func (s *inventarioService) UpdateInventario(ctx context.Context, id string, cantidad float64) error {
	ingrediente, err := s.ingrediente.Get(ctx, id)
	if err != nil {
		return err
	}

	if ingrediente.Cantidad < cantidad {
		return errors.New("error_update")
	}

	err = s.ingrediente.Update(ctx, id, ingrediente.Cantidad-cantidad)
	return err
}
