package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Signature struct {
	gorm.Model
	ID                  uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	BatchID             uuid.UUID `gorm:"not null"`
	NodeID              uuid.UUID `gorm:"not null"`
	OracleNodePublicKey string    `gorm:"not null"`
	Signature           string    `gorm:"not null"`
	CreatedAt           time.Time `gorm:"default:now()"`
	Batch               Batch     `gorm:"foreignKey:BatchID"`
}
