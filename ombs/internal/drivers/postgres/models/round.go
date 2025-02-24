package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Round struct {
	gorm.Model
	ID          uuid.UUID `gorm:"primaryKey"`
	RoundNumber int       `gorm:"unique;not null"`
	CreatedAt   time.Time `gorm:"default:now()"`
}
