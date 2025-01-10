package handlers

import (
	"log"

	"github.com/HenriqueSchroeder/golang-elk-kafka/pkg/contracts"
	"github.com/HenriqueSchroeder/golang-elk-kafka/pkg/kafka"
	"github.com/IBM/sarama"
)

type ConsumerHandler struct{}

func (ConsumerHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (ConsumerHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

func (h ConsumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		log.Printf("Mensagem recebida: topic=%s partition=%d offset=%d key=%s value=%s", message.Topic, message.Partition, message.Offset, string(message.Key), string(message.Value))
		session.MarkMessage(message, "")
	}
	return nil
}

func StartConsumers(consumer *kafka.Consumer) {
	topics := []string{
		contracts.ColorTopic,
		contracts.FamilyTopic,
		contracts.ProductTopic,
		contracts.CollectionTopic,
		contracts.ProductVariantTopic,
	}

	go func() {
		err := consumer.Consume(topics, ConsumerHandler{})
		if err != nil {
			log.Fatalf("Erro ao consumir mensagens: %v", err)
		}
	}()
}
