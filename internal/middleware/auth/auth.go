package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"net/http"
)

func Middleware(store sessions.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := store.Get(c.Request, "session-name")
		if err != nil || session.Values["user_id"] == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
