package models

import (
	"github.com/abaron/lebaros-backend/lib/common"
	"github.com/jinzhu/gorm"
)

// Sales data model
type Sales struct {
	gorm.Model
	Product   Product `gorm:"foreignkey:ProductID"`
	ProductID uint    `sql:"not null"`
	Price     float32 `sql:"not null"`
	Revenue   float32 `sql:"not null"`
	Quantity  uint16  `gorm:"size:4"`
	Currency  string  `gorm:"size:3"`
	Created   User    `gorm:"foreignkey:CreatedBy"`
	CreatedBy uint    `sql:"not null"`
	Updated   User    `gorm:"foreignkey:UpdatedBy"`
	UpdatedBy uint    `gorm:"default:null"`
	Deleted   User    `gorm:"foreignkey:DeletedBy"`
	DeletedBy uint    `gorm:"default:null"`
}

// Serialize serializes sales data
func (p Sales) Serialize() common.JSON {
	return common.JSON{
		"id":         p.ID,
		"price":      p.Price,
		"revenue":    p.Revenue,
		"quantity":   p.Quantity,
		"currency":   p.Currency,
		"created_at": p.CreatedAt,
		"created_by": p.CreatedAt,
		"updated_at": p.CreatedAt,
		"updated_by": p.CreatedAt,
		"deleted_at": p.CreatedAt,
		"deleted_by": p.CreatedAt,
	}
}
