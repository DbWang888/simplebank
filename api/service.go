package api

import (
	db "simplebank/db/sqlc"

	"github.com/gin-gonic/gin"
)

//提供HTTP 服务
type Server struct {
	store  *db.Store
	router *gin.Engine
}

//创建账户server
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

//监听api请求
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}


func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
