package models

type Brand struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
}
