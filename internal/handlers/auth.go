package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
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
	req := new(requests.SignInRequest)
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
	req := new(requests.SignUpRequest)
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

func (h *AuthHandler) CheckSession(r *http.Request) *Session {
	session, err := h.sessionStore.Get(r, sessionName)
	id := session.Values["user_id"]
	if err != nil || id == nil {
		return &Session{
			Status: "unauthorized",
			User:   nil,
		}
	}
	user, err := h.service.GetUserById(id.(string))
	if err != nil {
		return &Session{
			Status: "unauthorized",
			User:   nil,
		}
	}
	return &Session{
		Status: "logged_in",
		User:   user,
	}
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
