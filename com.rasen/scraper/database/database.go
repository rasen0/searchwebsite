package database

import (
	"com.rasen/common/database/mongodb"
	"com.rasen/common/structlog"
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
	"time"
)

func NewMongoDB(addr string) *mongo.Client{
	mgoDB := mongodb.NewMongo(addr)
	return mgoDB
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
		return true
	}
	dataMap.Range(saveD)
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
