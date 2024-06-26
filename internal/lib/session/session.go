package session

import (
	"github.com/gin-gonic/gin"
	"go_shop/internal/middleware"
	"log"
)

func GetSession(c *gin.Context) *middleware.Session {
	session := new(middleware.Session)
	ctxSession, ok := c.Get("session")
	if !ok {
		log.Println("session did not get")
		session.Status = "unauthorized"
	} else {
		session = ctxSession.(*middleware.Session)
	}
	return session
}
