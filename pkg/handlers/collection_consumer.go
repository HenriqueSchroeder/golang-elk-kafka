package handlers

import (
	"encoding/json"
	"log"

	"github.com/HenriqueSchroeder/golang-elk-kafka/pkg/contracts"
	"github.com/IBM/sarama"
)

type CollectionConsumerHandler struct{}

func (CollectionConsumerHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (CollectionConsumerHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

func (h CollectionConsumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		var collectionMessage contracts.CollectionMessage
		err := json.Unmarshal(message.Value, &collectionMessage)
		if err != nil {
			log.Printf("Erro ao decodificar mensagem de coleção: %v", err)
			continue
		}

		log.Printf("Coleção recebida: %+v", collectionMessage)
		session.MarkMessage(message, "")
	}
	return nil
}
