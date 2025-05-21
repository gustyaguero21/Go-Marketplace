package server

import "github.com/gin-gonic/gin"

func StartServer() *gin.Engine {
	server := gin.Default()

	UrlMapping(server)

	return server
}
