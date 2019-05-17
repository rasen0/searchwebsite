package database

import (
	"bufio"
	"com.rasen/common/database/mongodb"
	"com.rasen/common/structlog"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"os"
	"runtime"
	"sync"
	"time"
)

func NewMongoDB(addr string) *mongo.Client{
	mgoDB := mongodb.NewMongo(addr)
	return mgoDB
}
var MemFd *os.File

func init() {
	MemFd,_=os.OpenFile("memStats.txt",os.O_CREATE|os.O_WRONLY,0666)
}
func SaveDataMapToMgo(dataMap *sync.Map,collection *mongo.Collection){
	data := make([]interface{},0)
	saveD := func(key interface{}, value interface{}) bool{
		defer func(){
			if err := recover();err !=nil{
				structlog.Logger.WithFields(logrus.Fields{"err":err}).Error("save data to mgo fail")
			}
		}()
		m := bson.M{}
		m["title"] = key
		m["href"] = value
		data = append(data,m)
		dataMap.Delete(key)
		return true
	}
	memStats := &runtime.MemStats{}
	runtime.ReadMemStats(memStats)
	fdWriter:=bufio.NewWriter(MemFd)
	s := fmt.Sprintf("0||%v %+v %+v\n",memStats.HeapAlloc,memStats.NumGC,memStats.PauseTotalNs)
	io.WriteString(fdWriter,s)
	dataMap.Range(saveD)
	runtime.ReadMemStats(memStats)
	s = fmt.Sprintf("1||%v %+v %+v\n",memStats.HeapAlloc,memStats.NumGC,memStats.PauseTotalNs)
	io.WriteString(fdWriter,s)
	fdWriter.Flush()
	InsertManyMgo(data,collection)
}

func InsertManyMgo(data []interface{},collection *mongo.Collection){
	ctx,_ := context.WithTimeout(context.Background(),10 * time.Second)
	collection.InsertMany(ctx,data)// 存数据
}

func InsertOneMgo(data bson.D,collection *mongo.Collection){
	ctx,_ := context.WithTimeout(context.Background(),10 * time.Second)
	collection.InsertOne(ctx,data)
}
