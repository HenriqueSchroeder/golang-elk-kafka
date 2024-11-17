#!/bin/bash

# Verifica se o diretório de dados do Kafka já está inicializado
if [ ! -f "/var/lib/kafka/data/meta.properties" ]; then
  echo "Cluster ID não encontrado. Gerando novo Cluster ID..."

  # Gera o Cluster ID dinamicamente
  CLUSTER_ID=$(kafka-storage random-uuid)

  # Exibe o Cluster ID gerado
  echo "Cluster ID gerado: $CLUSTER_ID"

  # Formata o armazenamento do Kafka com o Cluster ID gerado
  kafka-storage format -t "$CLUSTER_ID" -c /etc/kafka/kraft/server.properties
  
  echo "Kafka formatado com o Cluster ID: $CLUSTER_ID"
else
  # Se o Kafka já foi inicializado, extrai o CLUSTER_ID do arquivo meta.properties
  CLUSTER_ID=$(cat /var/lib/kafka/data/meta.properties | grep "cluster.id" | cut -d '=' -f2)
  echo "Usando Cluster ID existente: $CLUSTER_ID"
fi

# Exporta o CLUSTER_ID como variável de ambiente para o Kafka
export CLUSTER_ID

# Inicia o Kafka (usando o entrypoint original da imagem)
exec /etc/confluent/docker/run
