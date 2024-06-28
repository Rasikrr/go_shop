package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go_shop/internal/models"
	"go_shop/internal/storage"
	"log"
	"os"
	"time"
)

type Storage struct {
	Client *mongo.Client
}

func ConnectDB() *Storage {
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")
	connURL := fmt.Sprintf("mongodb://%s:%s", host, port)

	clientOptions := options.Client().ApplyURI(connURL)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	storage := &Storage{Client: client}
	log.Println("Connected to MongoDB!")

	return storage
}

func (s *Storage) FillTestDate() {
	products := []interface{}{
		models.Product{
			ID:                "134",
			Name:              "Product 1",
			Brand:             "Maison Margiela",
			Description:       "Description 1",
			Category:          "Category 1",
			Subcategory:       "Subcategory 1",
			Price:             100.0,
			DiscountInPercent: 10.0,
			Sizes:             []models.Size{{"S", 10}, {"M", 20}, {"L", 30}, {"XL", 40}},
			PathsToPhotos:     []string{"path/to/photo1"},
			Slug:              "product-1",
		},
		models.Product{
			ID:                "2324",
			Name:              "Product 2",
			Brand:             "Zara",
			Description:       "Description 2",
			Category:          "Category 2",
			Subcategory:       "Subcategory 2",
			Price:             200.0,
			DiscountInPercent: 20.0,
			Sizes:             []models.Size{{"S", 10}, {"M", 20}, {"L", 30}, {"XL", 40}},
			PathsToPhotos:     []string{"path/to/photo2"},
			Slug:              "product-2",
		},
		models.Product{
			ID:                "344",
			Name:              "Product 3",
			Brand:             "Mango",
			Description:       "Description 3",
			Category:          "Category 3",
			Subcategory:       "Subcategory 3",
			Price:             300.0,
			DiscountInPercent: 30.0,
			Sizes:             []models.Size{{"S", 10}, {"M", 20}, {"L", 30}, {"XL", 40}},
			PathsToPhotos:     []string{"path/to/photo3"},
			Slug:              "product-3",
		},
		models.Product{
			ID:                "424",
			Name:              "Product 4",
			Brand:             "Acne studios",
			Description:       "Description 4",
			Category:          "Category 4",
			Subcategory:       "Subcategory 4",
			Price:             400.0,
			DiscountInPercent: 40.0,
			Sizes:             []models.Size{{"S", 10}, {"M", 20}, {"L", 30}, {"XL", 40}},
			PathsToPhotos:     []string{"path/to/photo4"},
			Slug:              "product-4",
		},
		models.Product{
			ID:                "544",
			Name:              "Product 5",
			Brand:             "Massimo Dutti",
			Description:       "Description 5",
			Category:          "Category 5",
			Subcategory:       "Subcategory 5",
			Price:             500.0,
			DiscountInPercent: 50.0,
			Sizes:             []models.Size{{"S", 10}, {"M", 20}, {"L", 30}, {"XL", 40}},
			PathsToPhotos:     []string{"path/to/photo5"},
			Slug:              "product-5",
		},
	}
	collection := s.Client.Database(storage.DB_NAME).Collection("products")
	_, err := collection.InsertMany(context.Background(), products)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Test data inserted into MongoDB!")

}

func (s *Storage) Close() {
	if err := s.Client.Disconnect(context.Background()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connection to MongoDB closed.")
}
