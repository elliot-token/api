package api

import (
	"errors"
	"net/http"

	"github.com/elliot-token/api/app/domain"
	"github.com/elliot-token/api/app/service"
	"github.com/gin-gonic/gin"
)

type signUpRequest struct {
	WalletAddress string `json:"walletAddress" binding:"required"`
	Username      string `json:"username" binding:"required"`
}

type getUserResponse struct {
	WalletAddress string `json:"walletAddress"`
	Username      string `json:"username"`
}

func (h *handler) SignUp(c *gin.Context) {
	var req signUpRequest
	if err := c.Bind(&req); err != nil {
		// TODO debug log here
		badRequest(c, "invalid body request")
		return
	}

	if err := h.userSvc.SignUp(&domain.UserEntity{
		WalletAddress: req.WalletAddress,
		Username:      req.Username,
	}); err != nil {
		if errors.Is(err, service.ErrUserConflict) {
			c.AbortWithStatusJSON(
				http.StatusConflict,
				errorMessage{
					Error: err.Error(),
				},
			)
			return
		}
		// TODO log error here
		internalServerError(c)
		return
	}

	c.Status(http.StatusCreated)
}

func (h *handler) GetUser(c *gin.Context) {
	walletAddr := c.Param("walletAddr")
	user, err := h.userSvc.GetUser(walletAddr)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		internalServerError(c)
		return
	}

	c.JSON(http.StatusOK, getUserResponse{
		WalletAddress: user.WalletAddress,
		Username:      user.Username,
	})
}
