package main

import (
	"log"

	"github.com/nsqio/go-nsq"
)

func main() {
	p, err := NewProducer("127.0.0.1:4150")
	if err != nil {
		log.Fatalf("NewProducer error:%v", err)
	}
	defer p.producer.Stop()

	err = p.producer.Publish("my_topic", []byte("该吃饭了"))
	if err != nil {
		log.Fatalf("producer.Publish error:%v", err)
	}
}

type producer struct {
	producer *nsq.Producer
}

func NewProducer(addr string) (*producer, error) {
	config := nsq.NewConfig()
	p, err := nsq.NewProducer(addr, config)
	if err != nil {
		return nil, err
	}

	return &producer{
		producer: p,
	}, nil
}

func (p *producer) publish(topic string, body []byte) error {
	err := p.producer.Publish(topic, body)
	if err != nil {
		return err
	}

	return nil
}
