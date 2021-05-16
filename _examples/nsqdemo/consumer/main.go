package main

import (
	"log"

	"github.com/nsqio/go-nsq"
)

type Consumer struct {

}

func (*Consumer) HandleMessage(message *nsq.Message) error {
	//var body string
	//err := json.Unmarshal(message.Body, &body)
	//if err != nil {
	//	log.Fatalf("json.Unmarshal error:%v", err)
	//	return err
	//}

	log.Printf("received message:%v", string(message.Body))

	message.Finish()
	return nil
}

func main()  {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("my_topic", "test", config)
	if err != nil {
		log.Fatalf("nsq.NewConsumer error:%v", err)
	}

	consumer.AddHandler(&Consumer{})

	err = consumer.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Fatalf("consumer.ConnectToNSQ error:%v", err)
	}

	select {

	}
}