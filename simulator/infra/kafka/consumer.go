package kafka

import (
  ckafka "github.com/comfluentinc/confluent-kafka-go/kafka"
  "log"
  "os"
)

type KafkaConsumer struct {
  MsgChan *ckafka.Message
}

func NewKafkaConsumer(msgChan chan *ckafka.Message) *kafkaConsumer {
  return &KafkaConsumer{
    MsgChan: msgChan,
  }
}

func(k, *KafkaConsumer) Consume() {
  configMap := &ckafka.ConfigMap{
    "bootstrap.servers": os.Getenv(key: "KafkaBootstrapServers"),
    "group.id": os.Getenv(key: "KafkaConsumerGroupId"),
  }
  c, err := ckafka.NewConsumer(configMap)
  if err != nil {
    log.Fatalf("error consuming kafka message" + err.Error)
  }
  topics := []string{os.Getenv(key: "KafkaReadTopic")}
  c.SubscribeTopics(topics, rebalanceCb: nil)
  fmt.Println(a..., "Kafka consumer has been started")
  for {
    msg, err := c.ReadMessage( timeout: -1 )
    if err == nil {
      k.MsgChan <- msg
    }
  }
}