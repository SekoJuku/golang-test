package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinRouter struct {
	Gin *gin.Engine
}

func NewGinRouter() GinRouter {

	httpRouter := gin.Default()

	httpRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Up and Running..."})
	})
	return GinRouter{
		Gin: httpRouter,
	}

}
