package api

import (
	db "simplebank/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

//提供HTTP 服务
type Server struct {
	store  db.Store
	router *gin.Engine
}

//创建账户server
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	//注册自定义验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency) //绑定的标签名称以及自定义验证器
	}

	router.POST("/users", server.CreateUser)

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	router.POST("/transfers", server.createTransfer)

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
