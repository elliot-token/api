package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type errorMessage struct {
	Error string `json:"error"`
}

func badRequest(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(
		http.StatusBadRequest,
		errorMessage{
			Error: msg,
		},
	)
}

func internalServerError(c *gin.Context) {
	c.AbortWithStatusJSON(
		http.StatusInternalServerError,
		errorMessage{
			Error: "internal server error",
		},
	)
}
