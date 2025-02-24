package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LeaderTransaction struct {
	gorm.Model
	LeaderID       string    `gorm:"not null"`
	RoundID        uuid.UUID `gorm:"not null"`
	BatchID        uuid.UUID `gorm:"not null"`
	SignaturesSent int       `gorm:"not null"`
	TxHash         string    `gorm:"not null"`
	Timestamp      time.Time `gorm:"default:now()"`
	Round          Round     `gorm:"foreignKey:RoundID"`
	Batch          Batch     `gorm:"foreignKey:BatchID"`
	Leader         Leader    `gorm:"foreignKey:LeaderID"`
}
