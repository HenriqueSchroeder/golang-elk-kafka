package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/HenriqueSchroeder/golang-elk-kafka/internal/kafka"
	"github.com/HenriqueSchroeder/golang-elk-kafka/pkg/contracts"
)

/**
 * Send a product variant.
 */
func SendProductVariantMessageHandler(producer *kafka.Producer) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			http.Error(response, "Método de requisição inválido", http.StatusMethodNotAllowed)
			return
		}

		var productVariantMessage contracts.ProductVariantMessage

		err := json.NewDecoder(request.Body).Decode(&productVariantMessage)

		if err != nil {
			http.Error(response, "Dados inválidos", http.StatusBadRequest)
			return
		}

		sendMessage(response, producer, contracts.ProductVariantTopic, productVariantMessage)
	}
}
