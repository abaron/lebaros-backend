package models

import (
	"github.com/abaron/lebaros-backend/lib/common"
	"github.com/jinzhu/gorm"
)

// User data model
type User struct {
	gorm.Model
	Username     string
	Fullname     string
	PasswordHash string
}

// Serialize serializes user data
func (u *User) Serialize() common.JSON {
	return common.JSON{
		"id":       u.ID,
		"username": u.Username,
		"fullname": u.Fullname,
	}
}

func (u *User) Read(m common.JSON) {
	u.ID = uint(m["id"].(float64))
	u.Username = m["username"].(string)
	u.Fullname = m["fullname"].(string)
}

// for migrate
type tabler interface {
	TableName() string
}
