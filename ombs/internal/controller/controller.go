package handler

import (
	"net/http"
	"ombs/internal/domain"
	"ombs/internal/service"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	AuthService *service.AuthService
}

// NewAuthController initializes AuthController
func NewAuthController(e *echo.Echo, authService *service.AuthService) {
	handler := &AuthController{AuthService: authService}
	e.POST("/auth", handler.Authenticate)
}

// Authenticate receives NodeInfo and returns a response
func (c *AuthController) Authenticate(ctx echo.Context) error {
	node := domain.NodeInfo{
		MessageHash: ctx.FormValue("messageHash"),
		Signature:   ctx.FormValue("signature"),
	}
	response := c.AuthService.AuthenticateNode(node)

	return ctx.JSON(http.StatusOK, echo.Map{"message": response})
}
