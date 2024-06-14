package models

type Category struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
}

type Subcategory struct {
	ID       string `bson:"_id"`
	Name     string `bson:"name"`
	Category string `bson:"category"`
}
