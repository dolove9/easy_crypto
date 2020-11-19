package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSignUtil_SignPKCS1v15WithSha256(t *testing.T) {
	signUtil := SignUtil{}

	keyUtil := KeyUtil{}
	keyBase64Encode, _ := FileUtil{}.readBytesFromFileName("C:/workspace/2021/dev/golang/resources/cert/dec_private.key")
	decryptedPriKey, _ := Base64Util{}.Decoding(string(keyBase64Encode))

	privateKey, _ := keyUtil.LoadPrivateKey(decryptedPriKey)

	signature := signUtil.SignPKCS1v15WithSha256("TestSign", *privateKey)
	assert := assert.New(t)
	assert.Equal("T3hQOZWAPaeTwFoJBl4ThdmUZMGno5pLOyeq3h/UbamkcyYj7FA8m+heX2tuMgi1IfKtM5xpmVqN2fhshifTkXwqw8oPIbBXKuZWEBId4dYGitN64FKBuxmcArS8dsjYp8lNmbdDDjjwgs6sjd7PoI8nfX0FYMjW5y7+AjdIm/42gPZ6/B9JEh7J3sIQnt5gvoxnSgds/mbGY3Z1Po9QHAuVPLA/o3Jv8snvuFXCWOMZje85G5XI9WU11xSlobY/UPvCZMA3yjeTJyl4gzvD4B0n6lGL5wyVgR2BeAYF54zy1oYYp8IXCXgbz/HhkcMaqcafRA6c0/sO5reg09AkqA==", signature)
}

func TestSignUtil_VerifyPKCS1v15withSha256(t *testing.T) {
	fileName := "C:/workspace/2021/dev/golang/resources/cert/signCert.der"
	certUtil := CertUtil{}
	certUtil.CertUtilFromFileName(fileName)

	publicKey := certUtil.PublicKey()

	signUtil := SignUtil{}
	oriMessage := "TestSign"
	signature := "T3hQOZWAPaeTwFoJBl4ThdmUZMGno5pLOyeq3h/UbamkcyYj7FA8m+heX2tuMgi1IfKtM5xpmVqN2fhshifTkXwqw8oPIbBXKuZWEBId4dYGitN64FKBuxmcArS8dsjYp8lNmbdDDjjwgs6sjd7PoI8nfX0FYMjW5y7+AjdIm/42gPZ6/B9JEh7J3sIQnt5gvoxnSgds/mbGY3Z1Po9QHAuVPLA/o3Jv8snvuFXCWOMZje85G5XI9WU11xSlobY/UPvCZMA3yjeTJyl4gzvD4B0n6lGL5wyVgR2BeAYF54zy1oYYp8IXCXgbz/HhkcMaqcafRA6c0/sO5reg09AkqA=="
	verifyResult := signUtil.VerifyPKCS1v15withSha256(signature, oriMessage, *publicKey)
	fmt.Println(verifyResult)

}
