package kafka

import (
	"github.com/Shopify/sarama"
	"time"
)

//连接kafka消费者
func ConnectKafkaConsumer() (consumer sarama.Consumer, err error) {
	config := sarama.NewConfig()
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Consumer.Offsets.Initial = sarama.OffsetNewest //初始从最新的offset开始
	consumer, err = sarama.NewConsumer([]string{"127.0.0.1:9092"}, config)
	return
}

//连接kafka生产者
func ConnectKafkaProducer() (producer sarama.SyncProducer, err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	producer, err = sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	return
}
