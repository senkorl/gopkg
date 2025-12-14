#!/bin/bash

BROKER=localhost:9092
TOPIC=adCallback
GROUP=testGroup
MESSAGE='{"order_id":1234,"status":"paid"}'

echo "=== 创建 topic ==="
docker exec -it kafka /opt/kafka/bin/kafka-topics.sh \
  --bootstrap-server $BROKER \
  --create \
  --topic $TOPIC \
  --partitions 1 \
  --replication-factor 1

echo "=== 列出 topic ==="
docker exec -it kafka /opt/kafka/bin/kafka-topics.sh \
  --bootstrap-server $BROKER \
  --list

echo "=== 生产消息 ==="
echo $MESSAGE | docker exec -i kafka /opt/kafka/bin/kafka-console-producer.sh \
  --bootstrap-server $BROKER \
  --topic $TOPIC

echo "=== 消费消息 ==="
docker exec -it kafka /opt/kafka/bin/kafka-console-consumer.sh \
  --bootstrap-server $BROKER \
  --topic $TOPIC \
  --from-beginning \
  --group $GROUP \
  --max-messages 1

echo "=== 查看消费者组 offset 和 lag ==="
docker exec -it kafka /opt/kafka/bin/kafka-consumer-groups.sh \
  --bootstrap-server $BROKER \
  --group $GROUP \
  --describe

