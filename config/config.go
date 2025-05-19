package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServicePort  string
	KafkaBrokers string
}

func LoadConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	servicePort := os.Getenv("SERVICE_PORT")
	kafkaBrokers := os.Getenv("KAFKA_BROKERS")

	log.Printf("Configuração carregada: SERVICE_PORT=%s, KAFKA_BROKERS=%s", servicePort, kafkaBrokers)

	return &Config{
		ServicePort:  servicePort,
		KafkaBrokers: kafkaBrokers,
	}
}
