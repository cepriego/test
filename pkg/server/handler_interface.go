package server

import "github.com/gin-gonic/gin"

type ServerHandlers interface {
	InitRoutes(router *gin.Engine)
}
