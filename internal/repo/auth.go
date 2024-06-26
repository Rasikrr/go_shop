package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go_shop/internal/models"
	mn "go_shop/internal/storage/mongo"
	"time"
)

type AuthRepo interface {
	GetUser(email string) (*models.User, error)
	CreateUser(user *models.User) (primitive.ObjectID, error)
	GetUserById(string) (*models.User, error)
	UpdateUser(update bson.D, email string) error
}

type AuthRepoImpl struct {
	storage *mn.Storage
}

func NewAuthRepo(storage *mn.Storage) *AuthRepoImpl {
	return &AuthRepoImpl{
		storage: storage,
	}
}

func (r *AuthRepoImpl) GetUser(email string) (*models.User, error) {
	collection := r.storage.Client.Database("go_shop").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	res := collection.FindOne(ctx, bson.D{{"email", email}})
	if err := res.Err(); err != nil {
		return nil, err
	}
	user := new(models.User)
	if err := res.Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *AuthRepoImpl) CreateUser(user *models.User) (primitive.ObjectID, error) {
	collection := r.storage.Client.Database("go_shop").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		return [12]byte{}, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

func (r *AuthRepoImpl) GetUserById(idHex string) (*models.User, error) {
	collection := r.storage.Client.Database("go_shop").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	user := new(models.User)

	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		return nil, err
	}
	if err := collection.FindOne(ctx, bson.D{{"_id", id}}).Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *AuthRepoImpl) UpdateUser(data bson.D, email string) error {
	collection := r.storage.Client.Database("go_shop").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	filter := bson.D{{"email", email}}

	updateData := bson.D{
		{"$set", data},
	}

	_, err := collection.UpdateOne(ctx, filter, updateData)
	if err != nil {
		return err
	}
	return nil

}
