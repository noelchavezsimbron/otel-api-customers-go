package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
	"log"
	"time"
)


type MongoConfig struct {
	Uri             string
	User            string
	Password        string
	Database        string
	ApplicationName string
	MaxPoolSize     uint64
	MinPoolSize     uint64
	AuthMechanism   string
}

type MongodbHelper struct {
	client *mongo.Client
	conf   *MongoConfig
	db     *mongo.Database
}

func NewMongodbHelper(conf *MongoConfig) *MongodbHelper {
	return &MongodbHelper{conf: conf}
}

func (helper *MongodbHelper) DatabaseName() string {
	return helper.conf.Database
}

func (helper *MongodbHelper) Database(name string) *mongo.Database {

	return helper.client.Database(name)
}

func (helper *MongodbHelper) Collection(name string) *mongo.Collection {

	return helper.client.Database(helper.conf.Database).Collection(name)
}

func (helper *MongodbHelper) OpenConnection() error {

	if helper.client != nil {

		return nil
	}

	var err error
	var credentials *options.Credential

	cnOptions := options.Client().ApplyURI(helper.conf.Uri)

	if helper.conf.User != "" && helper.conf.Password != "" {
		credentials = &options.Credential{Username: helper.conf.User, Password: helper.conf.Password}

		cnOptions.Auth = credentials
		cnOptions.Auth.AuthMechanism = helper.conf.AuthMechanism
		cnOptions.Auth.AuthSource = helper.conf.Database
	}

	cnOptions.SetAppName(helper.conf.ApplicationName)
	cnOptions.SetMaxPoolSize(helper.conf.MaxPoolSize)
	cnOptions.SetMinPoolSize(helper.conf.MinPoolSize)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	cnOptions.Monitor = otelmongo.NewMonitor("jaeger-tracing-go-service")

	helper.client, err = mongo.Connect(ctx, cnOptions)
	if err != nil {
		defer func() {
			fmt.Println("close connection")
			if err := helper.closeConnection(); err != nil {
				log.Printf("%v", err)
			}
		}()
		return err
	}

	fmt.Println("trying to connect...")

	err = helper.client.Ping(context.TODO(), nil)
	if err != nil {
		defer func() {
			fmt.Println("close connection")
			if err := helper.closeConnection(); err != nil {
				log.Printf("%v", err)
			}
		}()
		log.Printf("%v", err)
		return err
	}

	fmt.Println("Connected..!")

	return nil
}

func (helper *MongodbHelper) closeConnection() error {

	err := helper.client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nConnection to MongoDB closed.")
	return err
}
