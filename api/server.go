package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masterkeysrd/online-payment-platform/api/types"
)

type Server interface {
	Run()
	RegisterController(path string, controller types.Controller)
	RegisterControllerWithMiddleware(path string, controller types.Controller, middleware gin.HandlerFunc)
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
	server := &http.Server{
		Addr:    ":8080",
		Handler: s.router,
	}

	server.ListenAndServe()
}

func (s *server) RegisterController(path string, controller types.Controller) {
	router := s.router.Group(path)
	controller.RegisterRoutes(router)
}

func (s *server) RegisterControllerWithMiddleware(path string, controller types.Controller, middleware gin.HandlerFunc) {
	router := s.router.Group(path)
	router.Use(middleware)
	controller.RegisterRoutes(router)
}
