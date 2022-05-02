package api

import (
	"github.com/elliot-token/api/app/service"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	SignUp(c *gin.Context)
	GetUser(c *gin.Context)
}

type handler struct {
	userSvc service.UserService
}

func NewHandler(userSvc service.UserService) Handler {
	return &handler{
		userSvc: userSvc,
	}
}
