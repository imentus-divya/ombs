package models

import "gorm.io/gorm"

// OracleNode represents an oracle node in the system
type OracleNode struct {
	gorm.Model
	NodeAddress         string `gorm:"primaryKey"`
	OracleNodePublicKey string `gorm:"unique,not null"`
	RewardWalletAddress string
	ProxyWalletAddress  string
	StatusID            uint             `gorm:"not null"`
	RoleID              uint             `gorm:"not null"`
	Status              OracleNodeStatus `gorm:"foreignKey:StatusID"`
	Role                Role             `gorm:"foreignKey:RoleID"`
}
