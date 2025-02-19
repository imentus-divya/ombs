package models

// status.go
// OracleNodeStatus represents the status of an oracle node
type OracleNodeStatus struct {
	ID         uint   `gorm:"primaryKey"`
	StatusName string `gorm:"unique;not null"`
}

// role.go
// Role represents different roles within the system
type Role struct {
	ID       uint   `gorm:"primaryKey"`
	RoleName string `gorm:"unique;not null"`
}

// oracle_node.go
// OracleNode represents an oracle node in the system
type OracleNode struct {
	NodeAddress         string `gorm:"primaryKey"`
	OracleNodePublicKey string `gorm:"not null"`
	RewardWalletAddress string `gorm:"not null"`
	ProxyWalletAddress  string
	StatusID            uint
	RoleID              uint
	Status              OracleNodeStatus `gorm:"foreignKey:StatusID"`
	Role                Role             `gorm:"foreignKey:RoleID"`
}
