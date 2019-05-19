package main

import (
	"bufio"
	"com.rasen/analyser/nsq"
	"fmt"
	"os"
)

func main() {
	webProducer := nsq.NewProducer()
	webProducer.Set("addr", "127.0.0.1:4150")
	webProducer.Start()
	for {
		reader := bufio.NewReader(os.Stdin)
		line, _, _ := reader.ReadLine()
		if string(line) == "exit" {
			fmt.Println("exit")
			return
		}
		err := webProducer.Publish("classifyWeb", line)
		if err != nil {
			fmt.Printf("pulish err:%v\n", err)
		}
	}
}
