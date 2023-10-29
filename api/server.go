package api

import (
	"fmt"
	"github.com/go_sample/api/middlewares"
	"github.com/go_sample/api/routes"
	"github.com/go_sample/core/config"
	"github.com/go_sample/core/token"
	"github.com/go_sample/core/workers"
	"github.com/go_sample/database/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HttpServer struct {
	Config *config.Configuration
	Store  repository.Store
	Echo   *echo.Echo
}

func NewServer(config *config.Configuration, store repository.Store, dist workers.TaskDistributor) *HttpServer {
	tokenMaker, err := token.NewPasetoMaker(config.Token.SecretKey)
	if err != nil {
		panic(fmt.Sprintf("cannot create token maker %s", err.Error()))
	}

	echo := echo.New()
	echo.Use(middleware.CORS())
	echo.Use(middlewares.LoggerMiddleware())

	route := routes.NewRoutes(echo, store, tokenMaker, dist)
	route.Routes()

	return &HttpServer{
		Config: config,
		Store:  store,
		Echo:   echo,
	}
}

func (s *HttpServer) Run() {
	err := s.Echo.Start(fmt.Sprintf("%s:%v", s.Config.Server.Host, s.Config.Server.Port))
	if err != nil {
		panic(err.Error())
	}
}
