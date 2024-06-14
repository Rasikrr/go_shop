package app

import (
	"github.com/gin-gonic/gin"
	"go_shop/internal/handlers"
	"go_shop/internal/repo"
	"go_shop/internal/service"
	"go_shop/internal/storage/mongo"
	"html/template"
	"log"
	"os"
)

type App struct {
	storage *mongo.Storage
	app     *gin.Engine
}

func NewApp() *App {
	storage := mongo.ConnectDB()
	//storage.FillTestDate()
	app := gin.Default()

	return &App{
		storage: storage,
		app:     app,
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
		//"eq":  Eq,
	})
	a.app.LoadHTMLGlob("./frontend/templates/*")
	a.app.Static("/static", "./frontend/")

	productRepo := repo.NewProductRepo(a.storage)
	productsService := service.NewProductService(productRepo)
	productsHandler := handlers.NewProductHandler(productsService)

	a.app.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})

	})
	products := a.app.Group("/products")
	{
		products.GET("/", productsHandler.Get)
		products.GET("/:slug", productsHandler.GetOne)
	}

	a.app.GET("/products", productsHandler.Get)

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

func Eq(a, b float64) bool {
	return a == b
}
