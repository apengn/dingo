package log_tranter

import (
	elastic "gopkg.in/olivere/elastic.v2"
)

type Tweet struct {
	Topic   string
	Message string
}

var (
	client *elastic.Client
)

func initEs() (err error) {
	client, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(LogConfig.EsAddr))
	if err != nil {
		return err
	}
	return nil
}

func sendDatatoEs(topic, message string) error {
	tweet := &Tweet{Topic: topic, Message: message}
	_, err := client.Index().
		Index(topic).
			Type(topic).
				//Id().
					BodyJson(tweet).
						Do()
	if err != nil {
		return err
	}
	return nil
}
