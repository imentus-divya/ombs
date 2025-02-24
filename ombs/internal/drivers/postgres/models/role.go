package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	RoleName string `gorm:"size:255;not null;unique"`
}
