package util

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"os"
)

type SignUtil struct {

}

func (signUtil SignUtil) SignPKCS1v15WithSha256(plainText string, privateKey rsa.PrivateKey) string  {
	rng := rand.Reader
	sha256Hash := sha256.Sum256([]byte(plainText))

	sign, err := rsa.SignPKCS1v15(rng, &privateKey, crypto.SHA256, sha256Hash[:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from signing: %s\n", err)
		return "Error from signing"
	}
	return Base64Util{}.Encoding(sign)
}

func (signUtil SignUtil)VerifyPKCS1v15withSha256(signature string, plainText string, pubkey rsa.PublicKey) bool {
	signByte, _ := Base64Util{}.Decoding(signature)
	hashData := sha256.Sum256([]byte(plainText))

	err := rsa.VerifyPKCS1v15(&pubkey, crypto.SHA256, hashData[:], signByte)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from Verify: %s\n", err)
		return false
	}
	return true
}
