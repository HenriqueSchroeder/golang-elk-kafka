package main

import (
	"log"

	"github.com/HenriqueSchroeder/golang-elk-kafka/pkg/config"
	"github.com/HenriqueSchroeder/golang-elk-kafka/pkg/handlers"
	"github.com/HenriqueSchroeder/golang-elk-kafka/pkg/kafka"
)

var consumer *kafka.Consumer

func main() {
	cfg := config.LoadConfig()

	var err error
	consumer, err = kafka.NewConsumer("my-group", []string{cfg.KafkaBrokers})
	if err != nil {
		log.Fatalf("Falha ao iniciar o consumidor Kafka: %v", err)
	}
	defer consumer.Close()

	handlers.StartConsumers(consumer)

	log.Println("Consumidor Kafka iniciado com sucesso")
	select {} // Keep the application running
}
