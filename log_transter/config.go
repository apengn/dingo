package log_tranter

import (
	"github.com/astaxie/beego/config"
	"errors"
)

var (
	LogConfig *LogConf
)

type LogConf struct {
	KafkaAddr  string
	EsAddr     string
	KafkaTopic string
	LogLevel   string
	LogPath    string
}

func initConfig(adapterName, fileName string) (err error) {
	configer, err := config.NewConfig(adapterName, fileName)
	if err != nil {
		return
	}
	LogConfig = &LogConf{}

	logPath := configer.String("logs::filepath")
	if len(logPath) == 0 {
		err = errors.New("logs filepath length is 0")
		return
	}
	LogConfig.LogPath = logPath
	logLevel := configer.String("logs::level")
	if len(logLevel) == 0 {
		logLevel = "debug"
	}
	LogConfig.LogLevel = logLevel

	kafkaAddre := configer.String("kafka::address")
	if len(kafkaAddre) == 0 {
		err = errors.New("kafka addr is nil")
		return
	}
	LogConfig.KafkaAddr = kafkaAddre
	esAddr := configer.String("es::address")

	kafkaTopic := configer.String("kafka::topic")
	if len(kafkaTopic) == 0 {
		err = errors.New("kafka topic is  nil")
		return
	}
	LogConfig.KafkaTopic = kafkaTopic

	if len(esAddr) == 0 {
		err = errors.New("es addr is nil")
		return
	}
	LogConfig.EsAddr = esAddr

	return
}
