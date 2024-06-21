package app

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"go_shop/internal/handlers"
	"go_shop/internal/repo"
	"go_shop/internal/service"
	"go_shop/internal/storage/mongo"
	"html/template"
	"log"
	"math"
	"os"
)

type App struct {
	storage      *mongo.Storage
	app          *gin.Engine
	sessionStore sessions.Store
}

func NewApp() *App {
	storage := mongo.ConnectDB()
	//storage.FillTestDate()
	app := gin.Default()
	sessionStore := sessions.NewCookieStore([]byte(os.Getenv("SESSIONS_SECRET")))
	sessionStore.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600 * 8,
		HttpOnly: true,
	}

	return &App{
		storage:      storage,
		app:          app,
		sessionStore: sessionStore,
	}
}

func (a *App) Run() {
	a.app.SetFuncMap(template.FuncMap{
		"mul": Mul,
		"add": Add,
		"div": Div,
		"sub": Sub,
		"addInt": func(a, b int) int {
			return a + b
		},
		"divPag": DivPag,
		"ran":    Range,
	})
	a.app.LoadHTMLGlob("./frontend/templates/*.html")
	a.app.Static("/static", "./frontend/")

	productRepo := repo.NewProductRepo(a.storage)
	productsService := service.NewProductService(productRepo)
	productsHandler := handlers.NewProductHandler(productsService)

	products := a.app.Group("/products")
	{
		products.GET("/", productsHandler.Get)
		products.GET("/:slug", productsHandler.GetOne)
	}
	authRepo := repo.NewAuthRepo(a.storage)
	authService := service.NewAuthService(authRepo)
	authHandler := handlers.NewAuthHandler(authService, a.sessionStore)

	a.app.GET("/", func(c *gin.Context) {
		session := authHandler.CheckSession(c.Request)
		c.HTML(200, "index.html", gin.H{"session": session})
	})

	auth := a.app.Group("/auth")
	{
		auth.GET("/", authHandler.Get)
		auth.POST("/signin", authHandler.SignIn)
		auth.POST("/signup", authHandler.SignUp)
		auth.GET("/logout", authHandler.Logout)
		auth.GET("/profile/:user")
		auth.POST()
	}

	if err := a.app.Run(srvConfig()); err != nil {
		log.Fatalf("Failed to run server: %v", err)
		return
	}
}

func (a *App) Shutdown() {
	a.storage.Close()
	// TODO Graceful shutdown
}

func srvConfig() string {
	host := os.Getenv("SRV_HOST")
	port := os.Getenv("SRV_PORT")
	if host == "" || port == "" {
		log.Fatalf("HOST or PORT is not set")
	}
	return host + ":" + port
}

func Add(a, b float64) float64 {
	return a + b
}

func Mul(a, b float64) float64 {
	return a * b
}

func Div(a float64, b int) float64 {
	if b != 0 {
		return a / float64(b)
	}
	return 0
}

func Sub(a, b float64) float64 {
	return a - b
}

func DivPag(a, b int) int {
	return int(math.Ceil(float64(a) / float64(b)))
}

func Range(start, end int) []int {
	if end < start {
		return nil
	}
	nums := make([]int, end-start+1)
	for i := range nums {
		nums[i] = start + i
	}
	return nums
}
