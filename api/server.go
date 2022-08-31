package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/maars202/simplebank/db/sqlc"
)

type Server struct {
	store *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server{
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.CreateAccount)

	server.router = router
	return server
}

// start running the http server on a specific address:
func (server *Server) Start(address string) error{
	return server.router.Run(address)

}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}