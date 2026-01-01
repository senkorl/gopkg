package main

import (
	"log"

	"github.com/IBM/sarama"
)

//	var brokers = []string{
//		"alikafka-post-cn-j4g3vu44s00h-1-vpc.alikafka.aliyuncs.com:9092",
//		"alikafka-post-cn-j4g3vu44s00h-2-vpc.alikafka.aliyuncs.com:9092",
//		"alikafka-post-cn-j4g3vu44s00h-3-vpc.alikafka.aliyuncs.com:9092",
//	}
var brokers = []string{"localhost:9092"}
var topic = "dahanghai_click_callback"

func main() {
	// 创建消费者配置
	config := sarama.NewConfig()
	config.Version = sarama.V3_9_0_0
	config.Consumer.Return.Errors = true
	//
	//config.Net.SASL.Enable = true
	//config.Net.SASL.Mechanism = sarama.SASLTypePlaintext
	//config.Net.SASL.User = "alikafka_post-cn-j4g3vu44s00h"
	//config.Net.SASL.Password = "PLfSoumK6AdybSaL8HDwxPuYZN9jQYqZ"
	//config.Net.SASL.Handshake = true

	// 创建消费者客户端
	consumer, err := sarama.NewConsumer(brokers, config) // 替换成你的Kafka地址
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}
	defer consumer.Close()

	log.Println("Starting consumer...")
	// 订阅Kafka的Topic
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest) // 替换成你要订阅的Topic和Partition
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
