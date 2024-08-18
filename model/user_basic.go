package model

import (
	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
	ID    string `json:"id"`
}
