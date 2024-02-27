package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server interface {
	Run()
}

type server struct {
	router *gin.Engine
}

func NewServer() Server {
	return &server{
		router: gin.Default(),
	}
}

func (s *server) Run() {
	s.router.GET("/test", testHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: s.router,
	}

	server.ListenAndServe()
}

func testHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}