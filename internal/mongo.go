package internal

import (
	"context"
	"fmt"
	"log"

	"github.com/shammalie/network_ip_bearer/pkg/ipapi"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	ipDataCollection   = "ip_data"
	ipEventsCollection = "ip_events"
)

type Db struct {
	client *mongo.Client
	ipData *mongo.Collection
	ipHits *mongo.Collection
}

func NewMongoClient() *Db {
	dbHost := viper.GetString("MONGO_HOST")
	dbUser := viper.GetString("MONGO_USERNAME")
	dbPass := viper.GetString("MONGO_PASSWORD")
	dbPort := viper.GetInt("MONGO_PORT")

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/?maxPoolSize=20&w=majority", dbUser, dbPass, dbHost, dbPort)

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	log.Println("MongoDB: successfully connected")

	return &Db{
		client: client,
		ipData: client.Database("dev").Collection(ipDataCollection),
		ipHits: client.Database("dev").Collection(ipEventsCollection),
	}
}

func handleId(id primitive.ObjectID) primitive.ObjectID {
	if id.IsZero() {
		return primitive.NewObjectID()
	}
	return id
}

func (db *Db) InsertIpDetail(data *ipapi.IpDetail) error {
	data.Id = handleId(data.Id)
	_, err := db.ipData.InsertOne(context.TODO(), *data)
	return err
}

func (db *Db) InsertIpEvent(data *Event) error {
	data.Id = handleId(data.Id)
	_, err := db.ipHits.InsertOne(context.TODO(), *data)
	return err
}

func (db *Db) GetIpDataById(id primitive.ObjectID) (*ipapi.IpDetail, error) {
	var doc *ipapi.IpDetail
	err := db.ipData.FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&doc)
	if err != nil {
		return &ipapi.IpDetail{}, err
	}
	return doc, nil
}

func (db *Db) GetIpDataByIp(ip string) (*ipapi.IpDetail, error) {
	var doc *ipapi.IpDetail
	err := db.ipData.FindOne(context.TODO(), bson.D{{"ip", ip}}).Decode(&doc)
	if err != nil {
		return &ipapi.IpDetail{}, err
	}
	return doc, nil
}

func (db *Db) UpdateLastSeen(id primitive.ObjectID, epoch int64) (*ipapi.IpDetail, error) {
	filter := bson.D{{"_id", id.String()}}
	update := bson.D{{"$set",
		bson.D{{"last_seen", epoch}},
	}}
	_, err := db.ipData.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return &ipapi.IpDetail{}, err
	}
	return db.GetIpDataById(id)
}
