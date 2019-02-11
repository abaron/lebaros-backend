package models

import (
	"github.com/abaron/lebaros-backend/lib/common"
	"github.com/jinzhu/gorm"
)

// Post data model
type Post struct {
	gorm.Model
	Text   string `sql:"type:text;" sql:"not null"`
	User   User   `gorm:"foreignkey:UserID"`
	UserID uint   `sql:"not null"`
}

// Serialize serializes post data
func (p Post) Serialize() common.JSON {
	return common.JSON{
		"id":         p.ID,
		"text":       p.Text,
		"user":       p.User.Serialize(),
		"created_at": p.CreatedAt,
	}
}
