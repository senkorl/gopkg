package main

import (
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	// 创建Kafka生产者配置
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true // 配置生产者返回成功信息

	// 创建生产者客户端
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config) // 替换成你的Kafka地址
	if err != nil {
		log.Fatalf("Error creating producer: %v", err)
	}
	defer producer.Close()

	// 发送消息到Kafka
	msg := &sarama.ProducerMessage{
		Topic: "your_topic", // 替换成你要发送的Topic
		Value: sarama.StringEncoder(fmt.Sprint("{\"order_id\": \"123\"}")),
	}

	tk := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-tk.C:
			// 发送消息
			_, _, err = producer.SendMessage(msg)
			if err != nil {
				log.Fatalf("Error sending message: %v", err)
			}
			log.Println("Message sent successfully")
		}
	}

}
