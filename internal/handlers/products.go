package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	mySession "go_shop/internal/lib/session"
	"go_shop/internal/models"
	"go_shop/internal/service"
	"log"
	"net/http"
)

type ProductHandler struct {
	productService service.ProductService
	cartService    service.CartService
}

func NewProductHandler(prodService service.ProductService, cartService service.CartService) *ProductHandler {
	return &ProductHandler{
		productService: prodService,
		cartService:    cartService,
	}
}

func (h *ProductHandler) Get(c *gin.Context) {
	session := mySession.GetSession(c)

	sizes, err := h.productService.GetAllSizes()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	catsAndSubCats, err := h.productService.GetAllCatsAndSubCats()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	brands, err := h.productService.GetAllBrands()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	products, err := h.productService.GetProducts(c)
	if err != nil {
		log.Printf("Failed to get products: %v", err)
		c.JSON(500, gin.H{"error": err.Error()})
	}
	minPrice, maxPrice := h.productService.GetMinAndMaxPrice(products)
	c.HTML(200, "products.html", gin.H{
		"session":        session,
		"sizes":          sizes,
		"catsAndSubcats": catsAndSubCats,
		"brands":         brands,
		"products":       products,
		"minPrice":       minPrice,
		"maxPrice":       maxPrice,
		"page":           "shop",
	})
}

func (h *ProductHandler) GetOne(c *gin.Context) {
	session := mySession.GetSession(c)

	var overall float64
	cartOverall, ok := c.Get("cartOverall")
	if ok {
		overall = cartOverall.(float64)
	}

	slug := c.Param("slug")
	product, err := h.productService.GetBySlug(slug)
	relatedProducts, err := h.productService.GetRelatedProducts(product.Category, product.Subcategory)
	if err != nil {
		log.Printf("Failed to get product: %v", err)
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.HTML(200, "product-details.html", gin.H{
		"session":     session,
		"product":     product,
		"cartOverall": overall,
		"relatedProd": relatedProducts,
		"page":        "shop",
	})
}

func (h *ProductHandler) AddToCart(c *gin.Context) {
	session := mySession.GetSession(c)
	fmt.Println(session)
	if session.Status == "unauthorized" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unauthorized"})
		return
	}
	cartItem := new(models.CartItem)
	if err := c.ShouldBindBodyWithJSON(cartItem); err != nil {
		fmt.Println("HERE")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(cartItem)
	cartItem.UserID = session.User.ID
	err := h.cartService.AddToCart(cartItem)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"error": "ok"})
}
