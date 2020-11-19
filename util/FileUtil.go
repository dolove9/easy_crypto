package util

import (
	"encoding/hex"
	"io/ioutil"
)

type FileUtil struct {

}

func (fu FileUtil) readBytesFromFileName(path string) ([]byte, error)  {
	file, err := ioutil.ReadFile(path)
	return file, err
}

func (fu FileUtil) dumpByteArrayToString(bytes []byte)  string{
	hexString := hex.EncodeToString(bytes)
	return hexString
}

func (fu FileUtil) writeBytesToFile(bytes []byte, fileName string) error  {
	return nil
}