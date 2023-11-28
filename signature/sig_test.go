package signature

import (
	"fmt"
	"log"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
)

func Test_rc_recovery_addr2(t *testing.T) {
	privateKeyString := "41b68c90ce193729f0e1b73ad492da89431d13b43cf07b6e32a2fdd8458b0170"
	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		log.Fatal(err)
	}
	// Sign message
	signature, err := PersonalSign("test", privateKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(signature)
	addr := "0x72a7e74886b97A4FDfEe192Ecc83e9FbCEdbD6EE"
	verified, err := VerifySignature(addr, []byte("test"), []byte(signature))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(verified)
}
