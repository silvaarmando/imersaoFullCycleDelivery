package kafka

import (
	ckafka "github.com/comfluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
)

func NewKafkaProducer() *ckafka.Producer {
  configMap := &ckafka.ConfigMap{
    "bootstrap.servers": os.Getenv(key: "KafkaBootstrapServers"),
  }
  p, err := ckafka.NewProducer(configMap)
  if err != nil {
    log.Println(err.Error())
  }
  return p
}

func Publish(msg string, topic string, producer *ckafka.Producer) error {
  message := &ckafka.Message{
    TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
    Value: []byte(msg)
  }
  err := producer.Producer(message, deliveryChan: nil)
  if err != nil {
    return err
  }
  return nil
}