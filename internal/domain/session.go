package domain

import "go_shop/internal/models"

type Session struct {
	Status string       `json:"status"`
	User   *models.User `json:"user"`
}
