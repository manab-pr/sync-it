package datasource

import "go.mongodb.org/mongo-driver/mongo"


// MongoUserDatasource handles raw MongoDB operations
type MongoUserDatasource struct {
	collection *mongo.Collection
}


// NewMongoUserDatasource creates a new MongoUserDatasource instance
func NewMongoUserDatasource(collection *mongo.Collection) *MongoUserDatasource {
	return &MongoUserDatasource{
		collection: collection,
	}
}




