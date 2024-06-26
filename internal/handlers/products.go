package handlers

import (
	"github.com/gin-gonic/gin"
	mySession "go_shop/internal/lib/session"
	"go_shop/internal/service"
	"log"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

func (p *ProductHandler) Get(c *gin.Context) {
	session := mySession.GetSession(c)

	sizes, err := p.service.GetAllSizes()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	catsAndSubCats, err := p.service.GetAllCatsAndSubCats()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	brands, err := p.service.GetAllBrands()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	products, err := p.service.GetProducts(c)
	if err != nil {
		log.Printf("Failed to get products: %v", err)
		c.JSON(500, gin.H{"error": err.Error()})
	}
	minPrice, maxPrice := p.service.GetMinAndMaxPrice(products)
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

func (p *ProductHandler) GetOne(c *gin.Context) {
	session := mySession.GetSession(c)

	slug := c.Param("slug")
	product, err := p.service.GetBySlug(slug)
	relatedProducts, err := p.service.GetRelatedProducts(product.Category, product.Subcategory)
	if err != nil {
		log.Printf("Failed to get product: %v", err)
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.HTML(200, "product-details.html", gin.H{
		"session":     session,
		"product":     product,
		"relatedProd": relatedProducts,
		"page":        "shop",
	})
}
