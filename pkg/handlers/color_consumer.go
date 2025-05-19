package handlers

import (
	"encoding/json"
	"log"

	"github.com/HenriqueSchroeder/golang-elk-kafka/pkg/contracts"
	"github.com/IBM/sarama"
)

type ColorConsumerHandler struct{}

func (ColorConsumerHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (ColorConsumerHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

func (h ColorConsumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		var colorMessage contracts.ColorMessage
		err := json.Unmarshal(message.Value, &colorMessage)
		if err != nil {
			log.Printf("Erro ao decodificar mensagem de cor: %v", err)
			continue
		}

		log.Printf("Cor recebida: %+v", colorMessage)
		session.MarkMessage(message, "")
	}
	return nil
}
