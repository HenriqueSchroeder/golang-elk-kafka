package kafka

import (
	"context"
	"log"

	"github.com/IBM/sarama"
)

/**
 * Producer.
 */
type Producer struct {
	syncProducer sarama.SyncProducer
}

/**
 * Creates a new Kafka producer with the broken addresses provided.
 */
func NewProducer(brokers []string) (*Producer, error) {
	log.Printf("Tentando conectar aos brokers Kafka: %v", brokers)

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer(brokers, config)

	if err != nil {
		return nil, err
	}

	return &Producer{syncProducer: producer}, nil
}

/**
 * Send a message to the specified kafka topic.
 */
func (producer *Producer) SendMessage(topic, message string) (partition int32, offset int64, err error) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err = producer.syncProducer.SendMessage(msg)

	if err != nil {
		return -1, -1, err
	}

	return partition, offset, nil
}

/**
 * Close fecha o produtor Kafka.
 */
func (producer *Producer) Close() error {
	return producer.syncProducer.Close()
}

/**
 * Consumer.
 */
type Consumer struct {
	consumerGroup sarama.ConsumerGroup
}

/**
 * Creates a new Kafka consumer for the given group and brokers.
 */
func NewConsumer(group string, brokers []string) (*Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRoundRobin()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumerGroup, err := sarama.NewConsumerGroup(brokers, group, config)

	if err != nil {
		return nil, err
	}

	return &Consumer{consumerGroup: consumerGroup}, nil
}

/**
 * Consume messages from the specified Kafka topics.
 */
func (consumer *Consumer) Consume(topics []string, handler sarama.ConsumerGroupHandler) error {
	for {
		err := consumer.consumerGroup.Consume(context.Background(), topics, handler)
		if err != nil {
			return err
		}
	}
}

/**
 * Close fecha o consumidor Kafka.
 */
func (consumer *Consumer) Close() error {
	return consumer.consumerGroup.Close()
}
