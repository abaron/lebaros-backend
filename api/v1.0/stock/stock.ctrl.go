package stock

import (
	"time"

	"github.com/abaron/lebaros-backend/database/models"
	"github.com/abaron/lebaros-backend/lib/common"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Product stock type alias
type Product = models.Product
type ProductSize = models.ProductSize
type ProductColor = models.ProductColor
type ProductStock = models.ProductStock
type ProductStockLog = models.ProductStockLog

// User type alias
type User = models.User

// JSON type alias
type JSON = common.JSON

// Handler for display stock list
func in(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	from := c.PostForm("from")
	to := c.PostForm("to")

	type result struct {
		ID        uint      `json:"id"`
		Name      string    `json:"name"`
		Size      string    `json:"size"`
		Color     string    `json:"color"`
		Sku       string    `json:"sku"`
		Stock     uint      `json:"stock_in"`
		CreatedAt time.Time `json:"created_at"`
	}

	var results []result
	db.Table("product_stock_logs").Select(`
		products.id,
		products.name,
		product_sizes.size,
		product_colors.color,
		product_stocks.sku,
		product_stock_logs.stock,
		product_stock_logs.created_at
	`).Joins(`
		left join product_stocks on product_stocks.id = product_stock_logs.product_stock_id
	`).Joins(`
		left join product_colors on product_colors.sku = product_stocks.sku
	`).Joins(`
		left join product_sizes on product_sizes.sku = product_stocks.sku
	`).Joins(`
		left join products on products.id = product_stocks.id
	`).Where(`
		product_stock_logs.flag = ? AND
		(
			product_stock_logs.created_at
			BETWEEN ? AND ?
		)
	`, 0, from, to).Scan(&results)

	response := common.Response{
		Code:     200,
		Status:   "OK",
		Messages: "",
		Data:     results,
	}
	c.JSON(200, response)
}

func out(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	from := c.PostForm("from")
	to := c.PostForm("to")

	type result struct {
		ID        uint      `json:"id"`
		Name      string    `json:"name"`
		Size      string    `json:"size"`
		Color     string    `json:"color"`
		Sku       string    `json:"sku"`
		Stock     uint      `json:"stock_out"`
		CreatedAt time.Time `json:"created_at"`
	}

	var results []result
	db.Table("product_stock_logs").Select(`
		products.id,
		products.name,
		product_sizes.size,
		product_colors.color,
		product_stocks.sku,
		product_stock_logs.stock,
		product_stock_logs.created_at
	`).Joins(`
		left join product_stocks on product_stocks.id = product_stock_logs.product_stock_id
	`).Joins(`
		left join product_colors on product_colors.sku = product_stocks.sku
	`).Joins(`
		left join product_sizes on product_sizes.sku = product_stocks.sku
	`).Joins(`
		left join products on products.id = product_stocks.id
	`).Where(`
		product_stock_logs.flag = ? AND
		(
			product_stock_logs.created_at
			BETWEEN ? AND ?
		)
	`, 1, from, to).Scan(&results)

	response := common.Response{
		Code:     200,
		Status:   "OK",
		Messages: "",
		Data:     results,
	}
	c.JSON(200, response)
}
