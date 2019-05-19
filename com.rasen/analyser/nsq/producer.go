package nsq

import (
	"fmt"
	"github.com/nsqio/go-nsq"
)

type Producer struct {
	topic      string
	channel    string
	addr       string
	nsqds      []string
	nsqLookups []string
	*nsq.Producer
}

func NewProducer() *Producer {
	return &Producer{
		nsqds:      make([]string, 0),
		nsqLookups: make([]string, 0),
	}
}

func (p *Producer) Set(option string, value interface{}) {
	switch option {
	case "addr":
		v, ok := value.(string)
		if !ok {
			fmt.Printf("set addr fail. err:value is not string\n")
		}
		p.addr = v
	case "nsqs":
		v, ok := value.([]string)
		if !ok {
			fmt.Printf("set nsqs fail. err:value is not slice\n")
		}
		p.nsqds = v
	case "lookup":
		v, ok := value.([]string)
		if !ok {
			fmt.Printf("set nsqlookup fail. err:value is not slice\n")
		}
		p.nsqLookups = v
	default:
		fmt.Printf("option is't exist")
	}
}

func (p *Producer) Start() {
	cfg := nsq.NewConfig()
	var err error
	p.Producer, err = nsq.NewProducer(p.addr, cfg)
	if err != nil {
		fmt.Printf("new producer fail.err:%vn", err)
	}
}
