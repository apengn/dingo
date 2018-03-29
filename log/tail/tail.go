package tail

import (
	"github.com/hpcloud/tail"
	"github.com/dingoblog/dingo/config"
	"github.com/astaxie/beego/logs"
	"fmt"
	"time"
	"github.com/dingoblog/dingo/log/kafka"
	"github.com/Shopify/sarama"
)

//追踪日志文件
func init() {
	tails, err := tail.TailFile(config.Conf.LogConf.FileName, tail.Config{
		ReOpen: true,
		Follow: true,
		//Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})
	if err != nil {
		logs.Error(err.Error())
		return
	}

	kafkaClient, err := kafka.ConnectKafkaProducer()
	if err != nil {
		logs.Error("kafka connect failed, err:", err)
		return
	}
	go acceptMsg(tails, kafkaClient)
}

func acceptMsg(tails *tail.Tail, kafkaClient sarama.SyncProducer) {
	for {
		var msg *tail.Line
		var ok bool
		//每次文件更新都会进行追踪
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Println("tail file close reopen ,filename", tails.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}

		pmsg := &sarama.ProducerMessage{}
		pmsg.Topic = "http_log"
		pmsg.Value = sarama.StringEncoder(msg.Text)
		_, _, err := kafkaClient.SendMessage(pmsg)
		//partition, offset, err := kafkaClient.SendMessage(pmsg)
		if err != nil {
			logs.Error(err.Error())
		}
		//fmt.Printf("partition:%v,offset:%v\n", partition, offset)
	}
}
