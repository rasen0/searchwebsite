package main

import (
	"com.rasen/contentpool/nsq"
)

func main() {
	webConsumer := nsq.NewConsumer("analyse", "classifyWeb")
	nsqs := []string{"127.0.0.1:4150"}
	webConsumer.Set("nsqs", nsqs)
	nsqlookups := []string{"127.0.0.1:4160"}
	webConsumer.Set("lookup", nsqlookups)
	webConsumer.Set("concurrence", 15)

	webConsumer.Start()
	select {}
	//webConsumer.Stop()
	//fmt.Printf("stop:%v\n",<-webConsumer.StopChan)
}
