package client

import "github.com/drasko/edgex-export/mongo"

var repo *mongo.MongoRepository

func InitMongoRepository(r *mongo.MongoRepository) {
	repo = r
	return
}
