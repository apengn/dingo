package log_tranter

func run() (error) {
	setSendDataFunc(sendDatatoEs)
	return readKafkaData()
}
