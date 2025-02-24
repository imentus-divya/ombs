package models

import (
	"time"

	"gorm.io/gorm"
)

type JWTIssuance struct {
	gorm.Model
	NodeAddress     string     `gorm:"not null"`
	JWTAccessToken  string     `gorm:"not null"`
	JWTRefreshToken string     `gorm:"not null"`
	IssuedAt        time.Time  `gorm:"default:now()"`
	ExpiresAt       time.Time  `gorm:"not null"`
	Blacklisted     bool       `gorm:"default:false"`
	OracleNode      OracleNode `gorm:"foreignKey:NodeAddress"`
}
