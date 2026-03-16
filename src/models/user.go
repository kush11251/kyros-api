package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `json:"id" gorm:"primarykey"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
