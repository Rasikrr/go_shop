package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type CartItem struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"ID,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id" bson:"user_id"`
	ProductID string             `bson:"product_id" json:"ProductID"`
	Size      string             `bson:"size" json:"Size"`
	Amount    uint               `bson:"amount" json:"Amount"`
}
