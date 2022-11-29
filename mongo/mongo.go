package cache

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Adapter struct {
	db            *mongo.Database
	collection    *mongo.Collection
	dataLifecycle int64
}
type dataModel struct {
	ID        primitive.ObjectID `bson:"_id"`
	Data      string             `bson:"data"`
	SearchKey string             `bson:"searchKey"`
}

func New(connectionString, dataBaseName, collectionName string, dataLifecycle int64) (*Adapter, error) {
	connectionContext := context.Background()
	db, collection, err := connectDataBaseAndCollection(connectionContext, connectionString, dataBaseName, collectionName)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		db:            db,
		collection:    collection,
		dataLifecycle: dataLifecycle,
	}, nil
}

func connectDataBaseAndCollection(connectionContext context.Context, connectionString, dbName, collectName string) (db *mongo.Database, collection *mongo.Collection, err error) {
	mongoClient, err := mongo.Connect(connectionContext, options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Error("mongoDB Connect failed" + err.Error())
		return db, collection, err
	}
	db = mongoClient.Database(dbName)
	collection = db.Collection(collectName)
	return
}

func (a *Adapter) Get(key string) string {
	mongoCursor, err := a.findDataByKey(key)

	if err != nil {
		log.Error("mongoDB find data failed for key: " + key + " error: " + err.Error())
		return ""
	}
	defer mongoCursor.Close(context.Background())

	if !mongoCursor.Next(context.Background()) {
		log.Error("mongoDB find data failed for key: " + key)
		return ""
	}
	res := &dataModel{}

	err = mongoCursor.Decode(res)

	if err != nil {
		log.Error("mongoDB decode data failed " + err.Error())
		return ""
	}
	if (time.Now().Unix()-res.ID.Timestamp().Unix())/60 > a.dataLifecycle {
		a.deleteExpiredData(key)
		return ""
	}
	return res.Data
}
func (a *Adapter) findDataByKey(key string) (*mongo.Cursor, error) {
	Context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return a.collection.Find(Context, bson.M{"searchKey": key})
}
func (a *Adapter) deleteExpiredData(key string) {
	Context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	a.collection.DeleteMany(Context, bson.M{"searchKey": key})
	//delete all: collection.DeleteMany(context.Background(), bson.M{})
}
func (a *Adapter) Set(SearchKey, Data string) {
	Context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	data := dataModel{
		ID:        primitive.NewObjectIDFromTimestamp(time.Now()),
		Data:      Data,
		SearchKey: SearchKey,
	}
	_, err := a.collection.InsertOne(Context, data)

	if err != nil {
		log.Error("mongoDB insert data failed " + err.Error())
	}

}
