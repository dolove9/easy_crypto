package util

import (
	"fmt"
	"math/big"

	//"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCertUtil_CertUtilFromFileName(t *testing.T)  {
	assert := assert.New(t)
	fileName := "C:/workspace/2021/dev/golang/resources/cert/signCert.der"
	certUtil := CertUtil{}
	err := certUtil.CertUtilFromFileName(fileName)

	assert.NoError(err)
}

func TestCertUtil_CertBytes(t *testing.T) {
	assert := assert.New(t)

	fileName := "C:/workspace/2021/dev/golang/resources/cert/signCert.der"
	certUtil := CertUtil{}

	err := certUtil.CertUtilFromFileName(fileName)
	assert.NoError(err)

	assert.Equal(certUtil.certificate.Raw, certUtil.CertBytes())
}

func TestCertUtil_SerialNumber(t *testing.T) {
	assert := assert.New(t)

	fileName := "C:/workspace/2021/dev/golang/resources/cert/signCert.der"
	certUtil := CertUtil{}
	certUtil.CertUtilFromFileName(fileName)

	number := certUtil.SerialNumber()
	fmt.Println(number)
	assert.Equal("18814", number)
	//certUtil.certificate.NotBefore
	//certUtil.certificate.NotAfter
}

func TestCertUtil_SerialNumberBigInt(t *testing.T) {
	assert := assert.New(t)

	fileName := "C:/workspace/2021/dev/golang/resources/cert/signCert.der"
	certUtil := CertUtil{}
	certUtil.CertUtilFromFileName(fileName)

	bigInt := certUtil.SerialNumberBigInt()
	assert.Equal(*big.NewInt(18814), bigInt)
}

func TestCertUtil_SubjectDN(t *testing.T) {
	assert := assert.New(t)

	fileName := "C:/workspace/2021/dev/golang/resources/cert/signCert.der"
	certUtil := CertUtil{}
	certUtil.CertUtilFromFileName(fileName)

	assert.Equal("CN=고길동,OU=AccreditedCA+OU=RA,O=Korea Information Certificate Authority,C=KR", certUtil.SubjectDN())
	//fmt.Println(certUtil.SubjectDN())
}

func TestCertUtil_CrlDP(t *testing.T) {
	assert := assert.New(t)

	fileName := "C:/workspace/2021/dev/golang/resources/cert/signCert.der"
	certUtil := CertUtil{}
	certUtil.CertUtilFromFileName(fileName)

	assert.Equal("ldap://192.168.220.134:389/ou=dp3p13,ou=crl,ou=AccreditedCA,o=Korea Information Certificate Authority,c=KR", certUtil.CrlDP())
}

func TestCertUtil_SubjectAltName(t *testing.T) {
	assert := assert.New(t)

	fileName := "C:/workspace/2021/dev/golang/resources/cert/signCert.der"
	certUtil := CertUtil{}
	certUtil.CertUtilFromFileName(fileName)

	subjectAltName, _ := certUtil.SubjectAltName()
	assert.Equal("고길동", subjectAltName.realName)
	assert.Equal("1.2.410.200004.10.1.1.1", subjectAltName.digestAlg)
	assert.Equal("mPRs8hZm21IjHqAmeWbedOgXmHfR5NdYekT3lC3xVqM=", Base64Util{}.Encoding(subjectAltName.vid))

	//fmt.Println(subjectAltName.realName)
	//fmt.Println(subjectAltName.digestAlg)
	//fmt.Println(Base64Util{}.Encoding(subjectAltName.vid))
}