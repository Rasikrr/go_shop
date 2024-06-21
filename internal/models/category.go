package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	SubcatsID []string           `bson:"subcats_id"`
}

type Subcategory struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `bson:"name"`
	CategoryID string             `bson:"category_id"`
}
