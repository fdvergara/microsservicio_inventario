package main

import (
	"cmd/api/internal/application/controllers"
	"cmd/api/internal/domain/services"
	"cmd/api/internal/infrastructure/messaging"
	"cmd/api/internal/infrastructure/persistence"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AppControllers struct {
	ingredienteController controllers.IngredienteController
	recetaController      controllers.RecetaController
	inventarioController  controllers.InventarioController
	consumer              controllers.InventarioConsumer
}

var appControllers *AppControllers

func build() error {

	ingredienteConnection, err := mongoClient("ingrediente", "inventario")

	if err != nil {
		log.Fatal("Error connect to database for ingrediente")
		return err
	}

	recetaConnection, err := mongoClient("receta", "inventario")

	if err != nil {
		log.Fatal("Error connect to database for receta")
		return err
	}

	repositoryIngrediente := persistence.NewRepository(*ingredienteConnection)
	repositoryReceta := persistence.NewRepository(*recetaConnection)
	ingredienteService := services.NewIngredienteService(repositoryIngrediente)
	recetaService := services.NewRecetaService(repositoryReceta)
	inventarioService := services.NewInventarioService(ingredienteService)

	consumer := controllers.NewInventarioConsumerController(messageinClient(), inventarioService, "inventario_topic", "inventario.#")

	appControllers = &AppControllers{
		ingredienteController: controllers.NewIngredienteController(ingredienteService),
		recetaController:      controllers.NewrecetaController(recetaService),
		inventarioController:  controllers.NewInventarioController(inventarioService),
		consumer:              consumer,
	}

	return nil
}

func mongoClient(collection string, database string) (*mongo.Collection, error) {
	uri := "mongodb://root:password@mongo:27017"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Error to connect database: %v", err)
		return nil, err
	}

	// Verificar si la conexión es exitosa
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("No se pudo hacer ping al servidor MongoDB:", err)
	}

	connection := client.Database(database).Collection(collection)
	return connection, nil
}

func messageinClient() *messaging.RabbitMQClient {
	rabbitClient, err := messaging.NewRabbitMQClient("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("Error conectando a RabbitMQ: %v", err)
	}
	return rabbitClient
}
