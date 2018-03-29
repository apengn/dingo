package kafka

import (
	"testing"
	"fmt"
	"github.com/Shopify/sarama"
)

func TestConnectKafka(t *testing.T) {
	client, err := ConnectKafkaProducer()
	if err != nil {
		fmt.Println(err)
	}
	defer client.Close()
	fmt.Println("kafka connect success")
	msg := &sarama.ProducerMessage{}
	msg.Topic = "topic_test"
	msg.Value = sarama.StringEncoder("this is a good test, my message is good")
	_, _, err = client.SendMessage(msg)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("kafka send success",msg.Value)
}
