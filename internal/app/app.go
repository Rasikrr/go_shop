package app

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"go_shop/internal/handlers"
	mySession "go_shop/internal/lib/session"
	"go_shop/internal/middleware"
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
		MaxAge:   0,
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
	a.app.Static("/media", "./media/")

	cartRepo := repo.NewCartRepo(a.storage)
	cartService := service.NewCartService(cartRepo)

	productRepo := repo.NewProductRepo(a.storage)
	productsService := service.NewProductService(productRepo)
	productsHandler := handlers.NewProductHandler(productsService, cartService)

	authRepo := repo.NewAuthRepo(a.storage)
	authService := service.NewAuthService(authRepo)
	authHandler := handlers.NewAuthHandler(authService, a.sessionStore)

	// Middlewares
	authMiddleware := middleware.NewAuthMiddleware(a.sessionStore, authRepo)
	a.app.Use(authMiddleware.CheckSession())
	a.app.Use(middleware.CartMiddleware(cartService))
	a.app.Use(middleware.NotFound())

	products := a.app.Group("/products")
	{
		products.GET("/", productsHandler.Get)
		products.GET("/:slug", productsHandler.GetOne)
	}

	a.app.GET("/", func(c *gin.Context) {
		session := mySession.GetSession(c)
		topSales, err := productRepo.GetTopSales(6)
		if err != nil {
			log.Printf("failed to get top sales products | %v", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.HTML(200, "index.html", gin.H{
			"session":  session,
			"topSales": topSales,
			"page":     "home",
		})
	})

	a.app.GET("/cart", func(c *gin.Context) {
		c.HTML(200, "shopping-cart.html", gin.H{})
	})

	a.app.GET("/about", func(c *gin.Context) {
		session := mySession.GetSession(c)

		c.HTML(200, "about.html", gin.H{
			"session": session,
			"page":    "about",
		})
	})
	a.app.GET("/contacts", func(c *gin.Context) {
		session := mySession.GetSession(c)

		c.HTML(200, "contacts.html", gin.H{
			"session": session,
			"page":    "contacts",
		})
	})

	auth := a.app.Group("/auth")
	{
		auth.GET("/", authHandler.Get)
		auth.POST("/signin", authHandler.SignIn)
		auth.POST("/signup", authHandler.SignUp)
		auth.GET("/logout", authHandler.Logout)
		auth.GET("/profile/:user", authHandler.Profile)
		auth.POST("/profile/:user", authHandler.EditProfile)
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
