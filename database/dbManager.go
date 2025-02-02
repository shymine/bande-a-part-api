package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DbObject struct {
	client *mongo.Client
	ctx context.Context
	dbName string
	collections map[string]*mongo.Collection
}

func (db *DbObject) GetContext() context.Context {
	return db.ctx
}
func (db *DbObject) RegisterCollection(collName string) {
	db.collections[collName] = db.client.Database(db.dbName).Collection(collName)
}
func (db *DbObject) GetCollection(collName string) *mongo.Collection {
	return db.collections[collName]
}

func (db *DbObject) Disconnect() {
	log.Println("Disconnecting Database")
	if err := db.client.Disconnect(db.ctx); err != nil {
		panic(err)
	}
}

func SetDBManager(dbName string, collections []string) {
	var db = DbObject{
		ctx: context.TODO(),
		dbName: dbName,
		collections: map[string]*mongo.Collection{},
	}
	
	clientOptions := options.Client().ApplyURI("mongodb://bande-a-part-db:27017/")
	client, err := mongo.Connect(db.GetContext(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(db.GetContext(), nil)
	if err != nil {
		log.Fatal(err)
	}
	db.client = client

	for _, c := range collections {
		db.RegisterCollection(c)
	}

	DB_MANAGER = db
}

var DB_MANAGER DbObject