package api

import (
	db "new/learning/user/db/sqlc"

	"github.com/gin-gonic/gin"
)

// server serves http reqests for our banking service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(store *db.Store) *Server {

	server := &Server{
		store: store,
	}
	router := gin.Default()
	router.POST("/accounts", server.createAccount)
	router.GET("accounts/:id", server.getAccount)
	// router.GET("accounts/", server.ListAccount)
	router.POST("/transfers", server.createTransfer)
	server.router = router
	return server
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
