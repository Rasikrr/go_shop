package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go_shop/internal/models"
	mn "go_shop/internal/storage/mongo"
	"log"
	"time"
)

type ProductRepo interface {
	GetAllSizes() ([]*models.Size, error)
	GetAllCategories() ([]*models.Category, error)
	GetAllBrands() ([]*models.Brand, error)
	GetProducts(d bson.D) ([]*models.Product, error)
	GetProductBySlug(slug string) (*models.Product, error)
}

type ProductRepoImpl struct {
	storage *mn.Storage
}

func NewProductRepo(db *mn.Storage) *ProductRepoImpl {
	return &ProductRepoImpl{
		storage: db,
	}
}

func (p *ProductRepoImpl) GetProducts(filters bson.D) ([]*models.Product, error) {
	collection := p.storage.Client.Database("go_shop").Collection("products")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	cursor, err := collection.Find(ctx, filters)
	if err != nil {
		log.Printf("Failed to get products: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []*models.Product

	if err := cursor.All(ctx, &products); err != nil {
		log.Printf("Failed to decode product: %v", err)
		return nil, err
	}
	return products, nil
}

func (p *ProductRepoImpl) GetProductBySlug(slug string) (*models.Product, error) {
	collection := p.storage.Client.Database("go_shop").Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	var product *models.Product
	err := collection.FindOne(ctx, bson.D{{"slug", slug}}).Decode(&product)
	if err != nil {
		log.Printf("Failed to get product: %v", err)
		return nil, err
	}
	return product, nil
}

func (p *ProductRepoImpl) GetAllSizes() ([]*models.Size, error) {
	collection := p.storage.Client.Database("go_shop").Collection("sizes")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Printf("Failed to get products: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var sizes []*models.Size

	if err := cursor.All(ctx, &sizes); err != nil {
		log.Printf("Failed to decode product: %v", err)
		return nil, err
	}

	return sizes, nil
}

func (p *ProductRepoImpl) GetAllCategories() ([]*models.Category, error) {
	collection := p.storage.Client.Database("go_shop").Collection("categories")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Printf("Failed to get category: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var categories []*models.Category

	if err := cursor.All(ctx, &categories); err != nil {
		log.Printf("Failed to decode product: %v", err)
		return nil, err

	}
	return categories, nil
}

func (p *ProductRepoImpl) GetAllBrands() ([]*models.Brand, error) {
	collection := p.storage.Client.Database("go_shop").Collection("brands")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Printf("Failed to get brands: %v", err)
		return nil, err
	}
	var brands []*models.Brand

	if err := cursor.All(ctx, &brands); err != nil {
		log.Printf("Failed to decode product: %v", err)
		return nil, err
	}
	return brands, nil
}
