package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/go_sample/api/handlers"
	"github.com/go_sample/api/services"
	"github.com/labstack/echo/v4"
)

func (r RouteGroup) AuthRoutes(route *echo.Group) {
	val := validator.New()
	authService := services.NewAuthService(r.Store, r.TokenMaker, r.Distributor)
	authHandlers := handlers.NewAuthHandler(val, authService)

	// routes
	route.GET("/", authHandlers.Login)
}
