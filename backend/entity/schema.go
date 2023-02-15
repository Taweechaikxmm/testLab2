package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Address string `valid:"required~test"`
}
