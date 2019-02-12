package product

import (
	"strconv"
	"strings"

	"github.com/abaron/lebaros-backend/database/models"
	"github.com/abaron/lebaros-backend/lib/common"
	"github.com/abaron/lebaros-backend/lib/helpers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Product type alias
type Product = models.Product
type ProductSize = models.ProductSize
type ProductColor = models.ProductColor
type ProductStock = models.ProductStock
type ProductStockLog = models.ProductStockLog

// User type alias
type User = models.User

// JSON type alias
type JSON = common.JSON

// Import existing data
func importExists(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	data := getCsvProduct("./database/produkTokoIjah.csv")
	var imported []interface{}

	for _, row := range data {
		stock, _ := strconv.ParseUint(row.Stock, 10, 32)
		min := float32(100000)
		max := float32(500000)
		minf := float32(10000)
		maxf := float32(50000)
		var productID uint

		var id string
		r := db.Table("products").Where("name = ? AND is_deleted = ?", row.Name, 0).Select("id").Row()
		r.Scan(&id)

		if id == "" {

			// Bulk insert without transaction
			// Insert new product
			product := Product{
				UserID:    0,
				Name:      row.Name,
				CreatedBy: 0,
				IsDeleted: false,
				Uri:       strings.ToLower(strings.Replace(row.Name, " ", "-", -1)),
			}
			db.Create(&product)
			productID = product.ID
		} else {
			id, _ := strconv.ParseUint(id, 10, 32)
			productID = uint(id)
		}

		// Insert product size
		productSize := ProductSize{
			ProductID: uint(productID),
			Sku:       row.Sku,
			Size:      row.Size,
			CreatedBy: 0,
		}
		db.Create(&productSize)

		// Insert product color
		productColor := ProductColor{
			ProductID:     uint(productID),
			ProductSizeID: uint(productSize.ID),
			Sku:           row.Sku,
			Color:         row.Color,
			CreatedBy:     0,
		}
		db.Create(&productColor)

		// Insert product stock
		productStock := ProductStock{
			ProductID:      uint(productID),
			ProductColorID: uint(productColor.ID),
			Sku:            row.Sku,
			Price:          helpers.IRandFloat32(min, max, 1)[0],
			Revenue:        helpers.IRandFloat32(minf, maxf, 1)[0],
			Currency:       "IDR",
			Stock:          uint16(stock),
			CreatedBy:      0,
		}
		db.Create(&productStock)

		if productStock.ID != 0 {
			// Insert product stock log
			productStockLog := ProductStockLog{
				ProductStockID: uint(productStock.ID),
				Stock:          uint16(stock),
				Flag:           uint8(0),
				CreatedBy:      0,
			}
			db.Create(&productStockLog)
			imported = append(imported, row)
		}
	}

	response := common.Response{
		Code:     200,
		Status:   "OK",
		Messages: "",
		Data:     imported,
	}
	c.JSON(200, response)
}

// Handler for display stock list
func stocks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	type result struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Size  string `json:"size"`
		Color string `json:"color"`
		Sku   string `json:"sku"`
		Stock uint   `json:"stock"`
	}

	var results []result
	db.Table("products").Select(`
		products.id,
		products.name,
		product_sizes.size,
		product_colors.color,
		product_stocks.sku,
		product_stocks.stock
	`).Joins(`
		left join product_stocks on product_stocks.product_id = products.id
	`).Joins(`
		left join product_colors on product_colors.sku = product_stocks.sku
	`).Joins(`
		left join product_sizes on product_sizes.sku = product_stocks.sku
	`).Scan(&results)

	response := common.Response{
		Code:     200,
		Status:   "OK",
		Messages: "",
		Data:     results,
	}
	c.JSON(200, response)
}
