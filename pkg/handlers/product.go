package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/HenriqueSchroeder/golang-elk-kafka/pkg/contracts"
	"github.com/HenriqueSchroeder/golang-elk-kafka/pkg/kafka"
)

/**
 * Send a product.
 */
func SendProductMessageHandler(producer *kafka.Producer) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			http.Error(response, "Método de requisição inválido", http.StatusMethodNotAllowed)
			return
		}

		var productMessage contracts.ProductMessage

		err := json.NewDecoder(request.Body).Decode(&productMessage)

		if err != nil {
			http.Error(response, "Dados inválidos", http.StatusBadRequest)
			return
		}

		sendMessage(response, producer, contracts.ProductTopic, productMessage)
	}
}
