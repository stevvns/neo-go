package transaction

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	//TODO NEO3.0: Update binary
	// https://neotracker.io/tx/2c6a45547b3898318e400e541628990a07acb00f3b9a15a8e966ae49525304da
	rawClaimTX = "020004bc67ba325d6412ff4c55b10f7e9afb54bbb2228d201b37363c3d697ac7c198f70300591cd454d7318d2087c0196abfbbd1573230380672f0f0cd004dcb4857e58cbd010031bcfbed573f5318437e95edd603922a4455ff3326a979fdd1c149a84c4cb0290000b51eb6159c58cac4fe23d90e292ad2bcb7002b0da2c474e81e1889c0649d2c490000000001e72d286979ee6cb1b7e65dfddfb2e384100b8d148e7758de42e4168b71792c603b555f00000000005d9de59d99c0d1f6ed1496444473f4a0b538302f014140456349cec43053009accdb7781b0799c6b591c812768804ab0a0b56b5eae7a97694227fcd33e70899c075848b2cee8fae733faac6865b484d3f7df8949e2aadb232103945fae1ed3c31d778f149192b76734fcc951b400ba3598faa81ff92ebe477eacac"
	// https://neotracker.io/tx/fe4b3af60677204c57e573a57bdc97bc5059b05ad85b1474f84431f88d910f64
	rawInvocationTX = "d101590400b33f7114839c33710da24cf8e7d536b8d244f3991cf565c8146063795d3b9b3cd55aef026eae992b91063db0db53c1087472616e7366657267c5cc1cb5392019e2cc4e6d6b5ea54c8d4b6d11acf166cb072961424c54f6000000000000000001206063795d3b9b3cd55aef026eae992b91063db0db0000014140c6a131c55ca38995402dff8e92ac55d89cbed4b98dfebbcb01acbc01bd78fa2ce2061be921b8999a9ab79c2958875bccfafe7ce1bbbaf1f56580815ea3a4feed232102d41ddce2c97be4c9aa571b8a32cbc305aa29afffbcae71b0ef568db0e93929aaac"
)

func decodeTransaction(rawTX string, t *testing.T) *Transaction {
	b, err1 := hex.DecodeString(rawTX)
	assert.Nil(t, err1)
	tx, err := NewTransactionFromBytes(b)
	assert.NoError(t, err)
	return tx
}
