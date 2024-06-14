package models

type Product struct {
	ID                string   `bson:"_id"`
	Name              string   `bson:"name"`
	Brand             string   `bson:"brand"`
	Description       string   `bson:"description"`
	Category          string   `bson:"category"`
	Subcategory       string   `bson:"subcategory"`
	Price             float64  `bson:"price"`
	DiscountInPercent float64  `bson:"discount_in_percent"`
	Sizes             []Size   `bson:"sizes"`
	PathsToPhotos     []string `bson:"paths_to_photos,omitempty"`
	Slug              string   `bson:"slug"`
}
