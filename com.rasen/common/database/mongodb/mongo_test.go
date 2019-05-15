package mongodb_test

import (
	"com.rasen/common/database/mongodb"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
	"time"
)

func TestConn(t *testing.T) {
	client := mongodb.NewMongo("mongodb://127.0.0.1:27017")
	collection := client.Database("test").Collection("t_doc")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res,err := collection.InsertOne(ctx,bson.M{"name":"pi","value":"3.14"})
	if err != nil{
		fmt.Println("err:",err)
	}
	fmt.Println("res:",res)
	resD,err := collection.DeleteOne(ctx,bson.M{"value":"3.14"})
	if err != nil{
		fmt.Println("del err:",err)
	}
	fmt.Println("del res:",resD)
}
