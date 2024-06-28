package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go_shop/internal/models"
	"go_shop/internal/repo"
	"log"
	"strings"
)

type CartService interface {
	AddToCart(cartItem *models.CartItem) error
	GetUserCartItems(userId primitive.ObjectID) ([]*models.CartItem, error)
	GetCartOverall([]*models.CartItem) (float64, error)
}

type CartServiceImpl struct {
	repo repo.CartRepo
}

func NewCartService(repo repo.CartRepo) CartService {
	return &CartServiceImpl{
		repo: repo,
	}
}

func (s *CartServiceImpl) AddToCart(item *models.CartItem) error {
	item.Size = strings.ToUpper(item.Size)
	err := s.repo.AddToCart(item)
	if err != nil {
		log.Printf("failed to add to cart | %v", err)
		return err
	}
	return nil
}

func (s *CartServiceImpl) GetUserCartItems(userId primitive.ObjectID) ([]*models.CartItem, error) {
	items, err := s.repo.GetUserCartItems(userId)
	if err != nil {
		log.Printf("failed to get cart items | user: %v, err: %v", userId, err)
		return nil, err
	}
	return items, nil
}

func (s *CartServiceImpl) GetCartOverall(items []*models.CartItem) (float64, error) {
	var overall float64

	for _, item := range items {
		product, err := s.repo.GetProductById(item.ProductID)
		if err != nil {
			log.Printf("failed to get product | %v", err)
			return -1.0, err
		}
		overall += product.Price * float64(item.Amount)
	}
	return overall, nil
}
