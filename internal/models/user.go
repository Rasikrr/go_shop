package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	FirstName    string             `bson:"firstName" json:"FirstName"`
	LastName     string             `bson:"lastName" json:"LastName"`
	PasswordHash string             `bson:"passwordHash" json:"PasswordHash"`
	Email        string             `bson:"email" json:"Email"`
	Country      string             `bson:"country" json:"Country"`
	City         string             `bson:"city" json:"City"`
	Address      string             `bson:"address" json:"Address"`
	TelNumber    string             `bson:"telNumber" json:"TelNumber"`
}
