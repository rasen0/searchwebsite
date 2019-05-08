package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"com.rasen/protobuf/pbf/serverapi"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"google.golang.org/grpc"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts,grpc.WithInsecure())
	conn,err := grpc.Dial("127.0.0.1:5000",opts...)
	if err != nil{
		fmt.Println("cn:",err)
		level.Error(logger).Log("err: ",err)
	}
	defer conn.Close()
	client := stringsvr.NewStringServerClient(conn)
	resp,err := client.ShowString(context.Background(),&stringsvr.StringMsgReq{S:"show string client ctx error"})
	if err != nil{
		level.Error(logger).Log(err)
	}
	level.Debug(logger).Log("resp:",resp)
	fmt.Println("resp:",resp," err:",err)
}

var logger log.Logger

func init(){
	os.Mkdir("logs",os.ModePerm)
	logPath := filepath.Join(".","logs","log.txt")
	fd,err := os.OpenFile(logPath,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0666)
	defer fd.Close()
	if err !=nil{
		fmt.Fprint(os.Stderr,"open or create log file fail. err:",err)
	}
	logger = log.NewLogfmtLogger(fd)
	baseTime := time.Now()
	mockTime := func() time.Time {
		baseTime = baseTime.Add(time.Second)
		return baseTime
	}

	logger = log.With(logger, "time", log.Timestamp(mockTime), "caller", log.DefaultCaller)
}