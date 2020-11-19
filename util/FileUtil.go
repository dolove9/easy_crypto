package util

import (
	"encoding/hex"
	"io/ioutil"
)

type FileUtil struct {

}

func (fu FileUtil) ReadBytesFromFileName(path string) ([]byte, error)  {
	file, err := ioutil.ReadFile(path)
	return file, err
}

func (fu FileUtil) DumpByteArrayToString(bytes []byte)  string{
	hexString := hex.EncodeToString(bytes)
	return hexString
}

func (fu FileUtil) WriteBytesToFile(bytes []byte, fileName string) error  {
	return nil
}