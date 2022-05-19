package api

import (
	"errors"
	"net/http"

	"github.com/elliot-token/api/app/domain"
	"github.com/elliot-token/api/app/service"
	"github.com/ethereum/go-ethereum/common"
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

	addrFromReq, err := common.NewMixedcaseAddressFromString(req.WalletAddress)
	if err != nil {
		badRequest(c, err.Error())
	}

	if err := verifyAddress(c, addrFromReq.Address()); err != nil {
		c.AbortWithStatusJSON(
			http.StatusForbidden,
			errorMessage{
				Error: err.Error(),
			},
		)
		return
	}

	if err := h.userSvc.SignUp(&domain.UserEntity{
		WalletAddress: common.HexToAddress(req.WalletAddress),
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
	addrFromReq, err := common.NewMixedcaseAddressFromString(c.Param("walletAddr"))
	if err != nil {
		badRequest(c, err.Error())
	}

	user, err := h.userSvc.GetUser(addrFromReq.Address())
	if err != nil {
		if errors.Is(err, service.ErrWalletNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		internalServerError(c)
		return
	}

	c.JSON(http.StatusOK, getUserResponse{
		WalletAddress: user.WalletAddress.Hex(),
		Username:      user.Username,
	})
}
