package routes

import (
	"github.com/go_sample/core/token"
	"github.com/go_sample/database"
	"github.com/labstack/echo/v4"
)

type RouteGroup struct {
	Echo       *echo.Echo
	Store      database.Store
	TokenMaker *token.Maker
}

func NewRoutes(echo *echo.Echo, store database.Store, maker token.Maker) *RouteGroup {
	return &RouteGroup{
		Echo:       echo,
		Store:      store,
		TokenMaker: &maker,
	}
}

func (r *RouteGroup) Routes() {
	v1 := r.Echo.Group("/api/v1")
	r.AuthRoutes(v1)
}
