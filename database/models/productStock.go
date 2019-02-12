package models

import (
	"github.com/abaron/lebaros-backend/lib/common"
	"github.com/jinzhu/gorm"
)

// Product stock data model
type ProductStock struct {
	gorm.Model
	Product        Product      `gorm:"foreignkey:ProductID"`
	ProductID      uint         `sql:"not null"`
	ProductColor   ProductColor `gorm:"foreignkey:ProductColorID"`
	ProductColorID uint         `sql:"not null"`
	Sku            string       `gorm:"size:20" sql:"not null"`
	Stock          uint16       `gorm:"size:32"`
	Price          float32      `sql:"not null"`
	Revenue        float32      `sql:"not null"`
	Currency       string       `gorm:"size:3" sql:"not null"`
	Created        User         `gorm:"foreignkey:CreatedBy"`
	CreatedBy      uint         `sql:"not null"`
	Updated        User         `gorm:"foreignkey:UpdatedBy"`
	UpdatedBy      uint         `gorm:"default:null"`
	Deleted        User         `gorm:"foreignkey:DeletedBy"`
	DeletedBy      uint         `gorm:"default:null"`
}

// Serialize serializes product stock data
func (p ProductStock) Serialize() common.JSON {
	return common.JSON{
		"id":            p.ID,
		"product_id":    p.ProductID,
		"product_color": p.ProductColor,
		"sku":           p.Sku,
		"stock":         p.Stock,
		"price":         p.Price,
		"revenue":       p.Revenue,
		"currency":      p.Currency,
		"created_at":    p.CreatedAt,
		"created_by":    p.CreatedAt,
		"updated_at":    p.CreatedAt,
		"updated_by":    p.CreatedAt,
		"deleted_at":    p.CreatedAt,
		"deleted_by":    p.CreatedAt,
	}
}
