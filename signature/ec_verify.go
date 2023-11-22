package signature

import (
	"bytes"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func VerifySignature(addr string, data []byte, signature []byte) (bool, error) {
	hash := crypto.Keccak256Hash(data)
	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		return false, err
	}
	sigAddr := common.BytesToAddress(crypto.Keccak256(sigPublicKey[1:])[12:])

	matches := bytes.Equal([]byte(sigAddr.String()), []byte(addr))
	return matches, nil
}
