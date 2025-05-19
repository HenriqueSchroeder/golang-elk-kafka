package main

import (
	"log"
	"net/http"

	"github.com/HenriqueSchroeder/golang-elk-kafka/config"
	"github.com/HenriqueSchroeder/golang-elk-kafka/internal/kafka"
	"github.com/HenriqueSchroeder/golang-elk-kafka/pkg/handlers"
)

var producer *kafka.Producer

func main() {
	cfg := config.LoadConfig()

	var err error
	producer, err = kafka.NewProducer([]string{cfg.KafkaBrokers})

	if err != nil {
		log.Fatalf("Falha ao iniciar o produtor Kafka: %v", err)
	}

	defer producer.Close()

	http.HandleFunc("/send/logs", handlers.SendLogMessageHandler(producer))
	http.HandleFunc("/send/color", handlers.SendColorMessageHandler(producer))
	http.HandleFunc("/send/family", handlers.SendFamilyMessageHandler(producer))
	http.HandleFunc("/send/product", handlers.SendProductMessageHandler(producer))
	http.HandleFunc("/send/collection", handlers.SendCollectionMessageHandler(producer))
	http.HandleFunc("/send/product/variant", handlers.SendProductVariantMessageHandler(producer))

	log.Printf("Iniciando servidor na porta :%s", cfg.ServicePort)
	log.Println("Servidor iniciado com sucesso")
	log.Fatal(http.ListenAndServe(":"+cfg.ServicePort, nil))
}
