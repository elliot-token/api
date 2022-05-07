package server

import (
	"fmt"
	"net/http"

	"github.com/elliot-token/api/app/api"
	"github.com/elliot-token/api/config"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func New(srvConf config.Server, handler api.Handler) *http.Server {
	// Default engine Logger and Recovery middleware already attached
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.Default())

	apiV1 := router.Group("/api/v1")
	apiV1.POST("/signup", handler.GetAuth, handler.SignUp)
	apiV1.GET("/users/:walletAddr", handler.GetUser)

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", srvConf.Port),
		Handler: router,
	}
}
