package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/HenriqueSchroeder/golang-elk-kafka/pkg/contracts"
	"github.com/HenriqueSchroeder/golang-elk-kafka/pkg/kafka"
)

/**
 * Send a log.
 */
func SendLogMessageHandler(producer *kafka.Producer) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			fmt.Printf("Método de requisição inválido: %s\n", request.Method)

			http.Error(response, "Método de requisição inválido", http.StatusMethodNotAllowed)
			return
		}

		message := request.FormValue("message")

		if message == "" {
			http.Error(response, "Mensagem é obrigatória", http.StatusBadRequest)
			return
		}

		logMessage := contracts.LogMessage{
			Timestamp: time.Now().Format(time.RFC3339),
			Level:     "INFO",
			Message:   message,
		}

		sendMessage(response, producer, contracts.LogsTopic, logMessage)
	}
}
