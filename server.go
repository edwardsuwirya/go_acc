package main

import (
	"enigmacamp.com/goacc/config"
	"enigmacamp.com/goacc/delivery/api"
	"enigmacamp.com/goacc/delivery/middleware"
	"enigmacamp.com/goacc/logger"
	"github.com/gin-gonic/gin"
)

type AppServer interface {
	Run()
}

type appServer struct {
	routerEngine *gin.Engine
	cfg          config.Config
}

func (p *appServer) initHandlers() {
	p.routerEngine.Use(middleware.ErrorMiddleware())
	p.v1()
}

func (p *appServer) v1() {
	productApiGroup := p.routerEngine.Group("/product")
	api.NewProductApi(productApiGroup, p.cfg.UseCaseManager.ProductRegistrationUseCase())
}

func (p *appServer) Run() {
	p.initHandlers()
	logger.Log.Info().Msgf("Server run on %s", p.cfg.ApiConfig.Url)
	err := p.routerEngine.Run(p.cfg.ApiConfig.Url)
	if err != nil {
		logger.Log.Fatal().Msg("Server failed to run")
	}

}

func Server() AppServer {
	//gin.SetMode(gin.ReleaseMode)
	//r := gin.New()
	//r.Use(gin.Recovery())
	r := gin.Default()

	c := config.NewConfig(".", "config")
	return &appServer{
		routerEngine: r,
		cfg:          c,
	}
}
