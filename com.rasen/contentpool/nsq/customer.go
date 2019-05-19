package nsq

import (
	"fmt"
)

type Consumer struct {
	topic       string
	channel     string
	concurrence int
	nsqds       []string
	nsqLookups  []string
	*nsq.Consumer
}

func NewConsumer(channel, topic string) *Consumer {
	return &Consumer{
		topic:      topic,
		channel:    channel,
		nsqds:      make([]string, 0),
		nsqLookups: make([]string, 0),
	}
}

func (c *Consumer) Set(option string, value interface{}) {
	switch option {
	case "nsqs":
		v, ok := value.([]string)
		if !ok {
			fmt.Printf("set nsqs fail. err:value is not slice\n")
		}
		c.nsqds = v
	case "lookup":
		v, ok := value.([]string)
		if !ok {
			fmt.Printf("set nsqlookup fail. err:value is not slice\n")
		}
		c.nsqLookups = v
	case "concurrence":
		v, ok := value.(int)
		if !ok {
			fmt.Printf("set concurrence fail. err:value is not int\n")
		}
		c.concurrence = v
	default:
		fmt.Printf("option is't exist")
	}
}

func (c *Consumer) Start() {
	cfg := nsq.NewConfig()
	c.Consumer, _ = nsq.NewConsumer("classifyWeb", "analyse", cfg)
	c.Consumer.AddConcurrentHandlers(c, c.concurrence)
	if len(c.nsqds) > 0 {
		err := c.Consumer.ConnectToNSQDs(c.nsqds)
		if err != nil {
			fmt.Printf("conn nsq fail. err:%v\n", err)
		}
	}
	if len(c.nsqLookups) > 0 {
		err := c.Consumer.ConnectToNSQLookupds(c.nsqLookups)
		if err != nil {
			fmt.Printf("conn lookup fail. err:%v\n", err)
		}
	}

}

func (c *Consumer) HandleMessage(message *nsq.Message) error {
	fmt.Printf("nsq co:%s\n", message.Body)

	return nil //todo un..
}
