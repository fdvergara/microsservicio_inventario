package persistence

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Get(ctx context.Context, id string, result interface{}) error
	Create(ctx context.Context, structure interface{}) error
	Update(ctx context.Context, id string, update interface{}) error
	Delete(ctx context.Context, id string) error
}

type mongoDatabase struct {
	db mongo.Collection
}

func NewRepository(db mongo.Collection) Repository {
	return &mongoDatabase{db: db}
}

func (r *mongoDatabase) Get(ctx context.Context, id string, result interface{}) error {
	err := r.db.FindOne(ctx, bson.M{"id": id}).Decode(result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("Documento no encontrado.")
			return mongo.ErrNoDocuments
		}
		log.Println("Error to get into Database", err)
		return err
	}
	return nil
}

func (r *mongoDatabase) Create(ctx context.Context, structure interface{}) error {
	_, err := r.db.InsertOne(ctx, structure)
	if err != nil {
		log.Println("Error to save into Database", err)
		return err
	}
	return nil
}

func (r *mongoDatabase) Update(ctx context.Context, id string, update interface{}) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"id": id}, update)
	if err != nil {
		log.Println("Error to update into Database", err)
		return err
	}
	return nil
}

func (r *mongoDatabase) Delete(ctx context.Context, id string) error {
	_, err := r.db.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		log.Println("Error to remove into Database", err)
		return err
	}
	return nil
}
