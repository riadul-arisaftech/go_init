package routes

import (
	"github.com/go_sample/core/token"
	"github.com/go_sample/core/workers"
	"github.com/go_sample/database/repository/interfaces"
	"github.com/labstack/echo/v4"
)

type RouteGroup struct {
	Echo        *echo.Echo
	Store       interfaces.IStore
	TokenMaker  *token.Maker
	Distributor workers.TaskDistributor
}

func NewRoutes(echo *echo.Echo, store interfaces.IStore, maker token.Maker, dist workers.TaskDistributor) *RouteGroup {
	return &RouteGroup{
		Echo:        echo,
		Store:       store,
		TokenMaker:  &maker,
		Distributor: dist,
	}
}

func (r *RouteGroup) Routes() {
	v1 := r.Echo.Group("/api/v1")
	r.AuthRoutes(v1)
}
