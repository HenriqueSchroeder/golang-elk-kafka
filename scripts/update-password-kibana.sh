#!/bin/bash

if [ -z "$ELASTIC_USERNAME" ]; then
  echo "Erro: A variável ELASTIC_USERNAME não está definida!"
  exit 1
fi

if [ -z "$ELASTIC_PASSWORD" ]; then
  echo "Erro: A variável ELASTIC_PASSWORD não está definida!"
  exit 1
fi

if [ -z "$ELASTICSEARCH_USERNAME" ]; then
  echo "Erro: A variável ELASTICSEARCH_USERNAME não está definida!"
  exit 1
fi

if [ -z "$ELASTICSEARCH_PASSWORD" ]; then
  echo "Erro: A variável ELASTICSEARCH_PASSWORD não está definida!"
  exit 1
fi

# Aguardar o Elasticsearch ficar disponível
echo "Aguardando o Elasticsearch iniciar..."
until curl -s -u $ELASTIC_USERNAME:$ELASTIC_PASSWORD http://elasticsearch:9200/_cluster/health; do
  echo "Aguardando Elasticsearch iniciar..."
  sleep 5
done

echo "Elasticsearch está disponível!"

# Atualizar a senha da conta do usuário do kibana
echo "Atualizando a senha da conta do usuário $ELASTICSEARCH_USERNAME..."

curl -s -u $ELASTIC_USERNAME:$ELASTIC_PASSWORD -X POST "http://elasticsearch:9200/_security/user/$ELASTICSEARCH_USERNAME/_password" \
  -H 'Content-Type: application/json' \
  -d "{
    \"password\": \"$ELASTICSEARCH_PASSWORD\"
  }"

if [ $? -ne 0 ]; then
  echo "Erro ao atualizar a senha do usuário $ELASTICSEARCH_USERNAME. Verifique as permissões."
  exit 1
fi

# Inicia o Kibana
echo "Iniciando o Kibana..."
exec /usr/local/bin/kibana-docker
