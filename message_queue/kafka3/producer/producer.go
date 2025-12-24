package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/IBM/sarama"
)

func main() {

	brokers := []string{"localhost:9092"}
	topic := "dahanghai_click_callback"

	// 创建Kafka生产者配置
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true // 配置生产者返回成功信息

	// 创建生产者客户端
	producer, err := sarama.NewSyncProducer(brokers, config) // 替换成你的Kafka地址
	if err != nil {
		log.Fatalf("Error creating producer: %v", err)
	}
	defer producer.Close()

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
