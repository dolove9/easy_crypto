package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMDUtil_SHA1(t *testing.T)  {
	assert := assert.New(t)

	assert.Equal("", "")
}

func TestMDUtil_SHA256(t *testing.T)  {
	assert := assert.New(t)

	mdUtil := MDUtil{}
	mdUtil.init("SHA1")
	_, err := mdUtil.digest([]byte("Temp"))
	if err != nil {
		fmt.Println(err)
	}

	assert.Equal("", "")
}
