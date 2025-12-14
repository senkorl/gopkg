package main

import (
	"log"

	"github.com/IBM/sarama"
)

func main() {
	// 创建消费者配置
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// 创建消费者客户端
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config) // 替换成你的Kafka地址
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}
	defer consumer.Close()

	// 订阅Kafka的Topic
	partitionConsumer, err := consumer.ConsumePartition("your_topic", 0, sarama.OffsetNewest) // 替换成你要订阅的Topic和Partition
	if err != nil {
		log.Fatalf("Error starting partition consumer: %v", err)
	}
	defer partitionConsumer.Close()

	log.Printf("Waiting for messages")
	// 消费消息
	for msg := range partitionConsumer.Messages() {
		log.Printf("Received message: %s", string(msg.Value))
	}
}
