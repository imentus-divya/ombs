package service

import (
	"ombs/internal/domain"
)

// AuthService handles authentication logic
type AuthService struct{}

// NewAuthService initializes AuthService
func NewAuthService() *AuthService {
	return &AuthService{}
}

// AuthenticateNode processes NodeInfo and returns a message
func (s *AuthService) AuthenticateNode(node domain.NodeInfo) string {
	// In real implementation, verify `node.messageHash` and `node.signature`
	return "OK"
}
