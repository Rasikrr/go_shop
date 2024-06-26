package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	FirstName    string             `bson:"firstName" json:"FirstName"`
	LastName     string             `bson:"lastName" json:"LastName"`
	PasswordHash string             `bson:"passwordHash" json:"PasswordHash"`
	PhotoPath    string             `bson:"photoPath" json:"PhotoPath"`
	Email        string             `bson:"email" json:"Email"`
	Country      string             `bson:"country" json:"Country"`
	State        string             `bson:"state" json:"State"`
	City         string             `bson:"city" json:"City"`
	Address      string             `bson:"address" json:"Address"`
	PostCode     string             `bson:"postCode" json:"PostCode"`
	TelNumber    string             `bson:"telNumber" json:"TelNumber"`
}
