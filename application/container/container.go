package container

import (
	"api-customers/application/controllers"
	"api-customers/application/services"
	"api-customers/application/storage"
	"api-customers/application/storage/memory"
	database "api-customers/helpers"
)

func dbConfig() *database.MongoConfig {

	mongoUri := "mongodb://db-customers:27017/opentracing-customers?retryWrites=true&w=majority"

	return &database.MongoConfig{
		Uri:             mongoUri,
		User:            "mongo",
		Password:        "mongo",
		Database:        "opentracing-customers",
		ApplicationName: "open-telemetry-go",
		MinPoolSize:     1,
		MaxPoolSize:     5,
		AuthMechanism:   "SCRAM-SHA-1",
	}
}

func mongodbHelper() *database.MongodbHelper {
	return database.NewMongodbHelper(dbConfig())
}
func CustomerMongoRepository() storage.CustomerRepository {
	return storage.NewCustomerMongoRepository(mongodbHelper())
}
func CustomerInMemoryRepository() storage.CustomerRepository {
	return memory.NewCustomerInMemoryRepository()
}
func CustomerFinder() *services.CustomerFinder {
	return services.NewCustomerFinder(CustomerMongoRepository())
}

func CustomerController() *controllers.CustomerController {
	return controllers.NewCustomerController(CustomerFinder())
}
