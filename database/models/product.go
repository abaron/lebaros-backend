package models

import (
	"github.com/abaron/lebaros-backend/lib/common"
	"github.com/jinzhu/gorm"
)

// Product data model
type Product struct {
	gorm.Model
	User        User    `gorm:"foreignkey:UserID" sql:"not null"`
	UserID      uint    `sql:"not null"`
	Sku         string  `gorm:"size:20" sql:"not null"`
	Name        string  `gorm:"size:128" sql:"not null"`
	Description string  `sql:"type:text;"`
	Uri         string  `gorm:"size:32" sql:"not null"`
	Price       uint    `sql:"not null"`
	Revenue     float32 `sql:"not null"`
	Currency    string  `gorm:"size:3" sql:"not null"`
	Created     User    `gorm:"foreignkey:CreatedBy"`
	CreatedBy   uint    `sql:"not null"`
	Updated     User    `gorm:"foreignkey:UpdatedBy"`
	UpdatedBy   uint    `gorm:"default:null"`
	IsDeleted   bool    `gorm:"default:false" sql:"not null"`
	Deleted     User    `gorm:"foreignkey:DeletedBy"`
	DeletedBy   uint    `gorm:"default:null"`
}

// Serialize serializes product data
func (p Product) Serialize() common.JSON {
	return common.JSON{
		"id":          p.ID,
		"user":        p.User.Serialize(),
		"sku":         p.Sku,
		"name":        p.Name,
		"description": p.Description,
		"uri":         p.Uri,
		"price":       p.Price,
		"revenue":     p.Revenue,
		"currency":    p.Currency,
		"created_at":  p.CreatedAt,
		"created_by":  p.CreatedAt,
		"updated_at":  p.CreatedAt,
		"updated_by":  p.CreatedAt,
		"deleted_at":  p.CreatedAt,
		"deleted_by":  p.CreatedAt,
	}
}
