package models

import (
	"time"

	"gorm.io/gorm"
)

type Leader struct {
	gorm.Model
	ID          string     `gorm:"primaryKey"`
	NodeAddress string     `gorm:"unique;not null"`
	AssignedAt  time.Time  `gorm:"default:now()"`
	OracleNode  OracleNode `gorm:"foreignKey:NodeAddress"`
}
