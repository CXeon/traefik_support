package http

import (
	"github.com/CXeon/micro_contrib/log"
	"github.com/CXeon/traefik_support/config"
	"github.com/gin-gonic/gin"
)

func initRoutes(conf *config.Config, logger *log.Logger) *gin.Engine {
	r := gin.Default()
	flmGroup := r.Group("/flm")
	appGroup := flmGroup.Group("/traefik-support")

	//创建controller
	c := NewController(conf, logger)

	appGroup.GET("/auth", c.forwardAuth)

	return r
}
