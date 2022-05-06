package utils

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

/*
eyJzaWduYXR1cmUiOiIweDc4ZDZjMTcyZjI0NTMxOWFmMTliN2JiNDRjYTFiNjJkODk0ODQ0YjU4MWVlNWI0NmU0YWI5OGZhYmYzOTY1Y2EwM2Q4ZDg0MjkyMzFhNWZkYTI0OGNhYTJmYzcyMWI0NTAzYjJmODQ5YjczZTQxYTBmNGE1YmIxMTc5MmFjNjA1MWMiLCJib2R5IjoiVVJJOiBodHRwOi8vbG9jYWxob3N0OjMwMDAvXG5XZWIzIFRva2VuIFZlcnNpb246IDJcbk5vbmNlOiA5NzEyODA2OFxuSXNzdWVkIEF0OiAyMDIyLTA1LTAxVDIyOjIzOjA1LjIxMFpcbkV4cGlyYXRpb24gVGltZTogMjAyMi0wNS0wMlQyMjoyMzowNS4wMDBaIn0=
*/

var (
	msg = `URI: http://localhost:3000/
Web3 Token Version: 2
Nonce: 97128068
Issued At: 2022-05-01T22:23:05.210Z
Expiration Time: 2022-05-02T22:23:05.000Z`
	sig = "0x78d6c172f245319af19b7bb44ca1b62d894844b581ee5b46e4ab98fabf3965ca03d8d8429231a5fda248caa2fc721b4503b2f849b73e41a0f4a5bb11792ac6051c"
	//sig = "0x5ad531b00b7696033233bc04ef321913399af16724b104016b21bf567cc8efbe4d05a71b5105ed63890c532332fbb86e8f663e9043e51792d210b20d3659ba321c"
	//msg = "toto"
)

func TestName(t *testing.T) {
	sigDecode := hexutil.MustDecode(sig)
	sigDecode[crypto.RecoveryIDOffset] -= 27

	msgHash := accounts.TextHash([]byte(msg))
	recovered, err := crypto.SigToPub(msgHash, sigDecode)
	if err != nil {
		t.Fatal(err)
	}

	recoveredAddr := crypto.PubkeyToAddress(*recovered)
	t.Log(recoveredAddr.Hex())
}
