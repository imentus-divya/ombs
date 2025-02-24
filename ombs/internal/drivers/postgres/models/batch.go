package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Batch struct {
	gorm.Model
	ID              uuid.UUID `gorm:"primaryKey"`
	RoundID         uuid.UUID `gorm:"not null"`
	BatchNumber     int       `gorm:"not null"`
	MessageHash     string    `gorm:"not null"`
	MessageJSON     string    `gorm:"type:json;not null"`
	TotalSignatures int       `gorm:"default:0"`
	SentSignatures  int       `gorm:"default:0"`
	CreatedAt       time.Time `gorm:"default:now()"`
	Round           Round     `gorm:"foreignKey:RoundID"`
}
