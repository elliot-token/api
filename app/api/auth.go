package api

import (
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/elliot-token/api/utils/ether"
	"github.com/gin-gonic/gin"
)

const (
	walletAddrKey = "wallet"
)

type AuthHeaders struct {
	Authorization string `header:"Authorization" binding:"required"`
}

type Token struct {
	Signature string `json:"signature"`
	Body      string `json:"body"`
}

func (h *handler) GetAuth(c *gin.Context) {
	authHeaders := AuthHeaders{}
	if err := c.ShouldBindHeader(&authHeaders); err != nil {
		badRequest(c, "invalid Authorization header")
		return
	}

	authFields := strings.Fields(authHeaders.Authorization)
	if len(authFields) != 2 {
		badRequest(c, "invalid Authorization header: expected 'bearer <token>'")
		return
	}

	if strings.ToLower(authFields[0]) != "bearer" {
		badRequest(c, "invalid Authorization scheme: expected 'bearer'")
		return
	}

	rawToken, err := base64.StdEncoding.DecodeString(authFields[1])
	if err != nil {
		badRequest(c, "token is not a base 64 encoded string")
		return
	}

	token := Token{}
	if err := json.Unmarshal(rawToken, &token); err != nil {
		badRequest(c, "token is not a valid json")
		return
	}

	address, err := ether.GetAddressFromSignature(token.Signature, token.Body)
	if err != nil {
		badRequest(c, "malformed signature")
	}

	// Pass wallet address to the next handler
	c.Set(walletAddrKey, address)
}
