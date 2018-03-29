package kafka

import (
	"fmt"
	"testing"
	"github.com/Shopify/sarama"
	"sync"
)

//读取kafka中的数据
func TestKafkaCusutmer(t *testing.T) {

	wg := sync.WaitGroup{}

	consumer, err := ConnectKafkaConsumer()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("connect kafka success!!")
	}
	defer consumer.Close()

	partList, err := consumer.Partitions("http_log")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for partition := range partList {
		pc, err := consumer.ConsumePartition("http_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer pc.AsyncClose()

		wg.Add(1)
		go func(consumer sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range pc.Messages() {
				fmt.Println(string(msg.Value))
			}
		}(pc)
	}

	wg.Wait()

	fmt.Println("Done consuming topic http_log")
}
