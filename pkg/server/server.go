package server

import (
	"github.com/gin-gonic/gin"
)

const (
	serverAddress = "0.0.0.0:8080"
)

type ShoreServer struct {
	handlers ServerHandlers
}

func NewShoreServer(handlers ServerHandlers) *ShoreServer {
	return &ShoreServer{
		handlers: handlers,
	}
}

func (s *ShoreServer) Init() error {
	// Init Router
	router := gin.Default()
	s.handlers.InitRoutes(router)

	err := router.Run(serverAddress)
	return err
}
