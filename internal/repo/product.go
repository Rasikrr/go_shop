package repo

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go_shop/internal/models"
	"go_shop/internal/storage"
	mn "go_shop/internal/storage/mongo"
	"log"
)

var (
	ErrNoDocuments = errors.New("no documents found")
	ErrOutOfStock  = errors.New("product ouf of stock")
	ErrNoSize      = errors.New("size does not found")
)

type ProductRepo interface {
	GetAllSizes() ([]*models.Size, error)
	GetAllCategories() ([]*models.Category, error)
	GetAllBrands() ([]*models.Brand, error)
	GetProducts(d bson.D) ([]*models.Product, error)
	GetProductBySlug(slug string) (*models.Product, error)
	GetSubCatById(id primitive.ObjectID) (*models.Subcategory, error)
	GetTopSales(limit int) ([]*models.Product, error)
	GetRelatedProducts(category, subcategory string) ([]*models.Product, error)
}

type ProductRepoImpl struct {
	storage *mn.Storage
}

func NewProductRepo(db *mn.Storage) *ProductRepoImpl {
	return &ProductRepoImpl{
		storage: db,
	}
}

func (p *ProductRepoImpl) GetRelatedProducts(category, subcategory string) ([]*models.Product, error) {
	collection := p.storage.Client.Database(storage.DB_NAME).Collection(storage.PRODUCTS)
	ctx, cancel := context.WithTimeout(context.Background(), storage.TIMEOUT)

	defer cancel()

	filter := bson.D{{"$or", []bson.D{
		{{"category", category}},
		{{"subcategory", subcategory}},
	}}}
	opts := options.Find().SetLimit(4)

	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		log.Println("failed to get related products", err)
		return nil, err
	}

	var products []*models.Product

	if err := cursor.All(ctx, &products); err != nil {
		log.Println("failed to decode related products", err)
		return nil, err
	}
	return products, nil
}

func (p *ProductRepoImpl) GetProducts(filters bson.D) ([]*models.Product, error) {
	collection := p.storage.Client.Database(storage.DB_NAME).Collection(storage.PRODUCTS)

	ctx, cancel := context.WithTimeout(context.Background(), storage.TIMEOUT)

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

func (p *ProductRepoImpl) GetTopSales(limit int) ([]*models.Product, error) {
	collection := p.storage.Client.Database(storage.DB_NAME).Collection(storage.PRODUCTS)
	ctx, cancel := context.WithTimeout(context.Background(), storage.TIMEOUT)

	defer cancel()

	opt := options.Find()
	opt.SetSort(bson.D{{"discount_in_percent", -1}})
	if limit != -1 {
		opt.SetLimit(int64(limit))
	}

	cur, err := collection.Find(ctx, bson.D{{}}, opt)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return []*models.Product{}, nil
		}
		return nil, err
	}
	var products []*models.Product
	if err := cur.All(ctx, &products); err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductRepoImpl) GetProductBySlug(slug string) (*models.Product, error) {
	collection := p.storage.Client.Database(storage.DB_NAME).Collection(storage.PRODUCTS)
	ctx, cancel := context.WithTimeout(context.Background(), storage.TIMEOUT)

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
	collection := p.storage.Client.Database(storage.DB_NAME).Collection(storage.SIZES)
	ctx, cancel := context.WithTimeout(context.Background(), storage.TIMEOUT)

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
	collection := p.storage.Client.Database(storage.DB_NAME).Collection(storage.CATEGORIES)
	ctx, cancel := context.WithTimeout(context.Background(), storage.TIMEOUT)

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
	collection := p.storage.Client.Database(storage.DB_NAME).Collection(storage.BRANDS)
	ctx, cancel := context.WithTimeout(context.Background(), storage.TIMEOUT)

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

func (p *ProductRepoImpl) GetSubCatById(id primitive.ObjectID) (*models.Subcategory, error) {
	collection := p.storage.Client.Database(storage.DB_NAME).Collection(storage.SUBCATEGORIES)
	ctx, cancel := context.WithTimeout(context.Background(), storage.TIMEOUT)

	defer cancel()
	subCat := new(models.Subcategory)
	res := collection.FindOne(ctx, bson.D{{"_id", id}})
	err := res.Decode(subCat)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err
		}
	}
	return subCat, nil
}
