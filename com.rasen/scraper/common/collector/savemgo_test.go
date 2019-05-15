package collector

import (
	"com.rasen/scraper/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"testing"
)

var mgoClient *mongo.Client

func TestMain(m *testing.M){
	mgoClient = database.NewMongoDB("mongodb://127.0.0.1:27017")
	back := m.Run()
	os.Exit(back)
}

func TestSaveMgoD(t *testing.T) {
	d:=bson.D{bson.E{"title","lin"},
		bson.E{"href","www.dgd.com"}}

	database.InsertOneMgo()
}