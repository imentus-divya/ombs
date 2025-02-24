package models

import "gorm.io/gorm"

// OracleNodeStatus represents the status of an oracle node
type OracleNodeStatus struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey"`
	StatusName string `gorm:"unique;not null"`
}
