package server

import (
	"calculator.com/configs"
	"calculator.com/infra/transports"
	"calculator.com/infra/transports/server/handlers"
	"calculator.com/internal/engine"
	"github.com/gin-gonic/gin"
)

type ginServer struct {
	router *gin.Engine
	config configs.ServerConfig
}

var _ transports.Transport = (*ginServer)(nil)

func (s ginServer) InitializeDefaultHandlers(e engine.ServiceEngine) {
	s.router.GET("/", s.wrap(handlers.HandleHealthCheck, e))
}

func (s ginServer) InitializeVoucherHandlers(e engine.ServiceEngine) {
	s.router.POST("/", s.wrap(handlers.HandleVoucherMaxSubset, e))
}

type handler func(c *gin.Context, e engine.ServiceEngine)

func (s ginServer) wrap(h handler, e engine.ServiceEngine) gin.HandlerFunc {
	return func(c *gin.Context) {
		h(c, e)
	}
}

func (s ginServer) Start() error {
	return s.router.Run(s.config.Port)
}

func NewHttpServer(c configs.ServerConfig) *ginServer {
	return &ginServer{
		router: gin.Default(),
		config: c,
	}
}
