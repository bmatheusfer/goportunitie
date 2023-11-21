package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	// inicializa a Router utilizando as configurações Default do Gin
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	// estamos rodando nossa api
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
