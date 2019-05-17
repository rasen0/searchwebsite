package nsq

import "github.com/nsqio/go-nsq"

type Consumer struct{
	*nsq.Consumer
}

func(c *Consumer) HandleMessage(message *nsq.Message) error{

}
