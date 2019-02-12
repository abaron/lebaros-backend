package models

import (
	"github.com/abaron/lebaros-backend/lib/common"
	"github.com/jinzhu/gorm"
)

// Product size data model
type ProductSize struct {
	gorm.Model
	Product   Product `gorm:"foreignkey:ProductID"`
	ProductID uint    `sql:"not null"`
	Sku       string  `gorm:"size:20" sql:"not null"`
	Size      string  `gorm:"size:16"`
	Created   User    `gorm:"foreignkey:CreatedBy"`
	CreatedBy uint    `sql:"not null"`
	Updated   User    `gorm:"foreignkey:UpdatedBy"`
	UpdatedBy uint    `gorm:"default:null"`
	Deleted   User    `gorm:"foreignkey:DeletedBy"`
	DeletedBy uint    `gorm:"default:null"`
}

// Serialize serializes product size data
func (p ProductSize) Serialize() common.JSON {
	return common.JSON{
		"id":         p.ID,
		"product_id": p.ProductID,
		"sku":        p.Sku,
		"size":       p.Size,
		"created_at": p.CreatedAt,
		"created_by": p.CreatedAt,
		"updated_at": p.CreatedAt,
		"updated_by": p.CreatedAt,
		"deleted_at": p.CreatedAt,
		"deleted_by": p.CreatedAt,
	}
}
