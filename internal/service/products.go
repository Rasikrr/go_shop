package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go_shop/internal/models"
	"go_shop/internal/repo"
	"log"
	"strconv"
	"strings"
)

type ProductService interface {
	GetProducts(ctx *gin.Context) ([]*models.Product, error)
	GetAllSizes() ([]*models.Size, error)
	GetAllCategories() ([]*models.Category, error)
	GetAllBrands() ([]*models.Brand, error)
	GetBySlug(string) (*models.Product, error)
	GetAllCatsAndSubCats() ([]*CatsAndSubCats, error)
	GetRelatedProducts(category, subcategory string) ([]*models.Product, error)
}

type CatsAndSubCats struct {
	Category      *models.Category
	SubCategories []*models.Subcategory
}

type ProductServiceImpl struct {
	repo repo.ProductRepo
}

func NewProductService(repo repo.ProductRepo) *ProductServiceImpl {
	return &ProductServiceImpl{
		repo: repo,
	}
}

func (p *ProductServiceImpl) GetRelatedProducts(category, subcategory string) ([]*models.Product, error) {
	return p.repo.GetRelatedProducts(category, subcategory)
}

func (p *ProductServiceImpl) GetProducts(ctx *gin.Context) ([]*models.Product, error) {
	brandsFilt := ctx.QueryArray("brands")
	catsFilt := ctx.QueryArray("categories")
	subcatFilt := ctx.QueryArray("subcategories")
	sexFilt := ctx.Query("sex")
	sizesFilt := ctx.QueryArray("sizes")
	priceFilt := ctx.Query("price")
	nameFilt := ctx.Query("prod-name")

	low, high := p.parsePrice(priceFilt)

	query := bson.D{}
	if nameFilt != "" {
		query = append(query, bson.E{"name", primitive.Regex{Pattern: fmt.Sprintf(".*%s.*", nameFilt), Options: "i"}})
	}
	if low != -1 && high != -1 {
		query = append(query, bson.E{"price", bson.D{{"$gte", low}, {"$lte", high}}})
	}
	if len(brandsFilt) > 0 {
		query = append(query, bson.E{"brand", bson.D{{"$in", brandsFilt}}})
	}
	if len(catsFilt) > 0 {
		query = append(query, bson.E{"category", bson.D{{"$in", catsFilt}}})
	}
	if len(subcatFilt) > 0 {
		query = append(query, bson.E{"subcategory", bson.D{{"$in", subcatFilt}}})
	}
	if sexFilt != "" {
		query = append(query, bson.E{"sex", sexFilt})
	}
	if len(sizesFilt) > 0 {
		query = append(query, bson.E{Key: "sizes", Value: bson.D{{"$elemMatch", bson.D{
			{"size", bson.D{{"$in", sizesFilt}}},
			{"amount", bson.D{{"$gt", 0}}},
		}}}})
	}
	products, err := p.repo.GetProducts(query)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductServiceImpl) GetAllSizes() ([]*models.Size, error) {
	sizes, err := p.repo.GetAllSizes()
	if err != nil {
		return nil, err
	}
	return sizes, nil
}

func (p *ProductServiceImpl) GetAllCategories() ([]*models.Category, error) {
	categories, err := p.repo.GetAllCategories()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (p *ProductServiceImpl) GetAllBrands() ([]*models.Brand, error) {
	brands, err := p.repo.GetAllBrands()
	if err != nil {
		return nil, err
	}
	return brands, nil
}

func (p *ProductServiceImpl) GetBySlug(slug string) (*models.Product, error) {
	product, err := p.repo.GetProductBySlug(slug)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductServiceImpl) parsePrice(price string) (float64, float64) {
	ran := strings.Split(price, " - ")
	if len(ran) != 2 {
		return -1, -1
	}
	low, err := strconv.ParseFloat(ran[0][1:], 64)
	if err != nil {
		log.Println(err)
		return -1, -1
	}
	high, err := strconv.ParseFloat(ran[1][1:], 64)
	if err != nil {
		log.Println(err)
		return -1, -1
	}
	return low, high
}

func (p *ProductServiceImpl) GetAllCatsAndSubCats() ([]*CatsAndSubCats, error) {
	cats, err := p.repo.GetAllCategories()
	if err != nil {
		return nil, err
	}
	catsAndSubCats := make([]*CatsAndSubCats, 0, len(cats))
	for _, cat := range cats {
		subCats := make([]*models.Subcategory, 0, len(cat.SubcatsID))
		for _, idHex := range cat.SubcatsID {
			id, err := primitive.ObjectIDFromHex(idHex)
			if err != nil {
				log.Println("failed to get ID from hex")
				return nil, err
			}
			subcat, err := p.repo.GetSubCatById(id)
			if err != nil {
				return nil, err
			}
			fmt.Println(id, subcat.Name)
			subCats = append(subCats, subcat)
		}
		catAndSubcats := &CatsAndSubCats{
			Category:      cat,
			SubCategories: subCats,
		}
		catsAndSubCats = append(catsAndSubCats, catAndSubcats)
	}
	return catsAndSubCats, nil
}
