package repo

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongo2 "go.mongodb.org/mongo-driver/mongo"
	"go_shop/internal/models"
	"go_shop/internal/storage"
	"go_shop/internal/storage/mongo"
	"strings"
)

type CartRepo interface {
	AddToCart(cartItem *models.CartItem) error
	GetProductById(string) (*models.Product, error)
	GetUserCartItems(userId primitive.ObjectID) ([]*models.CartItem, error)
}

type CartRepoImpl struct {
	storage *mongo.Storage
}

func NewCartRepo(storage *mongo.Storage) CartRepo {
	return &CartRepoImpl{
		storage: storage,
	}
}

func (r *CartRepoImpl) GetUserCartItems(userId primitive.ObjectID) ([]*models.CartItem, error) {
	collection := r.storage.Client.Database(storage.DB_NAME).Collection(storage.CARTS)
	ctx, cancel := context.WithTimeout(context.Background(), storage.TIMEOUT)

	defer cancel()

	filter := bson.D{{"user_id", userId}}
	cur, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	var cartItems []*models.CartItem

	if err := cur.All(ctx, &cartItems); err != nil {
		return nil, err
	}
	return cartItems, nil
}

//func (r *CartRepoImpl) GetCartOverall(items []*models.CartItem) (float64, error) {
//	var overall float64
//	collection := r.storage.Client.Database(storage.DB_NAME).Collection(storage.PRODUCTS)
//	ctx, cancel := context.WithTimeout(context.Background(), storage.TIMEOUT)
//
//	defer cancel()
//
//	for _, item := range items {
//		filter := bson.D{{"_id", item.ProductID}}
//		res := collection.FindOne(ctx, filter)
//		if res.Err() != nil {
//			return -1.0, res.Err()
//		}
//		product := new(models.Product)
//		if err := res.Decode(product); err != nil {
//			return -1.0, err
//		}
//		overall += product.Price * float64(item.Amount)
//	}
//	return overall, nil
//}

func (r *CartRepoImpl) AddToCart(cartItem *models.CartItem) error {
	fmt.Printf("%+v\n", cartItem)
	cartsCollection := r.storage.Client.Database(storage.DB_NAME).Collection(storage.CARTS)
	prodsCollection := r.storage.Client.Database(storage.DB_NAME).Collection(storage.PRODUCTS)
	ctx, cancel := context.WithTimeout(context.Background(), storage.TIMEOUT)

	defer cancel()

	product := new(models.Product)

	res := prodsCollection.FindOne(ctx, bson.D{{"_id", cartItem.ProductID}})
	if res.Err() != nil {
		if errors.Is(res.Err(), mongo2.ErrNoDocuments) {
			return ErrNoDocuments
		}
		return res.Err()
	}
	if err := res.Decode(product); err != nil {
		return err
	}
	found := false
	fmt.Println(product)
	for _, size := range product.Sizes {
		if strings.ToUpper(size.Size) == cartItem.Size {
			if size.Amount-cartItem.Amount < 0 {
				return ErrOutOfStock
			}
			found = true
			break
		}
	}
	if !found {
		return ErrNoSize
	}
	filter := bson.D{{"user_id", cartItem.UserID}, {"product_id", cartItem.ProductID}, {"size", cartItem.Size}}
	update := bson.D{{"$inc", bson.D{{"amount", cartItem.Amount}}}}
	updateRes, err := cartsCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if updateRes.ModifiedCount == 0 {
		_, err := cartsCollection.InsertOne(ctx, cartItem)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (r *CartRepoImpl) GetProductById(id string) (*models.Product, error) {
	collection := r.storage.Client.Database(storage.DB_NAME).Collection(storage.PRODUCTS)
	ctx, cancel := context.WithTimeout(context.Background(), storage.TIMEOUT)

	defer cancel()

	product := new(models.Product)

	if err := collection.FindOne(ctx, bson.D{{"_id", id}}).Decode(product); err != nil {
		return nil, err
	}
	return product, nil
}
