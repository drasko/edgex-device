package mongo

import (
	"gopkg.in/mgo.v2"
)

const (
	DbName         string = "test"
	CollectionName string = "exclient"
)

type MongoRepository struct {
	Session *mgo.Session
}

func NewMongoRepository(ms *mgo.Session) *MongoRepository {
	return &MongoRepository{Session: ms}
}
