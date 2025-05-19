package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/HenriqueSchroeder/golang-elk-kafka/internal/kafka"
	"github.com/HenriqueSchroeder/golang-elk-kafka/pkg/contracts"
)

/**
 * Send a collection.
 */
func SendCollectionMessageHandler(producer *kafka.Producer) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			http.Error(response, "Método de requisição inválido", http.StatusMethodNotAllowed)
			return
		}

		var collectionMessage contracts.CollectionMessage

		err := json.NewDecoder(request.Body).Decode(&collectionMessage)

		if err != nil {
			http.Error(response, "Dados inválidos", http.StatusBadRequest)
			return
		}

		sendMessage(response, producer, contracts.CollectionTopic, collectionMessage)
	}
}
