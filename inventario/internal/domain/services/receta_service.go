package services

import (
	domain "cmd/api/internal/domain/entities"
	"cmd/api/internal/infrastructure/persistence"
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RecetaService interface {
	Create(ctx context.Context, receta domain.Receta) (*string, error)
	Get(ctx context.Context, id string) (domain.Receta, error)
	AddIngrediente(ctx context.Context, id string, ingredienteCantidad domain.IngredienteCantidad) error
	RemoveIngrediente(ctx context.Context, id string, ingrediente_id string) error
}

type recetaService struct {
	repository persistence.Repository
}

func NewRecetaService(repository persistence.Repository) RecetaService {
	return &recetaService{
		repository: repository,
	}
}

func (i *recetaService) Create(ctx context.Context, receta domain.Receta) (*string, error) {
	receta.Id = uuid.New().String()
	err := i.repository.Create(ctx, receta)
	if err != nil {
		log.Println("Error to save receta", err)
		return nil, errors.New("error_saved")
	}

	return &receta.Id, nil
}

func (i *recetaService) Get(ctx context.Context, id string) (domain.Receta, error) {
	var receta domain.Receta
	err := i.repository.Get(ctx, id, &receta)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.Receta{}, errors.New("not_found")
		}
		log.Println("Error to get ingrediente")
		return domain.Receta{}, errors.New("error_get")
	}
	return receta, nil
}

func (i *recetaService) AddIngrediente(ctx context.Context, id string, ingredienteCantidad domain.IngredienteCantidad) error {
	_, err := i.Get(ctx, id)
	if err != nil {
		return err
	}
	update := bson.M{
		"$push": bson.M{
			"ingredientes": ingredienteCantidad,
		},
	}
	err = i.repository.Update(ctx, id, update)
	if err != nil {
		log.Println("Error to update ingrediente", err)
		return errors.New("error_update")
	}
	return nil
}

func (i *recetaService) RemoveIngrediente(ctx context.Context, id string, ingrediente_id string) error {
	_, err := i.Get(ctx, id)
	if err != nil {
		return err
	}
	update := bson.M{
		"$pull": bson.M{
			"ingredientes": bson.M{
				"ingrediente_id": ingrediente_id,
			},
		},
	}

	err = i.repository.Update(ctx, id, update)
	if err != nil {
		log.Println("Error to remove ingrediente from receta", err)
		return errors.New("error_update")
	}
	return nil
}
