package api

import (
	CRUD "bank/sql_go"

	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking services.
type Server struct {
	store  *CRUD.Store
	router *gin.Engine
}

// New Server
func NewServer(store *CRUD.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/account", server.createAccount)
	router.GET("/account/:id", server.getAccount)
	router.GET("/account", server.listAccount)
	router.POST("/transfer", server.transferMoney)
	server.router = router
	return server
}

// Runs the server at address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// Resoponse Error
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
