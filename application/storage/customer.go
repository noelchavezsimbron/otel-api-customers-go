package storage

import (
	"api-customers/application/model"
	database "api-customers/helpers"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type customerMongoRepository struct {
	db *database.MongodbHelper
}

func NewCustomerMongoRepository(db *database.MongodbHelper) *customerMongoRepository {
	return &customerMongoRepository{db: db}
}

func (repo customerMongoRepository) FindAll(ctx context.Context) ([]model.Customer, error) {

	err := repo.db.OpenConnection()
	if err != nil {
		log.Fatal(err)
	}

	findOptions := options.Find()

	filter := bson.M{}
	cursor, err := repo.db.Collection("customers").Find(ctx, filter, findOptions)

	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	customers := make([]model.Customer, 0)

	for cursor.Next(ctx) {

		var customer model.Customer
		err := cursor.Decode(&customer)
		if err != nil {
			log.Fatal(err)
		}
		customers = append(customers, customer)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}
