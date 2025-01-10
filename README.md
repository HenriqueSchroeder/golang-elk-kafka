# Golang ELK Kafka Project

Este projeto é uma POC (Proof of Concept) que demonstra a integração entre Golang, Kafka (usando Kraft) e a stack ELK (Elasticsearch, Logstash, Kibana) para gerenciamento de logs.

## Pré-requisitos

- Docker e Docker Compose
- Go 1.19 ou superior
- Make

## Configuração

1. Clone o repositório:
```bash
git clone https://github.com/HenriqueSchroeder/golang-elk-kafka.git
cd golang-elk-kafka
```

2. Copie o arquivo de exemplo de ambiente:
```bash
cp .env.sample .env
```

3. Ajuste as variáveis no arquivo `.env` conforme necessário.

## Executando o Projeto

1. Primeiro, inicie a infraestrutura (Kafka, ELK stack):
```bash
make run
```

2. Aguarde alguns minutos para que todos os serviços estejam prontos. Você pode verificar o status com:
```bash
make status
```

3. Em um novo terminal, inicie a aplicação Go:
```bash
make run-app
```

A aplicação estará disponível em `http://localhost:7700`

## Endpoints Disponíveis

A API possui os seguintes endpoints para envio de mensagens:

- POST `/send/logs` - Envio de logs
- POST `/send/color` - Envio de cores
- POST `/send/family` - Envio de famílias
- POST `/send/product` - Envio de produtos
- POST `/send/collection` - Envio de coleções
- POST `/send/product/variant` - Envio de variantes de produtos

## Comandos Úteis

- `make status` - Verifica o status dos containers
- `make logs` - Exibe logs dos containers
- `make stop` - Para todos os containers
- `make clean` - Remove containers, volumes e redes
- `make restart` - Reinicia todos os containers
- `make test` - Executa os testes

## Visualização de Logs

Após a inicialização, você pode acessar:

- Kibana: http://localhost:5601
- Elasticsearch: http://localhost:9200

## Contribuindo

Sinta-se à vontade para abrir issues ou enviar pull requests com melhorias.
