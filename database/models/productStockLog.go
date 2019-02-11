package models

import (
	"github.com/abaron/lebaros-backend/lib/common"
	"github.com/jinzhu/gorm"
)

// Product stock log data model
type ProductStockLog struct {
	gorm.Model
	ProductStock   Product `gorm:"foreignkey:ProductStockID"`
	ProductStockID uint    `sql:"not null"`
	Stock          uint16  `sql:"not null"`
	Created        User    `gorm:"foreignkey:CreatedBy"`
	CreatedBy      uint    `sql:"not null"`
	Updated        User    `gorm:"foreignkey:UpdatedBy"`
	UpdatedBy      uint    `gorm:"default:null"`
	Deleted        User    `gorm:"foreignkey:DeletedBy"`
	DeletedBy      uint    `gorm:"default:null"`
}

// Serialize serializes product stock log data
func (p ProductStockLog) Serialize() common.JSON {
	return common.JSON{
		"id":               p.ID,
		"product_stock_id": p.ProductStockID,
		"stock":            p.Stock,
		"created_at":       p.CreatedAt,
		"created_by":       p.CreatedAt,
		"updated_at":       p.CreatedAt,
		"updated_by":       p.CreatedAt,
		"deleted_at":       p.CreatedAt,
		"deleted_by":       p.CreatedAt,
	}
}
