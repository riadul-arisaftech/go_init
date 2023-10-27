package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go_sample/api/handlers"
	"go_sample/api/services"
)

func (r RouteGroup) AuthRoutes(route *echo.Group) {
	val := validator.New()
	authService := services.NewAuthService(r.Store, r.TokenMaker)
	authHandlers := handlers.NewAuthHandler(val, authService)

	// routes
	route.GET("/", authHandlers.Login)
}
