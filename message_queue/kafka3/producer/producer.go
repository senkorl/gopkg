package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	log.Println("Starting producer...")
	brokers := []string{"localhost:9092"}
	//topic := "dahanghai_click_callback"
	//
	//brokers := []string{
	//	"alikafka-post-cn-j4g3vu44s00h-1-vpc.alikafka.aliyuncs.com:9092",
	//	"alikafka-post-cn-j4g3vu44s00h-2-vpc.alikafka.aliyuncs.com:9092",
	//	"alikafka-post-cn-j4g3vu44s00h-3-vpc.alikafka.aliyuncs.com:9092",
	//}

	//brokers := []string{
	//	"alikafka-post-cn-j4g3vu44s00h-1.alikafka.aliyuncs.com:9093",
	//	"alikafka-post-cn-j4g3vu44s00h-2.alikafka.aliyuncs.com:9093",
	//	"alikafka-post-cn-j4g3vu44s00h-3.alikafka.aliyuncs.com:9093",
	//}

	topic := "dahanghai_click_callback"

	// 创建Kafka生产者配置
	config := sarama.NewConfig()

	//config.Version = sarama.V3_0_0_0
	//config.Producer.Partitioner = sarama.NewHashPartitioner
	config.Producer.Return.Successes = true //是否等待成功和失败后的响应
	//config.Producer.Return.Errors = true
	//config.Producer.Timeout = time.Duration(10) * time.Second
	//
	//config.Net.MaxOpenRequests = 100
	//config.Net.DialTimeout = time.Duration(10) * time.Second
	//config.Net.ReadTimeout = time.Duration(10) * time.Second
	//config.Net.WriteTimeout = time.Duration(10) * time.Second

	//config.Net.SASL.Enable = true
	//config.Net.SASL.Mechanism = sarama.SASLTypePlaintext
	//config.Net.SASL.User = "alikafka_post-cn-j4g3vu44s00h"
	//config.Net.SASL.Password = "PLfSoumK6AdybSaL8HDwxPuYZN9jQYqZ"
	//config.Net.SASL.Handshake = true

	// 创建生产者客户端
	producer, err := sarama.NewSyncProducer(brokers, config) // 替换成你的Kafka地址
	if err != nil {
		log.Fatalf("Error creating producer: %v", err)
	}
	defer producer.Close()

	log.Println("Producer is ready")

	type OrderInfo struct {
		OrderID  string
		SkuName  string
		Quantity int64
		Price    float64
	}

	tk := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-tk.C:
			orderID := "123"
			orderInfo := OrderInfo{
				OrderID:  orderID,
				SkuName:  "小面包",
				Quantity: time.Now().Unix(),
				Price:    1.99,
			}
			orderInfoStr, _ := json.Marshal(orderInfo)
			// 发送消息到Kafka
			msg := &sarama.ProducerMessage{
				Topic:     topic, // 替换成你要发送的Topic
				Partition: 0,
				Key:       sarama.StringEncoder(orderID),
				Value:     sarama.StringEncoder(orderInfoStr),
			}
			// 发送消息
			_, _, err = producer.SendMessage(msg)
			if err != nil {
				log.Fatalf("Error sending message: %v", err)
			}
			log.Println("Message sent successfully")
		}
	}

}
