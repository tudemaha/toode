package mongodb

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Connection *mongo.Database

func StartConnection() {
	log.Println("INFO MongoConnection: starting MongoDB connection...")

	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority",
		os.Getenv("MONGO_USER"), os.Getenv("MONGO_PASS"), os.Getenv("MONGO_HOST"))

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatalf("ERROR MongoConnection fatal error: %v", err)
	}

	Connection = client.Database(os.Getenv("MONGO_DATABASE"))

	log.Println("INFO MongoConnection: connection establish successfully")
}
