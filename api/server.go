package api

import (
	db "learn/banking/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/account", server.createAccount)
	router.GET("/account/:id", server.getAccount)
	router.GET("/account", server.getListAccount)
	router.PUT("/account/:id", server.updateAccount)
	router.DELETE("/account/:id", server.deleteAccount)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
