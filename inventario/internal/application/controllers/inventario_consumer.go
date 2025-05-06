package controllers

import (
	"cmd/api/internal/domain/services"
	"cmd/api/internal/infrastructure/messaging"
	"context"
	"encoding/json"
	"log"
)

type InventarioConsumer struct {
	Client     *messaging.RabbitMQClient
	service    services.InventarioService
	Exchange   string
	RoutingKey string
}

type Message struct {
	Id       string `json:"ingrediente_id"`
	Cantidad int    `json:"cantidad"`
}

func NewInventarioConsumerController(c *messaging.RabbitMQClient, s services.InventarioService, exchange, routingKey string) InventarioConsumer {
	return InventarioConsumer{
		Client:     c,
		service:    s,
		Exchange:   exchange,
		RoutingKey: routingKey,
	}
}

func (lc InventarioConsumer) Start() error {
	ch := lc.Client.GetChannel()

	// Declarar el exchange tipo topic
	err := ch.ExchangeDeclare(
		lc.Exchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	// Crear una cola temporal/exclusiva
	q, err := ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	// Enlazar la cola al exchange con una routing key
	err = ch.QueueBind(
		q.Name,
		lc.RoutingKey,
		lc.Exchange,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	log.Printf("Escuchando mensajes con routing key '%s' en exchange '%s'...", lc.RoutingKey, lc.Exchange)
	go func() {
		for d := range msgs {
			var payload Message
			err := json.Unmarshal(d.Body, &payload)
			if err != nil {
				log.Println("Error al parsear JSON:", err)
				continue
			}
			log.Println("Mensaje recibido:", payload)
			lc.service.UpdateInventario(context.Background(), payload.Id, float64(payload.Cantidad))
		}
	}()

	// select {}
	return nil
}
