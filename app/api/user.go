package api

import (
	"errors"
	"net/http"
	"strings"

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
	if err := c.ShouldBind(&req); err != nil {
		badRequest(c, "invalid body request")
		return
	}

	walletAddr := c.GetString(walletAddrKey)
	if !strings.EqualFold(req.WalletAddress, walletAddr) {
		c.AbortWithStatusJSON(
			http.StatusForbidden,
			errorMessage{
				Error: "wallet address does not match the one in signature",
			},
		)
		return
	}

	if err := h.userSvc.SignUp(&domain.UserEntity{
		WalletAddress: req.WalletAddress,
		Username:      req.Username,
	}); err != nil {
		if errors.Is(err, service.ErrUsernameConflict) || errors.Is(err, service.ErrWalletConflict) {
			c.AbortWithStatusJSON(
				http.StatusConflict,
				errorMessage{
					Error: err.Error(),
				},
			)
			return
		}
		internalServerError(c)
		return
	}

	c.Status(http.StatusCreated)
}

func (h *handler) GetUser(c *gin.Context) {
	walletAddr := c.Param("walletAddr")
	user, err := h.userSvc.GetUser(walletAddr)
	if err != nil {
		if errors.Is(err, service.ErrWalletNotFound) {
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
