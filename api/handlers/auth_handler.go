package handlers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/go_sample/api/interfaces"
	"github.com/labstack/echo/v4"
)

type AuthHandlers struct {
	validator   *validator.Validate
	authService interfaces.IAuthService
}

func NewAuthHandler(val *validator.Validate, as interfaces.IAuthService) interfaces.IAuthHandlers {
	return &AuthHandlers{
		validator:   val,
		authService: as,
	}
}

// @Tag Auth
// @Summary Login user.
// @Description User can login and get user data with access token.
// @Accept application/json
// @Produce application/json
// @Param email query string true "Input email"
// @Param password query string true "Input password"
// @Success 200 {string} success
// @Router /login [post]
func (h AuthHandlers) Login(ctx echo.Context) error {
	fmt.Printf("hello world")
	return nil
}
