package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HenriqueSchroeder/golang-elk-kafka/pkg/kafka"
)

/**
 * Send message to Kafka.
 */
func sendMessage(response http.ResponseWriter, producer *kafka.Producer, topic string, message interface{}) {
	messageBytes, err := json.Marshal(message)

	if err != nil {
		http.Error(response, fmt.Sprintf("Falha ao converter mensagem para JSON: %v", err), http.StatusInternalServerError)
		return
	}

	partition, offset, err := producer.SendMessage(topic, string(messageBytes))

	if err != nil {
		http.Error(response, fmt.Sprintf("Falha ao enviar mensagem: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(response, "Mensagem enviada para a partição %d no offset %d\n", partition, offset)
}
