package handlers

import (
	"github.com/gin-gonic/gin"
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

func (p *ProductHandler) Get(ctx *gin.Context) {
	sizes, err := p.service.GetAllSizes()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	cats, err := p.service.GetAllCategories()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	brands, err := p.service.GetAllBrands()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	products, err := p.service.GetProducts(ctx)
	if err != nil {
		log.Printf("Failed to get products: %v", err)
		ctx.JSON(500, gin.H{"error": err.Error()})
	}

	ctx.HTML(200, "products.html", gin.H{
		"sizes":    sizes,
		"cats":     cats,
		"brands":   brands,
		"products": products,
	})
}

func (p *ProductHandler) GetOne(c *gin.Context) {
	slug := c.Param("slug")
	product, err := p.service.GetBySlug(slug)
	if err != nil {
		log.Printf("Failed to get product: %v", err)
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.HTML(200, "product-details.html", gin.H{
		"product": product,
	})
}
