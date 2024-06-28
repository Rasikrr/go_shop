package session

import (
	"github.com/gin-gonic/gin"
	"go_shop/internal/domain"
	"log"
)

func GetSession(c *gin.Context) *domain.Session {
	session := new(domain.Session)
	ctxSession, ok := c.Get("session")
	if !ok {
		log.Println("session did not get")
		session.Status = "unauthorized"
	} else {
		session = ctxSession.(*domain.Session)
	}
	return session
}
