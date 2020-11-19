package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMDUtil_SHA1(t *testing.T)  {
	assert := assert.New(t)

	mdUtil := MDUtil{}
	mdUtil.init("SHA1")
	digest, err := mdUtil.digest([]byte("TEST1234!@#%한글"))
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal("WYMNINcipVy8OG23lR1X4WJB/3c=", digest.base64Result)
	base64Util := Base64Util{}
	decoding, _ := base64Util.Decoding(digest.base64Result)
	assert.Equal(digest.digestResult, decoding)
	fileUtil := FileUtil{}
	fmt.Println(fileUtil.DumpByteArrayToString(digest.digestResult))
}

func TestMDUtil_SHA256(t *testing.T)  {
	assert := assert.New(t)

	mdUtil := MDUtil{}
	mdUtil.init("SHA256")
	digest, err := mdUtil.digest([]byte("Temp"))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(digest.base64Result)
	assert.Equal("1i9Da54bY46l3qGBZYa0Q9o8vOKB+sWnyObKBC8kNjA=", digest.base64Result)
	base64Util := Base64Util{}
	decoding, _ := base64Util.Decoding(digest.base64Result)
	assert.Equal(digest.digestResult, decoding)

	fileUtil := FileUtil{}
	fmt.Println(fileUtil.DumpByteArrayToString(digest.digestResult))
}

func TestMDUtil_SHA256_ERROR_NOALG(t *testing.T) {
	assert := assert.New(t)

	mdUtil := MDUtil{}
	mdUtil.init("SHA123")
	_, err := mdUtil.digest([]byte("Temp"))
	assert.Error(err)
}