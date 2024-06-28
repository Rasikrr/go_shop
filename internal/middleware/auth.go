package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"go_shop/internal/domain"
	"go_shop/internal/repo"
)

const (
	sessionName = "go_shop"
)

type AuthMiddleware interface {
	CheckSession() gin.HandlerFunc
}

type AuthMiddlewareImpl struct {
	sessionStore sessions.Store
	userRepo     repo.AuthRepo
}

func NewAuthMiddleware(store sessions.Store, userRepo repo.AuthRepo) *AuthMiddlewareImpl {
	return &AuthMiddlewareImpl{
		sessionStore: store,
		userRepo:     userRepo,
	}
}

func (m *AuthMiddlewareImpl) CheckSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := m.sessionStore.Get(c.Request, sessionName)
		id := session.Values["user_id"]

		ctxSession := &domain.Session{
			Status: "unauthorized",
			User:   nil,
		}

		if err == nil && id != nil {
			user, err := m.userRepo.GetUserById(id.(string))
			if err == nil {
				ctxSession.Status = "logged_in"
				ctxSession.User = user
			}
		}

		c.Set("session", ctxSession)
		c.Next()
	}
}
