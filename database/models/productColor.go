package models

import (
	"github.com/abaron/lebaros-backend/lib/common"
	"github.com/jinzhu/gorm"
)

// Product color data model
type ProductColor struct {
	gorm.Model
	Product       Product     `gorm:"foreignkey:ProductID"`
	ProductID     uint        `sql:"not null"`
	ProductSize   ProductSize `gorm:"foreignkey:ProductSizeID"`
	ProductSizeID uint        `sql:"not null"`
	Sku           string      `gorm:"size:20" sql:"not null"`
	Color         string      `gorm:"size:32" sql:"not null"`
	Created       User        `gorm:"foreignkey:CreatedBy"`
	CreatedBy     uint        `sql:"not null"`
	Updated       User        `gorm:"foreignkey:UpdatedBy"`
	UpdatedBy     uint        `gorm:"default:null"`
	Deleted       User        `gorm:"foreignkey:DeletedBy"`
	DeletedBy     uint        `gorm:"default:null"`
}

// Serialize serializes product color data
func (p ProductColor) Serialize() common.JSON {
	return common.JSON{
		"id":           p.ID,
		"product_id":   p.ProductID,
		"product_size": p.ProductSize,
		"sku":          p.Sku,
		"color":        p.Color,
		"created_at":   p.CreatedAt,
		"created_by":   p.CreatedAt,
		"updated_at":   p.CreatedAt,
		"updated_by":   p.CreatedAt,
		"deleted_at":   p.CreatedAt,
		"deleted_by":   p.CreatedAt,
	}
}
