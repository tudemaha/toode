package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type Shortener struct {
	ID      primitive.ObjectID `bson:"_id"`
	Shorten string             `bson:"shorten"`
	URL     string             `bson:"url"`
}
