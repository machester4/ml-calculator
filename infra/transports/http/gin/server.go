package ginserver

import (
	"calculator.com/configs"
	"calculator.com/infra/messaging"
	"calculator.com/infra/transports"
	"calculator.com/internal/engine"
	"github.com/gin-gonic/gin"

	ginhandlers "calculator.com/infra/transports/http/gin/handlers"
	ginmiddleware "calculator.com/infra/transports/http/gin/middleware"
)

type ginServer struct {
	router *gin.Engine
	config configs.ServerConfig
}

var _ transports.Transport = (*ginServer)(nil)

func (s ginServer) InitializeDefaultHandlers(e engine.ServiceEngine) {
	s.router.GET("/", s.wrapHandler(ginhandlers.HandleHealthCheck, e))
}

func (s ginServer) wrapHandler(h handler, e engine.ServiceEngine) gin.HandlerFunc {
	return func(c *gin.Context) {
		h(c, e)
	}
}

type handler func(c *gin.Context, e engine.ServiceEngine)

func (s ginServer) InitializeVoucherHandlers(e engine.ServiceEngine, m messaging.Messaging) {
	s.router.POST("/",
		s.wrapMiddleware(ginmiddleware.MiddlewareVoucherMaxSubset, m),
		s.wrapHandler(ginhandlers.HandleVoucherMaxSubset, e),
	)
}

func (s ginServer) wrapMiddleware(mi middleware, m messaging.Messaging) gin.HandlerFunc {
	return func(c *gin.Context) {
		mi(c, m)
	}
}

type middleware func(c *gin.Context, m messaging.Messaging)

func (s ginServer) Start() error {
	return s.router.Run(s.config.Port)
}

func New(c configs.ServerConfig) *ginServer {
	return &ginServer{
		router: gin.Default(),
		config: c,
	}
}
