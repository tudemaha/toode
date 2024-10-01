package service

import (
	"context"
	"log"

	"github.com/tudemaha/toode/internal/shortener/dto"
	pkg "github.com/tudemaha/toode/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

func GetRedirection(shortUrl string) (string, error) {
	connection := pkg.Connection
	coll := connection.Collection("redirection")

	var shortener dto.Shortener

	filter := bson.M{"shorten": shortUrl}
	if err := coll.FindOne(context.TODO(), filter).Decode(&shortener); err != nil {
		log.Fatalf("ERROR: GetRedirection fatal error: %v", err)
		return "", err
	}

	return shortener.URL, nil
}
