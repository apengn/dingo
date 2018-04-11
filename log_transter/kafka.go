package log_tranter

import (
	"github.com/Shopify/sarama"
	"time"
	"fmt"
	"github.com/astaxie/beego/logs"
	"sync"
)

var (
	consumer sarama.Consumer
	sendData func(topic, message string) error
	wg       sync.WaitGroup
)

func initKafka() error {
	return ConnectKafkaConsumer()
}

//连接kafka消费者
func ConnectKafkaConsumer() (err error) {
	config := sarama.NewConfig()
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Consumer.Offsets.Initial = sarama.OffsetNewest //初始从最新的offset开始
	consumer, err = sarama.NewConsumer([]string{LogConfig.KafkaAddr}, config)
	return
}

func setSendDataFunc(f func(topic, message string) error) {
	sendData = f
}
func readKafkaData() error {
	partList, err := consumer.Partitions(LogConfig.KafkaTopic)
	if err != nil {
		logs.Error(err)
		return err
	}
	for partition := range partList {
		pc, err := consumer.ConsumePartition(LogConfig.KafkaTopic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer pc.AsyncClose()

		wg.Add(1)
		go func(consumer sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range pc.Messages() {
				fmt.Println(string(msg.Value))

				err = sendData(msg.Topic, string(msg.Value))
				if err != nil {
					logs.Warn("send es fail", err)
				}
			}

		}(pc)
	}
	wg.Wait()
	return err
}
