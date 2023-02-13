package ports

import (
	"urlShortener/pkg/app"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	app    app.Application
	engine *gin.Engine
}

func NewHttpServer(app app.Application) HttpServer {
	srv := HttpServer{
		app:    app,
		engine: gin.Default(),
	}

	srv.engine.POST("/register", srv.RegisterOrLogin)
	srv.engine.GET("/:id", srv.HandleGetOriginalLink)

	auth := srv.engine.Group("/", srv.authorize)
	auth.POST("/short_link", srv.HandleShorteningLink)

	return srv
}

func (s HttpServer) Run(addr ...string) error {
	return s.engine.Run(addr...)
}
