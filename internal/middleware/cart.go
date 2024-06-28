package middleware

import (
	"github.com/gin-gonic/gin"
	mySession "go_shop/internal/lib/session"
	"go_shop/internal/service"
)

func CartMiddleware(srv service.CartService) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := mySession.GetSession(c)
		if session.Status == "logged_in" {
			cartItems, err := srv.GetUserCartItems(session.User.ID)
			if err == nil {
				overall, err := srv.GetCartOverall(cartItems)
				if err == nil {
					c.Set("cartItems", cartItems)
					c.Set("cartOverall", overall)
				}
			}
		}
		c.Next()
	}
}
