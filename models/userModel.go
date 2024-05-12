package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Email    string `gorm:"unique"`
	Password string `json:"-" gorm:"min:6"`
	Photos   []*Photo `gorm:"foreignKey:UserID"`
}
