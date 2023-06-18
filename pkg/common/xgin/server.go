package xgin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"rcsv/pkg/middleware"
	"strconv"
)

type GinServer struct {
	Engine *gin.Engine
}

func NewGinServer() *GinServer {
	var (
		engine *gin.Engine
	)
	gin.SetMode(gin.ReleaseMode)
	engine = gin.New()

	engine.Use(gin.Recovery())

	engine.Use(middleware.Cors())
	return &GinServer{engine}
}

func (s *GinServer) Use(middleware ...gin.HandlerFunc) gin.IRoutes {
	return s.Engine.Use(middleware...)
}

func (s *GinServer) Run(port int) {
	var (
		addr string
		err  error
	)
	addr = ":" + strconv.Itoa(port)
	err = s.Engine.Run(addr)
	if err != nil {
		fmt.Println("GinServer Start Failed.", err.Error())
	}
}
