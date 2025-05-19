package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/HenriqueSchroeder/golang-elk-kafka/internal/kafka"
	"github.com/HenriqueSchroeder/golang-elk-kafka/pkg/contracts"
)

/**
 * Send a family.
 */
func SendFamilyMessageHandler(producer *kafka.Producer) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			http.Error(response, "Método de requisição inválido", http.StatusMethodNotAllowed)
			return
		}

		var familyMessage contracts.FamilyMessage

		err := json.NewDecoder(request.Body).Decode(&familyMessage)

		if err != nil {
			http.Error(response, "Dados inválidos", http.StatusBadRequest)
			return
		}

		sendMessage(response, producer, contracts.FamilyTopic, familyMessage)
	}
}
