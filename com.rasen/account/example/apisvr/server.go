package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"

	"com.rasen/protobuf/pbf/serverapi"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"google.golang.org/grpc"
)
type stringServer struct{}

func(svr *stringServer) ShowString(ctx context.Context,req *stringsvr.StringMsgReq)(*stringsvr.StringMsgResp,error){
	if ok:= strings.Contains(req.S,"error");ok{
		fmt.Println("error:","content contains error string")
		return &stringsvr.StringMsgResp{S:"",Err:"content contains error string"},errors.New("error string")
	}
	b := fmt.Sprintf("call ShowString method. show content: %s",req.S)
	fmt.Println(b)
	return &stringsvr.StringMsgResp{S:b},nil
}

func(svr *stringServer) ValidString(ctx context.Context,req *stringsvr.StringMsgReq)(*stringsvr.StringMsgResp,error){
	if ok := strings.ContainsRune(req.S,rune('错')); ok{
		return &stringsvr.StringMsgResp{S:"",Err:"content contains 错 string"},errors.New("error string")
	}
	b := fmt.Sprintf("call ValidString method. show content: %s",req.S)
	fmt.Println(b)
	return &stringsvr.StringMsgResp{S:b},nil
}

func main() {
	conn,err := net.Listen("tcp",":5000")
	if err != nil{
		fmt.Println(logger)
		level.Error(logger).Log("err: ",err)
	}
	level.Debug(logger).Log("err: ",err)
	ss := &stringServer{}
	server := grpc.NewServer()
	stringsvr.RegisterStringServerServer(server,ss)
	err = server.Serve(conn)
	if err != nil{
		level.Error(logger).Log("err:",err)
	}
	level.Debug(logger).Log("resp",)
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
