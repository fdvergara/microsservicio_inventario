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

type IngredienteService interface {
	Create(ctx context.Context, ingrediente domain.Ingrediente) (*string, error)
	Get(ctx context.Context, id string) (domain.Ingrediente, error)
	Update(ctx context.Context, id string, cantidad float64) error
	Delete(ctx context.Context, id string) error
}

type ingredienteService struct {
	repository persistence.Repository
}

func NewIngredienteService(repository persistence.Repository) IngredienteService {
	return &ingredienteService{
		repository: repository,
	}
}

func (i *ingredienteService) Get(ctx context.Context, id string) (domain.Ingrediente, error) {
	var ingrediente domain.Ingrediente
	err := i.repository.Get(ctx, id, &ingrediente)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.Ingrediente{}, errors.New("not_found")
		}
		log.Println("Error to get ingrediente")
		return domain.Ingrediente{}, errors.New("error_get")
	}
	return ingrediente, nil
}

func (i *ingredienteService) Create(ctx context.Context, ingrediente domain.Ingrediente) (*string, error) {
	ingrediente.Id = uuid.New().String()
	err := i.repository.Create(ctx, ingrediente)
	if err != nil {
		log.Println("Error to save ingrediente", err)
		return nil, errors.New("error_saved")
	}

	return &ingrediente.Id, nil
}

func (i *ingredienteService) Update(ctx context.Context, id string, cantidad float64) error {
	_, err := i.Get(ctx, id)
	if err != nil {
		return err
	}
	update := bson.M{
		"$set": bson.M{
			"cantidad": cantidad,
		},
	}

	err = i.repository.Update(ctx, id, update)
	if err != nil {
		log.Println("Error to update ingrediente", err)
		return errors.New("error_update")
	}
	return nil
}

func (i *ingredienteService) Delete(ctx context.Context, id string) error {
	return nil
}
