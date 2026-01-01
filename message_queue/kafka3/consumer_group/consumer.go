package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
)

func main() {

	brokers := []string{"localhost:9092"}

	//brokers := []string{
	//	"alikafka-post-cn-j4g3vu44s00h-1-vpc.alikafka.aliyuncs.com:9092",
	//	"alikafka-post-cn-j4g3vu44s00h-2-vpc.alikafka.aliyuncs.com:9092",
	//	"alikafka-post-cn-j4g3vu44s00h-3-vpc.alikafka.aliyuncs.com:9092",
	//}

	topics := []string{"dahanghai_click_callback"}
	groupID := "dahanghai_click_callback"

	config := sarama.NewConfig()

	//assignor := "range"
	//switch assignor {
	//case "sticky":
	//	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategySticky()}
	//case "roundrobin":
	//	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}
	//case "range":
	//	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRange()}
	//default:
	//	log.Panicf("Unrecognized consumer group partition assignor: %s", assignor)
	//}
	//config.Version = sarama.V3_0_0_0
	config.Producer.Partitioner = sarama.NewHashPartitioner
	//config.Producer.Return.Successes = false //是否等待成功和失败后的响应
	//config.Producer.Return.Errors = true
	//config.Producer.Timeout = time.Duration(10) * time.Second

	config.Net.MaxOpenRequests = 100
	config.Net.DialTimeout = time.Duration(10) * time.Second
	config.Net.ReadTimeout = time.Duration(10) * time.Second
	config.Net.WriteTimeout = time.Duration(10) * time.Second

	group, _ := sarama.NewConsumerGroup(brokers, groupID, config)

	fmt.Println("Created consumer group: ", groupID)
	for {
		err := group.Consume(context.Background(), topics, Consumer{})
		if err != nil {
			log.Fatalf("Error consuming messages: %v", err)
		}
		time.Sleep(10 * time.Second)
	}

}

type Consumer struct{}

func (Consumer) Setup(sarama.ConsumerGroupSession) error {
	log.Println("consumer group setup")

	return nil
}
func (Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	log.Println("consumer group cleanup")

	return nil
}

func (Consumer) ConsumeClaim(
	session sarama.ConsumerGroupSession,
	claim sarama.ConsumerGroupClaim,
) error {
	for msg := range claim.Messages() {
		// 业务处理
		log.Printf(
			"topic=%s partition=%d offset=%d key=%s value=%s",
			msg.Topic,
			msg.Partition,
			msg.Offset,
			string(msg.Key),
			string(msg.Value),
		)
		session.MarkMessage(msg, "")
	}
	return nil
}
