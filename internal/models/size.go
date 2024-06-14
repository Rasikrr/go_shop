package models

type Size struct {
	Size   string `bson:"size"`
	Amount uint   `bson:"amount"`
}
