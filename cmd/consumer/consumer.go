package main

import (
	"log"

	"github.com/HenriqueSchroeder/golang-elk-kafka/config"
	"github.com/HenriqueSchroeder/golang-elk-kafka/internal/kafka"
	"github.com/HenriqueSchroeder/golang-elk-kafka/pkg/contracts"
	"github.com/HenriqueSchroeder/golang-elk-kafka/pkg/handlers"
)

func startColorConsumer(brokers []string) {
	consumer, err := kafka.NewConsumer("color-group", brokers)
	if err != nil {
		log.Fatalf("Falha ao iniciar o consumidor de cores: %v", err)
	}
	defer consumer.Close()

	err = consumer.Consume([]string{contracts.ColorTopic}, handlers.ColorConsumerHandler{})
	if err != nil {
		log.Fatalf("Erro ao consumir mensagens de cor: %v", err)
	}
}

func startCollectionConsumer(brokers []string) {
	consumer, err := kafka.NewConsumer("collection-group", brokers)
	if err != nil {
		log.Fatalf("Falha ao iniciar o consumidor de coleções: %v", err)
	}
	defer consumer.Close()

	err = consumer.Consume([]string{contracts.CollectionTopic}, handlers.CollectionConsumerHandler{})
	if err != nil {
		log.Fatalf("Erro ao consumir mensagens de coleção: %v", err)
	}
}

func main() {
	cfg := config.LoadConfig()
	brokers := []string{cfg.KafkaBrokers}

	go startColorConsumer(brokers)
	go startCollectionConsumer(brokers)

	log.Println("Consumidores Kafka iniciados com sucesso")
	select {}
}
