package mongodb

import (
	"context"
	"time"

	"com.rasen/common/structlog"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongo(address string) *mongo.Client{
	//client, err := mongo.NewClient(options.Client().ApplyURI(address))
	//if err != nil {
	//	structlog.Logger.WithFields(logrus.Fields{"err": err}).Error("new mongo db client fail")
	//}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//client.Connect(ctx)

	client, err := mongo.Connect(ctx,options.Client().ApplyURI(address))
	if err != nil{
		structlog.Logger.WithFields(logrus.Fields{"err":err}).Error("new mongo db client fail")
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		structlog.Logger.WithFields(logrus.Fields{"err": err}).Error("ping mongo fail")
	}
	return client
}
