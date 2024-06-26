package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	mySession "go_shop/internal/lib/session"
	"go_shop/internal/models"
	"go_shop/internal/requests"
	"go_shop/internal/service"
	"net/http"
)

const (
	sessionName = "go_shop"
)

type AuthHandler struct {
	service      service.AuthService
	sessionStore sessions.Store
}

type Session struct {
	Status string       `json:"status"`
	User   *models.User `json:"user"`
}

func NewAuthHandler(service service.AuthService, sessionStore sessions.Store) *AuthHandler {
	return &AuthHandler{
		service:      service,
		sessionStore: sessionStore,
	}
}

func (h *AuthHandler) Get(c *gin.Context) {
	c.HTML(200, "auth.html", nil)
}

func (h *AuthHandler) SignIn(c *gin.Context) {
	req := new(requests.SignIn)
	if err := c.ShouldBindBodyWithJSON(req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	user, err := h.service.GetUser(req.Email, req.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	h.createSession(c, user.ID.Hex())
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	req := new(requests.SignUp)
	if err := c.ShouldBindBodyWithJSON(req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	id, err := h.service.CreateUser(req.Name, req.Email, req.Password1, req.Password2)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	h.createSession(c, id)
}

func (h *AuthHandler) createSession(c *gin.Context, id string) {
	session, err := h.sessionStore.Get(c.Request, sessionName)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	session.Values["user_id"] = id
	if err := h.sessionStore.Save(c.Request, c.Writer, session); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, nil)
}

func (h *AuthHandler) Logout(c *gin.Context) {
	session, err := h.sessionStore.Get(c.Request, sessionName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	session.Options.MaxAge = -1
	delete(session.Values, "user_id")

	err = session.Save(c.Request, c.Writer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "logged_out"})

}

func (h *AuthHandler) Profile(c *gin.Context) {
	session := mySession.GetSession(c)
	email := c.Param("user")
	if session.User.Email != email {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "unauthorized",
		})
		return
	}
	c.HTML(200, "profile.html", gin.H{
		"session": session,
		"user":    session.User,
	})
}

func (h *AuthHandler) EditProfile(c *gin.Context) {
	session := mySession.GetSession(c)
	req := new(requests.EditProfile)
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errror": err.Error()})
		return
	}

	file, err := c.FormFile("photo")

	if err != nil {
		// TODO
		if err.Error() != "http: no such file" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}
	if err := h.service.UpdateUser(req, file, session.User.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusMovedPermanently, c.Request.URL.Path)
}
